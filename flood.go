package main

import (
	"log"
	"net"
	"time"
)

//create a ping to the specified server
func Ping(proto, addr string, out chan<- string) {
	c, err := net.Dial(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	msg := []byte("This could be a very long message")
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

func main() {

	start := time.Now()
	//ping server repeatedly with each routine
	ch := make(chan string, 1)
	for i := 0; i < 100; i++ {
		go Ping("tcp", "127.0.0.1:4040", ch)
	}

	var m string
	for i := 0; i < 100; i++ {
		m = <-ch
		log.Println(i+1, m)
	}

	log.Println(time.Since(start))

}