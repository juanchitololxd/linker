`launch-new-vm.sh`
Script utilizado para crear entornos bajo demanda para su aplicación `linker` 

Esta nueva version del script se puede invocar de dos maneras

```sh
sh launch-new-vm.sh -n "change_vm_name" -p
```

si se invoca con la opción `-p` el script crea una nueva VM y modifica el balanceador de cargas para que la url `su_equipo.unli.ink` use la nueva VM usando el puerto 8080. Utilice esta VM como entorno de despliegue para el entorno de Producción

```sh
sh launch-new-vm.sh -n "change_vm_name"
```

si se invoca sin la opcion `-p` se crea una VM que se asume que utilizará para el entorno de desarrollo, y que va a ser utilizada en el pipeline como el primer entorno de despligue, y que ud utilizará para probar que su aplicación funciona.


>  Atención
>  Su equipo solo puede utilizar 2 VM simultaneas. Intentar desplegar una 3ra maquina virtual resultará en error.
