Este proyecto está hecho para que, dado el excel que agregues a la carpeta root de este proyecto (donde está el archivo main.go), puedas leerlo, tomar todas las columnas de informacion, almacenarlas en mapas y mandarlas a un tercero, todo esto para tener "bases de datos" agiles (ya que la info guardada en una BBDD será la que esté en el excel)

-------------------------------------------------------------------------------

Pasos de instalacion y uso:

Antes que nada, necesitarás una base de datos en firestore y una cuenta de servicio que tenga, al menos, los permisos para escribir en dicha base. Luego...

1.- Clone este proyecto usando git clone 

2.- En la carpeta del proyecto crea un archivo .env que contenga;

GOOGLE_APPLICATION_CREDENTIALS=path/to/your/service/account
excel_name=excel-name (sin el '.xlsx')
destiny_collection=firestore_collection (donde se guardará)

3.- Agregue el excel a la carpeta root del proyecto (donde está el archivo main.go)

4.- Ejecute en la consola el comando "go run main.go" 

El programa leerá el excel, tomará el nombre de las columnas y luego, por cada fila, lo guardará en un map (map[string]string) y lo mandará al cliente de firestore para ser guardado.

