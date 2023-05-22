package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

var linkTags = map[string]string{
	"a": "href", "img": "src", "script": "src", "link": "href",
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for k, _ := range linkTags {
		for _, link := range visit(k, nil, doc) {
			fmt.Println(link)
		}
	}
}

func visit(target string, links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == target {
		for _, a := range n.Attr {
			if a.Key == linkTags[target] {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(target, links, c)
	}
	return links
}

func visitRecursive(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if n.FirstChild != nil {
		links = visitRecursive(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visitRecursive(links, n.NextSibling)
	}
	return links
}
