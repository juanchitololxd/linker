#!/usr/bin/env bash

is_prod=false

while getopts "n:p" opt; do
  case $opt in
    n)
      vm_name=$OPTARG
      ;;
    p)
      is_prod=true
      ;;
    *) echo "usage: $0 [-n] vm_name [-p]" >&2
      exit 1 ;;
  esac
done

#Common resources
lb_id="ocid1.loadbalancer.oc1.sa-bogota-1.aaaaaaaaralkuy6edcv5zmia2sj2go2ddzd4pwnt34jkkejmmsckm2vyzejq"
subnet_id="ocid1.subnet.oc1.sa-bogota-1.aaaaaaaaeafqiyk3e33b3e6o4e6pfkozmnhefo27yavcrodi2jypl44sfbea"

#Linker team specific ids
cmp_id="ocid1.compartment.oc1..aaaaaaaaxjfveyd7rumokeakkf3ujxsvmdhs7l2oxfjn7zchum3ur4rracta"
nsg_ids='["ocid1.networksecuritygroup.oc1.sa-bogota-1.aaaaaaaaynqnwl3wo7go3mmayhqr7o37ndoowxr7eabhn3ot6xhzqjcgcfkq", "ocid1.networksecuritygroup.oc1.sa-bogota-1.aaaaaaaalunetwg7lducrgt6xzmwkjimw2crzpnthc6db4jpbjja2hrmsnha"]'  
secret_id="ocid1.vaultsecret.oc1.sa-bogota-1.amaaaaaalnzgzyyamht6r2wg5w4nucbg2fpst2mnoq2dhn7bj2b63re3uwlq"
bckset_name="linker-1"

echo "Fetching vm secret"
touch temp_vm_pub.key
vm_pubkey=$(oci secrets secret-bundle get --secret-id $secret_id --query 'data."secret-bundle-content".content' --raw-output | base64 --decode)

echo "$vm_pubkey" > temp_vm_pub.key

echo "Provisioning VM"
new_vm_id=$(oci compute instance launch \
  --availability-domain DriT:SA-BOGOTA-1-AD-1 \
  --compartment-id $cmp_id \
  --subnet-id $subnet_id --nsg-ids "$nsg_ids" \
  --assign-public-ip false \
  --display-name "$vm_name" \
  --shape VM.Standard3.Flex \
  --shape-config '{ "baselineOcpuUtilization": "BASELINE_1_8", "memoryInGBs":4, "ocpus": 1 }' \
  --image-id "ocid1.image.oc1.sa-bogota-1.aaaaaaaa34rokd6rvzhzi3bfn5nil5utdqptxwbincppslbagperditew7da" \
  --ssh-authorized-keys-file temp_vm_pub.key\
  --user-data-file cloud-init.yaml \
  --wait-for-state RUNNING --wait-for-state TERMINATED --wait-for-state STOPPED --wait-interval-seconds 10 \
  --query 'data.id' --raw-output) 

rm temp_vm_pub.key

if $is_prod ; then
echo "Replacing LB backend"
empty_bcks=$(oci lb backend list --backend-set-name $bckset_name --load-balancer-id $lb_id --all --query 'data[].name')

echo "$empty_bcks" | jq -r '.[]' | while read bck;
do
  echo "To delete $bck"
  oci lb backend update --load-balancer-id $lb_id --backend-set-name $bckset_name --backend-name "$bck" --backup false --drain true --offline true --weight 1 --wait-for-state SUCCEEDED --wait-for-state FAILED;
  oci lb backend delete --load-balancer-id $lb_id --backend-set-name $bckset_name --backend-name "$bck"  --force --wait-for-state SUCCEEDED --wait-for-state FAILED;
done

new_pip=$(oci compute instance list-vnics --instance-id "$new_vm_id" --query 'data[0]."private-ip"' --raw-output)

oci lb backend create --load-balancer-id $lb_id --backend-set-name $bckset_name --ip-address "$new_pip" --port 8080 --wait-for-state SUCCEEDED --wait-for-state FAILED

echo "Added new ØB backend with IP $new_pip"
fi
