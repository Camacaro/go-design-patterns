package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

/*
	Primer Parametro: Nombre del flag
	Segundo Parametro: Valor por defecto
	Tercer Parametro: Descripcion del flag
*/
var site = flag.String("site", "scanme.nmap.org", "url del sitio a escanear")

/*
	Este paquete NET nos permite crear coneciones TCP a diferentes
	servidores y escanear los diferentes puertos que tenga

	Flags:
	go run portConcurrente.go --site=scanme.nmap.org
	go run portConcurrente.go --site=scanme.webscantest.com
*/
func main() {
	// Esto es para dejar disponible el flag en la variable site
	flag.Parse()

	start := time.Now()
	var wg sync.WaitGroup

	/*
		Puerto 21 abierto, Time 18.628875ms
		Puerto 25 abierto, Time 18.711208ms
		Puerto 80 abierto, Time 218.522375ms
		Puerto 22 abierto, Time 219.03ms
	*/

	fmt.Println("Ejecutando escaneo de puertos sitio:", *site)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))

			if err != nil {
				// Puerto cerrado
				return
			}

			conn.Close()
			fmt.Printf("Puerto %d abierto, Time %s \n", port, time.Since(start))
		}(i)
	}

	wg.Wait()
}
