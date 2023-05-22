package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "textTag: %v\n")
		os.Exit(1)
	}
	findText(doc)
}

func findText(n *html.Node) {
	if n.Type == html.ElementNode && (n.Data == "style" || n.Data == "script") {
		return
	}
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findText(c)
	}
}
