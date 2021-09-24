package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
	CONN_TYPE = "tcp"
)

func main() {
	listener, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		log.Fatal("Error establishing server")
	}
	defer listener.Close()
	log.Println("Listening on " + CONN_HOST + CONN_PORT)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting: ", err.Error())
		}
		handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
	}
	fmt.Println("Message recieved from the client: ", string(message))
	conn.Write([]byte("Hello gophers I have been for so long" + "\n"))
	conn.Close()
}
