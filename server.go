package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"runtime"
	"strings"
	"time"
)

//handler function implemented from net package
func handleConnection(conn net.Conn) {

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Print("Message Received is:", message)
		//close client out if message is END
		endMsg := strings.TrimSpace(string(message))
		if endMsg == "END" {
			break
		}

		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))

	}
	conn.Close()

}

func main() {

	numThreads := loopInput2()
	runtime.GOMAXPROCS(numThreads)

	fmt.Println("Server is running")
	time1 := time.Now()
	// create server with net listen function
	ln, err := net.Listen("tcp", "127.0.0.1:4040")
	if err != nil {
		log.Fatal(err)
		return
	}

	connClients := 0
	// run connection function as routine
	for {
		conn, err := ln.Accept()
		connClients++
		fmt.Println(connClients)
		log.Println(time.Since(time1))
		if err != nil {
			log.Fatal(err)
			return
		}

		go handleConnection(conn)

	}

	log.Println(time.Since(time1))
	fmt.Println("Server closed")

}

func loopInput2() int {
	needInput := true
	input := []int{0}
	for needInput {
		procsNum, updateNeedInput := askInput2()
		input[0] = procsNum
		needInput = updateNeedInput
	}
	return input[0]
}

func askInput2() (int, bool) {
	fmt.Println("This program will setup the server")

	fmt.Println("Input the amount of processors to use")
	var numProcs int
	_, err := fmt.Scanln(&numProcs)
	if err != nil {
		fmt.Println("Invalid number of threads. Try again with an integer")
		return 0, true
	}
	fmt.Println("Starting the server using a GOMAXPROCS number of ", numProcs, " .")
	fmt.Println("----------------------------------------------")
	return numProcs, false

}
