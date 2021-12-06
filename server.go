package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

//handler function
func handleConnection(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {

		message := scanner.Text()

		fmt.Print("Message Received is:", message)

		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}

}

func main() {
	fmt.Println("Server is running")
	time1 := time.Now()
	ln, err := net.Listen("tcp", "127.0.0.1:4040")
	if err != nil {
		log.Fatal(err)
	}

	// run connection function as routine
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)

		// have a timeout so server isnt running too long
		go func() {
			<-time.After(time.Duration(10) * time.Second)
			conn.Close()
		}()

	}

	log.Println(time.Since(time1))
	fmt.Println("Server closed")

}
