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
