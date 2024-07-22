#!/bin/bash

# Verificar si se proporcionó el nombre de la VM como argumento
if [ -z "$1" ]; then
    echo "Uso: $0 <nombre_de_la_VM>"
    exit 1
fi
echo $1

# Variables
vm_name="$1"  # Nombre de la VM pasado como argumento
compartment_id="ocid1.compartment.oc1..aaaaaaaaxjfveyd7rumokeakkf3ujxsvmdhs7l2oxfjn7zchum3ur4rracta"  # Reemplaza con el OCID de tu compartimento

# Obtener el ID de la instancia basada en su nombre
instance_id=$(oci compute instance list --compartment-id $compartment_id --display-name $vm_name --query 'data[0].id' --raw-output)

# Verificar si se encontró la instancia
if [ -z "$instance_id" ]; then
    echo "No se encontró ninguna instancia con el nombre $vm_name en el compartimento $compartment_id."
    exit 1
fi

# Eliminar la instancia
oci compute instance terminate --instance-id $instance_id --force
echo "La instancia con el nombre $vm_name y ID $instance_id ha sido eliminada."