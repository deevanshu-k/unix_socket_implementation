package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var clients = make(map[net.Conn]chan string)

func main() {
	vars := os.Args
	if len(vars) != 3 {
		fmt.Println("Usage:\n go run main.go server <filename>\n go run main.go client <filename>")
		return
	}

	socket_url := vars[2]

	if vars[1] == "server" {
		os.Remove(socket_url)
		go server(socket_url)

		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Printf("> ")
			text, err := reader.ReadString('\n')
			if err != nil {
				panic("Error reading input")
			}

			text = strings.TrimSpace(text)

			if text == "exit" {
				break
			}

			for _, ch := range clients {
				ch <- text
			}
		}

		os.Remove(socket_url)
	} else if vars[1] == "client" {
		client(socket_url)
	}
}
