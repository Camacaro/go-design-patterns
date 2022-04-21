package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

// Enviar Mensajes a traves del chat
type Client chan<- string

var (
	incomingClients = make(chan Client)
	leavingClients  = make(chan Client)
	messages        = make(chan string)
)

var (
	host = flag.String("h", "localhost", "host del servidor")
	port = flag.Int("p", 3090, "puerto del servidor")
)

// Client1 -> Server -> HandleConnection(Client1)
func HandleConnection(conn net.Conn) {
	defer conn.Close()
	// Un canal de mensajes para este cliente
	message := make(chan string)
	go MessageWrite(conn, message)

	// client1:2560 -> clientname = client1:2560
	clientName := conn.RemoteAddr().String()

	// Enviar un mensaje al cliente (Este, client1)
	message <- fmt.Sprintf("Welcome to the server, your name %s\n", clientName)

	// Mandar un mensaje a los demas clientes
	messages <- fmt.Sprintf("%s has joined the chat\n", clientName)

	incomingClients <- message

	// Utiliza esto para escribi un nuevo mensaje y enviarselo a los demas
	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		// Enviar el mensaje a los demas clientes
		// Nombre del cliente, Mensaje del cliente
		messages <- fmt.Sprintf("%s: %s", clientName, inputMessage.Text())
	}

	// Si cierra la conexion del scanner (inputMessage), entonces se quiere ir
	// de la sala de chat, cerrar la consola
	leavingClients <- message
	messages <- fmt.Sprintf("%s has left the chat\n", clientName)
}

// Vaya Escribiendo los mensajes que se vayan recibiendo (recibe mensaje)
// Esto es para leer los mensajes que se vaya recibiendo
func MessageWrite(conn net.Conn, messages <-chan string) {
	for msg := range messages {
		fmt.Fprintln(conn, msg)
	}
}
