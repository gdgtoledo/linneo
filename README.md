# Linneo
Side project de la comunidad del GDG Toledo.

Su objetivo es posicionar en un mapa los Almeces existentes en la ciudad.

## Desarrollo

El lenguaje utilizado para crear el API será Golang, utilizando `Go modules`. De este modo es necesario que la variable de entorno `GO111MODULE` está activada con un valor `on`.

```shell
# Si usas la terminal de VSCOde, ejecuta esta línea en ella cada vez que lo abras
$ export GO111MODULE=on

# o abre el IDE así para que tenga acceso a los Go modules, con el entorno ya configurado
$ export GO111MODULE=on && code /path/al/proyecto
```