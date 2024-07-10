# Laboratorio

## 1. Instale OCI CLI en su entotno local
Sigas las instrucciones de la guía [Instalar OCI CLI](https://docs.oracle.com/en-us/iaas/Content/API/SDKDocs/cliinstall.htm)

## 2. Crear su propia copia del script para crear un entorno
En su repositorio encontrará un nuevo archivo llamado `scripts/launch-new-vm.sh`. Está preconfigurado para desplegar una maquina virtual en su compartimiento.
El scrip se puede ejecutar de la siguiente manera:
```sh
sh scripts/launch-new-vm.sh "nombre-de-su-vm"
```
- La nueva VM es Ubuntu
- El usuario por defecto es `ubuntu`

Ahora tiene que modificar el archivo de `cloud-init.yaml` para que tenga el `runtime` de su aplicación linker

Ejecute el script para crear una máquina virtual en su entorno.

## 3. __Juegue__ con el pipeline y su nueva máquina
1. Cree un script que se llame `.github/workflows/ci.yml` use como guía el [Archivo en Linker](https://github.com/co-eiv-devsecops/linker/blob/main/.github/workflows/ci.yml). Como sugerencia le recomiendo que los pasos de su pipeline sean el siguiente
  - Compilar
  - Prueba unitaria
  - Empaquetar
  - Desplegar a su entorno

Utilice la capabilidad de [Github Secrets](https://docs.github.com/en/actions/security-guides/using-secrets-in-github-actions?tool=webui) para alamcenar los detalles de conexión a su VM

## 4. Cómo puedo probar qué linker funciona?
Su maquina virtual corre pero solo es accesible por medio del puerto 22, no más puertos están expuestos, ya que esto es una práctica común de seguridad
Para verificar que su aplicación funciona ud puede crear un sesión de _port forwarding_ usando el comando
```sh
ssh -i sullave.key -N -L 8080:localhost:8080 ubuntu@direccion_ip -vvv
```
