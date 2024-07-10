#!/usr/bin/env bash

vm_name=$1

cmp_id="ocid1.compartment.oc1..aaaaaaaaxjfveyd7rumokeakkf3ujxsvmdhs7l2oxfjn7zchum3ur4rracta"
subnet_id="ocid1.subnet.oc1.sa-bogota-1.aaaaaaaaeafqiyk3e33b3e6o4e6pfkozmnhefo27yavcrodi2jypl44sfbea"
nsg_ids='["ocid1.networksecuritygroup.oc1.sa-bogota-1.aaaaaaaaynqnwl3wo7go3mmayhqr7o37ndoowxr7eabhn3ot6xhzqjcgcfkq", "ocid1.networksecuritygroup.oc1.sa-bogota-1.aaaaaaaadp4trkdamzsbdvmmcldcpez2varttblwqp7zjbgkwkw6g4mdcswq"]'

secret_id="ocid1.vaultsecret.oc1.sa-bogota-1.amaaaaaalnzgzyyamht6r2wg5w4nucbg2fpst2mnoq2dhn7bj2b63re3uwlq"


touch temp_vm_pub.key
vm_pubkey=$(oci secrets secret-bundle get --secret-id $secret_id --query 'data."secret-bundle-content".content' --raw-output | base64 --decode)

echo "$vm_pubkey" > temp_vm_pub.key

oci compute instance launch \
  --availability-domain DriT:SA-BOGOTA-1-AD-1 \
  --compartment-id $cmp_id \
  --subnet-id $subnet_id --nsg-ids "$nsg_ids" \
  --assign-public-ip true \
  --display-name "$vm_name" \
  --shape VM.Standard3.Flex \
  --shape-config '{ "baselineOcpuUtilization": "BASELINE_1_8", "memoryInGBs":4, "ocpus": 1 }' \
  --image-id "ocid1.image.oc1.sa-bogota-1.aaaaaaaa34rokd6rvzhzi3bfn5nil5utdqptxwbincppslbagperditew7da" \
  --ssh-authorized-keys-file temp_vm_pub.key\
  --user-data-file cloud-init.yaml \
  --wait-for-state RUNNING --wait-for-state TERMINATED --wait-for-state STOPPED --wait-interval-seconds 10  

  rm temp_vm_pub.key
