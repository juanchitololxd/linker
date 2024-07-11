# Linker

## Escuela Colombiana de Ingeniería Julio Garavito

### Integrantes:

- [Juan Pablo Fonseca Cardenas](https://github.com/juanchitololxd)
- [Diego Alejandro Murcia Cespedes](https://github.com/Diego-Murcia)
- [Juan Sebastian Garcia Hincapie](https://github.com/jgarciahincapie)
- [Juan Pablo Daza Pinzon](https://github.com/JuanPabloDaza)

### Índice

- [¿Qué es linker?](#qué-es-linker)
- [Instrucciones y uso](#instrucciones-y-uso)
- [¿Cómo ayudar en el desarrollo?](#cómo-ayudar-en-el-desarrollo)


## ¿Qué es linker?

Linker es un acortador de URL que tiene como finalidad el aprendizaje de distintas técnicas de 
desarrollo impulsadas por el profesor [Sebastian Aguado Bedoya](https://github.com/saguadob) en el desarrollo del curso
[Desarrollo de aplicaciones basadas en arquitecturas nativas de la nube y metodologías DevSecOps](https://www.escuelaing.edu.co/es/programas/curso-desarrollo-de-aplicaciones-basadas-en-arquitecturas-nativas-de-la-nube-y-metodologias-devsecops/)
dictado en el periodo intersemestral 2024-I en la Escuela Colombiana de Ingeniería Julio Garavito.

Como el nombre del curso lo indica, linker esta pensado para poder ser desplegado en un ambiente de nube y poder analizar y
experimentar el desarrollo de una aplicación nativa de la nube y el uso de distintas herramientas y prácticas que componen las metodologías DevSecOps.

## Instrucciones y uso

Linker esta desarrollado en varios lenguajes, para su frontend se usa [HTML](https://developer.mozilla.org/es/docs/Web/HTML), [JavaScript](https://developer.mozilla.org/es/docs/Web/JavaScript)y como librería [Bootstrap](https://getbootstrap.com). Para su backend, [Go](https://go.dev).

### Guia para usar linker de manera remota:

Para poder ejecutar linker es necesario clonar este repositorio:

![](/img/CloneRepo.png)<br>

Ahora debemos ingresar al proyecto a travez de la terminal y ejecutar los comandos:

```
go mod tidy
```

Este comando se usa para limpiar y actualizar dependencias.

```
go run main.go
```

Este comando ejecutara linker.

Ahora debemos acceder a [localhost:8080](http://localhost:8080) y podremos ver la página de linker:

![](/img/Linker1.png)<br>

Para probar el funcionamiento deberás copiar el link que deseas acortar, para ese paso usaremos un video de YouTube como ejemplo:

![](/img/YoutubeVideo.png)<br>

Posteriormente vas a pegar el link copiado en la seccion que dice “url” dentro de la aplicacion de linker:

![](/img/PasteURL.png)<br>

Ahora, presiona el botón y veras como se genera el link acortado:

![](/img/ShortenURL.png)<br>

Puedes darle click al boton copiar:

![](/img/CopyButton.png)<br>

Ahora, ve a tu navegador, y pega el link copiado y dale enter:

![](/img/PasteShortenURL.png)<br>

Si todo salió bien ahora puedes accceder al recurso original:

![](/img/Video.png)<br>

### Despliegue de linker en una ambiente de nube:

Como ya habiamos mencionado anteriormente linker esta diseñado para poder ser desplegado en la nube, en este caso y por efectos del curso estaremos usando maquinas virtuales proporcionadas por Oracle.

Linker ya esta desplegado en estas, usando este [link](http://1.unli.ink) se pude comprobar.

Para poder llegar a desplegar linker tenemos un proceso que detallaremos y explicaremos su paso a paso.

Para empezar, una de las herramientas que vamos a usar es proporcionada por GitHub, esta es [GitHub Actions](https://docs.github.com/en/actions) con la cual podemos crear [workflows](https://docs.github.com/en/actions/using-workflows)
que nos ayudan a automatizar todo el despliegue de linker en nuestro ambiente de nube.

### Worflow

Nuestro workflow se encarga de compilar el codigo del proyecto y ejecutar las pruebas con la finalidad de probar el correcto funcionamiento del código antes de realizar un despliegue en la nube, el código de este proceso esta en el archivo [go.yml](/.github/workflows/go.yml).

![](/img/Workflow1.PNG)<br>

Durante este proceso se genero un ejecutable, url-shortener.exe el cual es el resutlado de el comando go build -v ./... el siguiente paso es enviar ese ejecutable a nuestra maquina virtual.

![](/img/Workflow2.PNG)<br>

Vamos a ver un poco el archivo [go.yml](/.github/workflows/go.yml) para ver que succede en la transferencia del ejecutable:

![](/img/Workflow3.PNG)<br>

En el archivo podemos ver que usamos nuestra llave para poder establecer la conexión con nuestra maquina virtual, esta llave
esta guardada como un secreto dentro del repositorio de GitHub. Luego de eso, enviamos nuestro [script](/scripts/kill-current-process.sh) el cual se encarga de eliminar cualquier proceso que este usando el puerto :8080 ya que es el que usa linker
para escuchar las peticiones.

Una vez ejecutamos nuestro script, entonces ahora enviamos el ejecutable url-shortener, y lo ejecutamos en segundo plano.

Si todo esta bien podemos comprobar que [linker](http://1.unli.ink) funciona.

## ¿Cómo ayudar en el desarrollo?

Linker esta pensado para seguir la estrategia de [trunk based development](https://trunkbaseddevelopment.com) y tambien el principio de [small batch development](https://dora.dev/capabilities/working-in-small-batches/).
Con esto en mente, para poder ayudarnos en linker es necesario la creación de [ramas](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/about-branches) en donde se desarrollen pequeñas funcionalidades
o caracteristicas que luego puedan ser revisables por cualquiera de manera sencilla.

Para poder incluir esta funcionalidad es necesario crear un [pull request](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/about-pull-requests)
con la finalidad de que este sea revisado por los integrantes del proyecto y pueda ser aceptado o rechazado y se pueda dar una retroalimentación de la decisión.

También tenemos a disposicion otra herrameinta de GitHub, [GitHub Codespaces](https://github.com/features/codespaces) la cual permite crear ambientes ailados de desarrollo
con la finalidad de que cualquier persona desde el navegador pueda realizar cambios y aportar a nuestro proyecto.