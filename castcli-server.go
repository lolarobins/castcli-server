package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

var host string = "0.0.0.0"
var port uint16 = 9430
var active bool = true

func main() {
	args := os.Args[1:]

	// parse arguments
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-h":
			i++
			host = args[i]
		case "-p":
			i++
			in, err := strconv.Atoi(args[i])

			if err != nil {
				fmt.Printf("ERROR: Could not parse specified port. Using default port: %d\n", port)
			} else {
				port = uint16(in)
			}
		}
	}

	server, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))

	if err != nil {
		fmt.Printf("ERROR: Failed to bind to %s:%d\n", host, port)
		os.Exit(1)
	} else {
		fmt.Printf("INFO: Server bound to %s:%d\n", host, port)

	}

	defer server.Close()

	go Accept(server)

	reader := bufio.NewReader(os.Stdin)

ProgramLoop:
	for {
		input, _ := reader.ReadString('\n')
		input = strings.Replace(strings.ToLower(input), "\n", "", -1)

		switch input {
		case ":q":
			fmt.Println("INFO: Stopping the server")
			break ProgramLoop
		}
	}

}

func Accept(server net.Listener) {
	for {
		client, err := server.Accept()

		if err != nil {
			continue
		}

		go Handle(client)
	}
}

func Handle(client net.Conn) {
	buffer := make([]byte, 2048)
	len, err := client.Read(buffer)

	if err != nil {
		client.Close()
		return
	}

	var request, value string

	fmt.Sscanf(string(buffer[:len]), "%s;%s", request, value)

	switch request {
	case "youtube":

	case "link":

	}

	client.Close()
}
