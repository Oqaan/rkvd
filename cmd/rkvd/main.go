package main

import (
	"errors"
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	log.Println("listening on :6379")

	conn, err := l.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if errors.Is(err, io.EOF) {
			log.Printf("client disconnected")
			return
		} else if err != nil {
			log.Fatal(err)
		}
		log.Printf("%q", buf[:n])

		// RESP simple string: "+" marks the type, "\r\n" ends the message
		_, err = conn.Write([]byte("+PONG\r\n"))
		if err != nil {
			log.Fatal(err)
		}
	}
}
