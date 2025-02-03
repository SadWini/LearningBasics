package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}
func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	var depth int
	startElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			var attribute string
			for _, attr := range n.Attr {
				attribute += fmt.Sprintf("%s=%q ", attr.Key, attr.Val)
			}
			child := ""
			if n.Data == "img" && n.FirstChild == nil {
				child = "/"
			}
			if len(attribute) > 1 {
				fmt.Printf("%*s<%s %s%s>\n", depth*2, "", n.Data, attribute, child)
			} else {
				fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, child)
			}
			depth++
		} else if n.Type == html.TextNode || n.Type == html.CommentNode {
			if !(n.Parent.Type == html.ElementNode &&
				(n.Parent.Data == "script" || n.Parent.Data == "style")) {
				for _, line := range strings.Split(n.Data, "\n") {
					line = strings.TrimSpace(line)
					if line != "" && line != "\n" {
						fmt.Printf("%*s%s\n", depth*2, "", line)
					}
				}
			}
		}
	}
	endElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			if !(n.Data == "img" && n.FirstChild == nil) {
				fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
			}
		}
	}
	forEachNode(doc, startElement, endElement)
	return nil
}

func ElementById(doc *html.Node, id string) *html.Node {
	pre := func(doc *html.Node) bool {
		for _, attr := range doc.Attr {
			if attr.Key == "id" && attr.Val == id {
				return true
			}
		}
		return false
	}
	if node, found := forEachNodeSearch(doc, pre, nil); found {
		return node
	}
	return nil
}

func forEachNodeSearch(n *html.Node, pre, post func(n *html.Node) bool) (*html.Node, bool) {
	if pre != nil {
		if pre(n) {
			return n, true
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if node, ok := forEachNodeSearch(c, pre, post); ok {
			return node, true
		}

	}
	if post != nil {
		if post(n) {
			return n, true
		}
	}
	return nil, false
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
