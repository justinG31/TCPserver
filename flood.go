package main

import (
	"fmt"
	"log"
	"net"
	"runtime"
	"time"
)

func main() {

	start := time.Now()
	numClient, numThread := loopInput()

	runtime.GOMAXPROCS(numThread)
	//ping server repeatedly with each routine
	pingChan := make(chan string, numClient)
	for i := 0; i < numClient; i++ {
		go Ping("tcp", "127.0.0.1:4040", pingChan)
	}

	var m string
	for i := 0; i < numClient; i++ {
		m = <-pingChan
		log.Println(i+1, m)
	}

	log.Println(time.Since(start))
	log.Println("flood closed")

}

//create a ping to the specified server
func Ping(proto, addr string, out chan<- string) {
	c, err := net.Dial(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	msg := []byte("0101010101010101010101001010101001010101")
	_, err = c.Write(msg)
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)
	_, err = c.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	out <- string(buf)
}

func loopInput() (int, int) {
	needInput := true
	input := []int{0, 0}
	for needInput {
		clientsNum, procsNum, updateNeedInput := askInput()
		input[0] = clientsNum
		input[1] = procsNum
		needInput = updateNeedInput
	}
	return input[0], input[1]
}

func askInput() (int, int, bool) {
	fmt.Println("This program will provide clients to flood a server")

	fmt.Println("Input  the amount of clients to send to the server")
	var numClient int
	_, err1 := fmt.Scanln(&numClient)
	if err1 != nil {
		fmt.Println("not a valid input, requires an integer ")
		return 0, 0, true
	}

	fmt.Println("Input the amount of processors to use")
	var numProcs int
	_, err2 := fmt.Scanln(&numProcs)
	if err2 != nil {
		fmt.Println("Invalid number of threads. Try again with an integer")
		return 0, 0, true
	}
	fmt.Println("Starting the flood with ", numClient, " clients using a GOMAXPROCS number of ", numProcs, " .")
	fmt.Println("----------------------------------------------")
	return numClient, numProcs, false

}
