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
- [Postmortem](#postmortem)
- [Change log](#change-log)


## ¿Qué es Linker?

Linker es un acortador de URL que tiene como finalidad el aprendizaje de distintas técnicas de 
desarrollo impulsadas por el profesor [Sebastian Aguado Bedoya](https://github.com/saguadob) en el desarrollo del curso
[Desarrollo de aplicaciones basadas en arquitecturas nativas de la nube y metodologías DevSecOps](https://www.escuelaing.edu.co/es/programas/curso-desarrollo-de-aplicaciones-basadas-en-arquitecturas-nativas-de-la-nube-y-metodologias-devsecops/)
dictado en el periodo intersemestral 2024-I en la Escuela Colombiana de Ingeniería Julio Garavito.

Como el nombre del curso lo indica, linker esta pensado para poder ser desplegado en un ambiente de nube y poder analizar y
experimentar el desarrollo de una aplicación nativa de la nube y el uso de distintas herramientas y prácticas que componen las metodologías DevSecOps.

## Instrucciones y uso

### Tecnologías usadas

Linker está desarrollado en varios lenguajes. Para el frontend, se utiliza [HTML](https://developer.mozilla.org/es/docs/Web/HTML), [JavaScript](https://developer.mozilla.org/es/docs/Web/JavaScript) y, como librería, [Bootstrap](https://getbootstrap.com). Para el backend, se utiliza [Go](https://go.dev) y almacena los enlaces que se acortan en una base de datos [MySQL](https://www.mysql.com).

Como parte de la metodología de trabajo del proyecto se usaron diferentes herramientas para lograr los objetivos propuestos. Para los pipelines usamos GitHub Actions,
implementamos [Prometheus](https://prometheus.io) en la aplicación para generar metricas y [Grafana](https://grafana.com) para poder visualizarlas.

### Guía para usar Linker de manera remota:

Para poder ejecutar linker es necesario clonar este repositorio:

![](/img/CloneRepo.png)<br>

Ahora debemos ingresar al proyecto a través de la terminal y se deben ejecutar los siguientes comandos:

```
go mod tidy
```

Este comando se usa para limpiar y actualizar dependencias.

```
.\scripts\setup-project-linux.sh dev
```
Este comando para configurar el environment para que el proyecto tenga configuradas las variables de entorno (deberás pedir a soporte@linker1.com credenciales para la BD y sobreeescribirlas en el archivo .env del root). Si usas windows deberás usar `.\scripts\setup-project-windows.ps1 -RUN_ENV dev`

```
go run ./cmd/api/main.go
```

Este comando ejecutará Linker.

Ahora debemos acceder a [localhost:8080](http://localhost:8080) y podremos ver la página de Linker:

![](/img/Linker1.png)<br>

Para probar el funcionamiento, deberás copiar el link que deseas acortar, para ese paso usaremos un video de YouTube como ejemplo:

![](/img/YoutubeVideo.png)<br>

Posteriormente vas a pegar el enlace copiado en la sección que dice “url” dentro de la aplicacion de linker:

![](/img/PasteURL.png)<br>

Ahora, presiona el botón y veras como se genera el link acortado:

![](/img/ShortenURL.png)<br>

Puedes darle click al botón copiar:

![](/img/CopyButton.png)<br>

Ahora, ve a tu navegador, y pega el link copiado y dale enter:

![](/img/PasteShortenURL.png)<br>

Si todo salió bien ahora puedes acceder al recurso original:

![](/img/Video.png)<br>

### Despliegue de Linker en un ambiente de nube:

Como ya habíamos mencionado anteriormente Linker está diseñado para poder ser desplegado en la nube, en este caso y por efectos del curso estaremos usando máquinas virtuales proporcionadas por Oracle.

Linker ya esta desplegado en estas, usando este [link](http://1.unli.ink) se puede comprobar.

Para poder llegar a desplegar Linker tenemos un proceso que detallaremos y explicaremos su paso a paso.

Para empezar, una de las herramientas que vamos a usar es proporcionada por GitHub, esta es [GitHub Actions](https://docs.github.com/en/actions), con la cual podemos crear [workflows](https://docs.github.com/en/actions/using-workflows)
que nos ayudan a automatizar todo el despliegue de linker en nuestro ambiente de nube.

### Workflow

Nuestro workflow tiene dos pasos, una vez se hace un pull request hacia la rama main, se verifica la calidad del código, se ejecutan pruebas y se compila para 
poder evitar problemas con la ejecución, luego de esto de comienza la trasferencia del ejecutable generado hacia nuestra maquina de desarrollo con la finalidad de probar los cambios en un ambiente controlado, y luego de que se comprueba que Linker funciona en el ambiente de desarrollo si se acepta el pull request 
entonces se realizara el mismo proceso pero para el ambiente de producción.

Los detalles de este proceso pueden ser vistos en los archivos [quilitycode.yml](./.github/workflows/qualitycode.yml) y [deploy.yml](./.github/workflows/deploy.yml), tambien el proceso esta mas detallado en nuestra wiki, en el apartado de la [Guía de uso](https://github.com/co-eiv-devsecops/linker-1-app/wiki/Guia-de-uso).

Una vez ejecutamos nuestro script, entonces ahora enviamos el ejecutable url-shortener, y lo ejecutamos en segundo plano.

Si todo está bien podemos comprobar que [Linker](http://1.unli.ink) funciona.

## ¿Cómo ayudar en el desarrollo?

Linker está pensado para seguir la estrategia de [trunk based development](https://trunkbaseddevelopment.com) y también el principio de [small batch development](https://dora.dev/capabilities/working-in-small-batches/).
Con esto en mente, para poder ayudarnos en linker es necesario la creación de [ramas](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/about-branches) en donde se desarrollen pequeñas funcionalidades
o características que luego puedan ser revisables por cualquiera de manera sencilla.

Para poder incluir esta funcionalidad es necesario crear un [pull request](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/about-pull-requests)
con la finalidad de que este sea revisado por los integrantes del proyecto y pueda ser aceptado o rechazado y se pueda dar una retroalimentación de la decisión.

También tenemos a disposición otra herramienta de GitHub, [GitHub Codespaces](https://github.com/features/codespaces) la cual permite crear ambientes aislados de desarrollo
con la finalidad de que cualquier persona desde el navegador pueda realizar cambios y aportar a nuestro proyecto.

## Postmortem

Como parte de uno de los ejercicios que trabajamos en clase, realizamos un análisis postmortem de un error en otra aplicación, tenemos desplegado como una pagina haciendo uso de GitHub Pages, el analisis se puede ver [aquí](https://co-eiv-devsecops.github.io/linker-1-app/).

## 12 principios

Uno de los objetivos del proyecto es crear una aplicación que cumpla con la metodología de los [doce principios](https://12factor.net/es/), con esto en mente
vamos a desglosar como aplicamos cada principio dentro de Linker, esto está detallado dentro de nuestra [wiki](https://github.com/co-eiv-devsecops/linker-1-app/wiki/12-Principios)

## Change Log
### v 1.0.1
- Configuración de variables de entorno 
- Uso de secrets en el pipeline para reemplazar datos privados como usuario y contraseña de la BD.
- Scripts que generan el archivo .env (se debe de ejecutar el correspondiente antes de correr el proyecto o de hacer tests)