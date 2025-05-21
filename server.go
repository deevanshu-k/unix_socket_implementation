package main

import (
	"fmt"
	"net"
)

func server(url string) {
	listener, err := net.Listen("unix", url)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			break
		}

		clients[conn] = make(chan string)
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	ch := clients[conn]

	go func() {
		for {
			buf := make([]byte, 1024)
			_, err := conn.Read(buf)
			if err != nil {
				if err.Error() == "EOF" {
					// Client closed the connection
				} else {
					fmt.Println("Error reading from client:", err)
				}
				break
			}
		}
		delete(clients, conn)
		close(ch)
	}()

	for {
		d, ok := <-ch
		if !ok {
			break
		}
		_, err := conn.Write([]byte(d))
		if err != nil {
			fmt.Println("Error writing to client:", err)
			break
		}
	}
}
