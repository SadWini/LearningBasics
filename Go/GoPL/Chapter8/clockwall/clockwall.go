package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	for _, locate := range os.Args[2:] {
		s := strings.Split(locate, "=")
		go getTime(s[1])
	}

	s := strings.Split(os.Args[1], "=")
	getTime(s[1])
}

func mustCopy(w io.Writer, r io.Reader) {
	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}
}

func getTime(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	mustCopy(os.Stdout, conn)
}
