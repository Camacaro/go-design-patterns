# Go avanzado: concurrencia, patrones de diseño y net

Go soporta la concurrencia de manera nativa

Go utiliza un modelo de concurrencia llamado CSP
(Communicating Sequential Processes)

Estas goRoutines son inicializadas con un stack de 2kb. En comparación, Java utiliza threads de apoximadamente 1MB (Cada threads)

Para comunicarse se utilizan Channels que pueden
ser "buffered" o "unbuffered". (Con buffered podemos trabajar los canales como semaforos ejecutar algunos en concurrencia mientras otros esperan que algun canal sea liberado)

Diferentes técnicas para sincronización, tales como WaitGroups e iteración en Channels. Esto es para bloquear el programa principal (main) y esperar que terminen las goRoutine, ir las escuchando de ser necesario

Go implementa interfaces de manera implicita sin palabras clave como implements

Mediante la utilizacion de Go Modules podemos escapar de la presion del GOPATH. (Antes en la version 1 siempre debiamos estar debajo de la GOPATH)

```go test``` nos ayudará a ejecutar los test que hayamos escrito para nuesto programas

Librerias Standars: Prometheus, Helm, etcd, containerd, Jaeger, gRPC y NATS son algunos de los proyectos que más llaman la atención en Go.


## En este proyecto
* Concurrencia de un valor compartido: cómo puedo sincronizar las GoRoutines cuando están accediendo al mismo valor y evitar la condición de carrera.
* Patrones de diseño: aquellos más frecuentes y que mejores resultados traerán a nuestros problemas.
* Utilizar el paquete Net para la creación de conexiones TCP que permiten diferentes utilidades como escanear puertos entre otros.

## Patrones de diseño 
Patrones de diseño: Son planos que nos ayuda a solventar problemas muy comunes a la hora de diseñar software, estos planos los podemos adaptar a nuestras necesidades y a lenguajes de programación específicos.

Los patrones de diseño se categorizan en:
* Patrones creacionales: Establecen mecanismos para que la creación de objetos pueda ser reutilizable y flexible. Ejemplo: **Factory** y **Singleton**.
* Patrones estructurales: Establecen mecanismos de como crear objetos en estructuras mas grandes sin perder esa flexibilidad y reusabilidad. Ejemplo: **Adapter**.
* Patrones de comportamiento: Establecen mecanismos de comunicación efectiva entre estos objetos, asimismo la asignación de responsabilidades de estos. Ejemplo: **Observer** y **Strategy**.

## Recursos 
* [explicación por patron y ejemplos](https://refactoring.guru/es/design-patterns/go)
* [Patrones de Diseño en Go - E-book](https://github.com/danielspk/designpatternsingoebook)
* [los patrones de diseño de Gang of Four](https://springframework.guru/gang-of-four-design-patterns/)


