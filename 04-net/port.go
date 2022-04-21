package main

import (
	"fmt"
	"net"
	"time"
)

/*
	Este paquete NET nos permite crear coneciones TCP a diferentes
	servidores y escanear los diferentes puertos que tenga
*/
func main() {
	start := time.Now()

	/*
		Puerto 21 abierto, Time 4.235809583s
		Puerto 22 abierto, Time 4.44112275s
		Puerto 25 abierto, Time 4.862002167s
		Puerto 80 abierto, Time 16.357994708s
	*/

	for i := 0; i < 100; i++ {
		// 1, 2, ..., 99
		// (puerto) Sitio:1, Sitio:2, ..., Sitio:99
		// ese puerto 1 -> open, 2 -> closed, ..., 99 -> open

		/*
			Primer Parametro: Protocolo
			Segundo Parametro: Sitio:Puerto
		*/
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "scanme.nmap.org", i))

		if err != nil {
			// Puerto cerrado
			continue
		}

		conn.Close()
		fmt.Printf("Puerto %d abierto, Time %s \n", i, time.Since(start))
	}
}
