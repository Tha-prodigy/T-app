package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println("Error!, couldn't get a listener object")
		return
	}
	defer listener.Close()
	fmt.Println("Listening on port :8080")
	for {

		conn, err2 := listener.Accept()
		if err2 != nil {
			log.Println(err2)
			return
		}
		HandleConnection(conn)

	}

}
