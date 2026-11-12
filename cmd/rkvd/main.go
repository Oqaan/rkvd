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

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr()
	log.Printf("client connected: %s", addr)

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if errors.Is(err, io.EOF) {
			log.Printf("client disconnected: %s", addr)
			break
		} else if err != nil {
			log.Printf("read from %s: %v", addr, err)
			break
		}
		log.Printf("received from %s: %q", addr, buf[:n])

		// RESP simple string: "+" marks the type, "\r\n" ends the message
		_, err = conn.Write([]byte("+PONG\r\n"))
		if err != nil {
			log.Printf("write to %s: %v", addr, err)
			break
		}
	}
}
