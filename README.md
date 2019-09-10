# Linneo
Side project de la comunidad del GDG Toledo.

Su objetivo es posicionar en un mapa los Almeces existentes en la ciudad.

## Desarrollo

Requisitos para el desarrollo:

- Golang v1.12.7, o superior
- Docker Compose v1.24.1, o superior

El lenguaje utilizado para crear el API será Golang, utilizando `Go modules`. De este modo es necesario que la variable de entorno `GO111MODULE` está activada con un valor `on`.

```shell
# Si usas la terminal de VSCOde, ejecuta esta línea en ella cada vez que lo abras
$ export GO111MODULE=on

# o abre el IDE así para que tenga acceso a los Go modules, con el entorno ya configurado
$ export GO111MODULE=on && code /path/al/proyecto
```

## Ejecución

El proyecto usa `Make` como wrapper para lanzar los comandos, que usa internamente `Docker Compose` para levantar el stack en local. Deberías lanzar los comandos en el siguiente orden para tener el proyecto corriendo en un estado adecuado:

```shell
$ make build   # construye la imagen Docker con el binario de Go
$ make start   # arranca el stack utilizando Docker Compose
$ make seed    # crea los índices adecuados (espera a que Elasticsearch arranque)
$ make destroy # destruye el stack utilizando Docker Compose down y haciendo un prune de los volúmenes
```