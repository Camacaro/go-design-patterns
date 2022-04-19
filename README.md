# Go avanzado: concurrencia, patrones de diseño y net

Go soporta la concurrencia de manera nativa

Go utiliza un modelo de concurrencia llamado CSP
(Communicating Sequential Processes)

Estas goRoutines son inicializadas con un stack de 2kb. En comparación, Java utiliza threads de apoximadamente 1MB (Cada threads)

Para comunicarse se utilizan Channels que pueden
ser "buffered" o "unbuffered". (Con buffered podemos trabajar los canales como semaforos ejecutar algunos en paralelo en esos canales habilitados cuando se desocupen se agregan los otros)

Diferentes técnicas para sincronización, tales como WaitGroups e iteración en Channels. Esto es para bloquear el programa principal (main) y esperar que terminen las goRoutine, ir las escuchando de ser necesario

Go implementa interfaces de manera implicita sin palabras clave como implements

Mediante la utilizacion de Go Modules podemos escapar de la presion del GOPATH. (Antes en la version 1 siempre debiamos estar debajo de la GOPATH)

```go test``` nos ayudará a ejecutar los test que hayamos escrito para nuesto programas

Librerias Standars: Prometheus, Helm, etcd, containerd, Jaeger, gRPC y NATS son algunos de los proyectos que más llaman la atención en Go.


## En este proyecto
* Concurrencia de un valor compartido: cómo puedo sincronizar las GoRoutines cuando están accediendo al mismo valor y evitar la condición de carrera.
* Patrones de diseño: aquellos más frecuentes y que mejores resultados traerán a nuestros problemas.
* Utilizar el paquete Net para la creación de conexiones TCP que permiten diferentes utilidades como escanear puertos entre otros.