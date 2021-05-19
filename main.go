package main

import (
	"log"
	"net"
)

func main() {
	//creates server and runs it
	s := newServer()
	go s.run()

	//if err has something then there was an error. print the error
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Unable to start the server: %s", err.Error())
	}

	defer listener.Close()
	log.Printf("Server started on port 8888")

	//endless loop to accept commands
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Unable to accept connection: %s", err.Error())
			continue
		}
		go s.newClient(conn)
	}
}