package main

import (
	"fmt"
	"golang.org/x/net/html/links"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	BreadthFirst(crawl, os.Args[1:])
}

func BreadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				for _, tag := range f(item) {
					worklist = append(worklist, tag)
				}
			}
		}
	}
}

var domain string

func crawl(a string) []string {
	fmt.Println(a)

	if domain == "" {
		p, err := url.Parse(a)
		if err != nil {
			log.Fatalf("crawl %s get : %v", err)
		}
		domain = p.Hostname()
		if strings.HasPrefix(domain, "www.") {
			domain = domain[4:]
		}
		fmt.Printf("Domain: %s \n", domain)
	}

	list, err := links.Extract(a)
	if err != nil {
		log.Print(err)
	}

	out := list[:0]
	for _, link := range list {
		p, err := url.Parse(link)
		if err != nil {
			continue
		}
		if strings.Contains(p.Hostname(), domain) {
			fmt.Println(link)
			out = append(out, link)
		}
	}
	return out
}
