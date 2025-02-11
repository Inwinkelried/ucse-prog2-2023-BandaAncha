# ucse-prog2-2023-BandaAncha

Trabajo integrador para la materia Programación 2. UCSE DAR 2023.

Integrantes del grupo: 
- Facundo Schillino
- Santiago Inwinkelried

## Estructura
Tenemos tres directorios principales en el root del proyecto:
* go: es la API, el backend.
* web: es el frontend con HTML, CSS y JavaScript.
* data: archivos JSON para importar a mongoDB y agilizar las pruebas.

## Instrucciones para levantar el proyecto
1. Abrir una terminal parados en el root, y correr el comando `docker-compose up`.
2. En el explorador de preferencia, ingresar a `localhost:80` para visualizar el frontend.
### Para usar datos de prueba del directorio "data"
1. Abrir MongoDB Compass y crear las colecciones `Pedidos`, `Productos`, `Camiones` y `Envios` dentro de la base de datos `BandaAncha` (se crea automáticamente luego de abrir el frontend).
2. En cada colección, importar el archivo .json con el mismo nombre
