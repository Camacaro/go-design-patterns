package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var (
	port = flag.Int("p", 3090, "port to scan")
	host = flag.String("h", "localhost", "url del sitio a escanear")
)

/*
	Para los que no sepan que es netcat, es una herramienta “simple”
	que nos sirve para hacer conexiones tcp o udp y enviar datos por
	las mismas, se puede enviar por ejemplo, una petición http (o si
	nos vamos a algo mas malvado, una reverse Shell).
	Si quieren hacer una petición http solo tienen que levantar un server,
	y escribir las siguientes líneas en que programa que realizamos
*/

/*
	Este es nuestro cliente

	-> host:port
	Escribir -> host:port
	Leer -> host:port

	Hola -> host:port -> [Hola]
*/
func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))

	if err != nil {
		log.Fatal(err)
	}

	// Se deja un strcuct muchas veces para indicar que no nos interesa
	// el valor que devuelve, solo nos interesa que bloquee la funcion main
	done := make(chan struct{})

	// Todo lo que lee de esa conexion
	go func() {
		/*
			Nuestro escritor sera el stdout, ya que el stdout es el que
			escribe en la consola

			conn sera el lector, ya que el lector es el que lee de la conexion
		*/
		io.Copy(os.Stdout, conn)

		// Cuando termine de leer, cerramos el canal, para que continue
		done <- struct{}{}
	}()

	CopyContents(conn, os.Stdin)

	conn.Close()
	<-done
}

/*
	dest: destinatario que es el escritor
	src un lector
*/
func CopyContents(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)

	if err != nil {
		log.Fatal(err)
	}

}
