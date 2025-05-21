package main

import (
	"fmt"
	"net"
)

func client(url string) {
	conn, err := net.Dial("unix", url)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("Server closed the connection.")
			} else {
				fmt.Println("Error reading from server:", err)
			}
			break
		}
		fmt.Printf("> %s\n", string(buf[:n]))
	}
}
