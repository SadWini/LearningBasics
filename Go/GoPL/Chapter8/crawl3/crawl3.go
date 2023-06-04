package main

import (
	"flag"
	"fmt"
	"links/links/links"
	"log"
)

var depth = flag.Int("depth", 1,
	"Only URLs reachable by at most `depth` links will be fetched")

func main() {
	flag.Parse()
	worklist := make(chan []string)
	unseenLinks := make(chan string)
	var step, d int

	step++

	counter := make([]int, *depth+2)
	counter[d] = step
	go func() { worklist <- flag.Args() }()

	seen := make(map[string]bool)

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}
	for ; step > 0; step-- {
		list := <-worklist

		if d > *depth {
			continue
		}
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				step++
				counter[d+1]++

				unseenLinks <- link
			}
		}
		if counter[d]--; counter[d] == 0 {
			d++
		}
	}
	close(unseenLinks)
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
