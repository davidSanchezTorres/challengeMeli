# Challenge MercadoLibre

_API con algoritmo para reclutar mutantes._

## Función 🚀

_Detectar si un humano es mutante basándose en su secuencia de ADN. El algoritmo buscar secuencias de 4 secuenciales de forma horizontal, vertical y cualquier oblicua_

Mira **Deployment** para conocer como desplegar el proyecto.


### Pre-requisitos 📋

_Descargar el proyecto y ejecuta el siguiente comando_

```
go mod vendor
```


### Cómo ejecutar la alicación

_Después de ejecutar el comando "go mod vendor" ejecuta el siguiente comando en la raíz del programa_

```
go run main.go
```

_El programa arrancará en localhost:3000_

_Con los dos servicios /v1/mutant/ y /v1/stats/_

_Está desplegado en AWS en una instancia gratuita y podrán ser consumidas como se muestra a continuación_

_POST_
```
ec2-3-18-225-241.us-east-2.compute.amazonaws.com/v1/mutant/
```
Body
```
{
    "dna":["ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"]
}
```

_GET_
``` 
ec2-3-18-225-241.us-east-2.compute.amazonaws.com/v1/stats/
```

_Comparto la colección de Postman para importación. Nombre de archivo "Mutant collection.postman_collection.json"._

_Los datos serán guardados en la base de datos Redis. La clave es el AND y el valor true o false según corresponda si es mutante o no._

_Ejemplo de guardado en base de datos_

_ATGCGA+CAGTGC+TTATGT+AGAAGG+CCCCTA+TCACTG+ true_

_Se guarda 1 por and ya que la clave es la secuencia y por lo tanto no puede ser repetida_

## Ejecutando las pruebas ⚙️

_Para ejecutar las pruebas y obtener el coverage, te situas en la raíz del proyecto y ejecuta los siguientes comandos_

_Para Linux_
```
go test -coverprofile="coverage.out" ./...
```

_Para Windows_
```
go test .\... -coverprofile="coverage.out"
```

_Para Windows y Linux_
```
go tool cover --func=coverage.out
```

```
go tool cover --html=coverage.out
```
