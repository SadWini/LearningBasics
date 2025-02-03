package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	/*
		if err := Ping(); err != nil {
			log.Printf("ping failed: %v; networking disabled", err)
		}
		if err := Ping; err != nil{
			log.Fprintf(os.Stderr, "ping failed: %v; networking disabled", err)
		}
	*/
	dir, err := os.MkdirTemp("", "scratch")
	if err != nil {
		fmt.Errorf("failed to create temp dir: %v", err)
	}
	os.RemoveAll(dir) // ignore errors; $TMPDIR is cleaned periodically
	if err := WaitForServer(os.Args[1]); err != nil {
		// fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		log.Fatalf("Site is down %v\n", err)
		os.Exit(1)
	}
}

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	log.SetPrefix("wait: ")
	log.SetFlags(0)
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err != nil {
			return nil
		}
		log.Printf("server not respongind (%s); retrying....", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
