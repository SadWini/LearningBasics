package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	count := make(map[string]int)
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "elements: %v\n", err)
		os.Exit(1)
	}
	countTags(doc, count)
	for k, v := range count {
		fmt.Println(k, v)
	}
}

func countTags(n *html.Node, count map[string]int) {
	if n.Type == html.ElementNode {
		count[n.Data]++
	}
	if n.FirstChild != nil {
		countTags(n.FirstChild, count)
	}
	if n.NextSibling != nil {
		countTags(n.NextSibling, count)
	}
}
