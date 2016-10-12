package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

type stream chan []byte

type StreamReader struct {
	path string
}

func (s *StreamReader) read(noise stream) {
	fmt.Println("reader started for dev: ", s.path)
	fh, err := os.Open(s.path)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	scanner := bufio.NewScanner(fh)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		noise <- scanner.Bytes()
	}
}

type GoRand struct {
	noise stream
}

func (g *GoRand) rng() {
	for {
		i := rand.Intn(255)
		g.noise <- []byte(string(i))
		time.Sleep(time.Duration(rand.Intn(100)) * time.Microsecond)
	}
}

func (g *GoRand) addSource(path string) {
	s := &StreamReader{path}
	go s.read(g.noise)
}

func (g *GoRand) run(path string) {
	listen, err := net.Listen("unix", path)
	if err != nil {
		panic(err)
	}
	defer listen.Close()
	go g.rng()

	fmt.Printf("listening on %s\n", path)
	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		go g.handler(conn)
	}
}

func (g *GoRand) handler(conn net.Conn) {
	var b []byte
	defer conn.Close()
	rand.Seed(int64(time.Now().Nanosecond()))

	for {
		b = <-g.noise
		if rand.Intn(2) == 0 {
			conn.Write(b)
		}
	}
}

func main() {
	g := &GoRand{make(stream)}
	if len(os.Args) > 1 {
		for _, a := range os.Args[1:] {
			g.addSource(a)
		}
	}
	g.run("/tmp/rand.sock")
}
