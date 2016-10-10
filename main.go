package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

func handler(conn net.Conn) {
	defer conn.Close()
	rand.Seed(int64(time.Now().Nanosecond()))
	for {
		i := rand.Intn(255)
		b := []byte(string(i))
		conn.Write(b)
	}
}

func run(path string) {
	listen, err := net.Listen("unix", path)
	if err != nil {
		panic(err)
	}
	defer listen.Close()
	fmt.Printf("listening on %s\n", path)
	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		go handler(conn)
	}
}

func main() {
	run("/tmp/rand.sock")
}
