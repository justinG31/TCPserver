package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	fmt.Println("Inside function")
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {

		// will listen for message to process ending in newline (\n)
		message := scanner.Text()
		// output message received
		fmt.Print("Message Received:", message)

		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}

}

func main() {
	fmt.Println("Launching server...")
	fmt.Println("Listen on port")
	ln, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Accept connection on port")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Calling handleConnection")
		go handleConnection(conn)
	}

}
