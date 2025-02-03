package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	outgoing chan<- string // an outgoing message channel
	name     string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.outgoing <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			cli.outgoing <- "Currently online:"
			for c := range clients {
				if c != cli {
					cli.outgoing <- c.name
				}
			}

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.outgoing)
		}
	}
}

func handleConn(conn net.Conn) {
	var cli client
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	cli.outgoing = ch
	cli.name = who

	ch <- "You are " + who
	messages <- who + "has arrived"
	entering <- cli

	timeout := time.NewTimer(5 * time.Minute)
	go func() {
		<-timeout.C
		conn.Close()
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		timeout.Reset(5 * time.Minute)
		messages <- who + ": " + input.Text()
	}

	leaving <- cli
	messages <- who + "has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintf(conn, msg)
	}
}
