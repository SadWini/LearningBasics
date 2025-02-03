package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	timeout := time.NewTimer(10 * time.Second)
	text := make(chan string)

	var wg sync.WaitGroup
	go func() {
		for input.Scan() {
			text <- input.Text()
		}
		close(text)
	}()

	for {
		select {
		case t, ok := <-text:
			if ok {
				wg.Add(1)
				timeout.Reset(10 * time.Second)
				go func() {
					defer wg.Done()
					echo(c, t, 1*time.Second)
				}()
			} else {
				wg.Wait()
				c.Close()
				return
			}
		case <-timeout.C:
			timeout.Stop()
			c.Close()
			fmt.Println("disconnect client, no new sounds in the last 10 seconds")
			return
		}
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
