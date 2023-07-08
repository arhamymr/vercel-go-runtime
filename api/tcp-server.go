package handler

import (
	"fmt"
	"io"
	"net"
	"net/http"
)

func TcpHandler(w http.ResponseWriter, r *http.Request) {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Listening on", ln.Addr())

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Handle the connection (read from/write to conn)
	// ...
	buffer := make([]byte, 1024) // Create a buffer to store received data

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed by remote side")
				return
			}
			fmt.Println("Error reading data:", err)
			return
		}

		data := buffer[:n] // Extract the actual data from the buffer

		// Process the received data as per your application's requirements
		
		fmt.Println("Received data:", string(data))
		conn.Write([]byte("ok i received the data :" + string(data)))
	}
}
