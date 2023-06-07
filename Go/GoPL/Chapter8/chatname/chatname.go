package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

type client struct {
	outgoing chan<- string // an outgoing message channel
	name     string
}

func main() {
	go broadcaster()
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
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

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	var who string

	// get client name
	for k, v := range r.Form {
		if k == "name" {
			who = v[0]
			break
		}
	}
	if who == "" {
		who = r.RemoteAddr
	}

	hi, ok := w.(http.Hijacker)
	if !ok {
		log.Fatalln("Can't Hijack.")
	}
	conn, _, err := hi.Hijack()
	if err != nil {
		log.Fatalln("Hijack error")
	}

	var cli client
	ch := make(chan string, 20) // outgoing client messages
	go clientWriter(conn, ch)

	cli.outgoing = ch
	cli.name = who

	ch <- "You are " + who
	messages <- who + " has arrived"
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
	// NOTE: ignoring potential errors from input.Err()

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintf(conn, msg)
	}
}
