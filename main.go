// Multi-threaded server
// Ref: https://www.youtube.com/watch?v=f9gUFy-9uCM&t=134s

package main

import (
	"log"
	"net"
	"time"
)

func raise_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func do(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	raise_error(err)
	log.Println("processing for client request")
	time.Sleep(8 * time.Second)
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello World!\r\n"))
	log.Println("Finished processing")
	conn.Close()
}
func main() {
	listener, err := net.Listen("tcp", ":1729")
	raise_error(err)

	for {
		// Accept is blocking call
		// Wait for client to connect
		log.Println("Waiting for client to connect")
		conn, err := listener.Accept()
		raise_error(err)
		log.Println("Client connected")

		go do(conn)
	}
}
