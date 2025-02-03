package main

import (
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/links"
	"net/http"
	"os"
)

func main() {
	for _, link := range os.Args[1:] {
		resp, err := http.Get(link)
		if err != nil {
			fmt.Printf("%v", err)
		}
		doc, err := html.Parse(resp.Body)
		if err != nil {
			fmt.Printf("parsing %s as HTML: %v", link, err)
		}
		title, err := soleTitle(doc)
		if err != nil {
			fmt.Printf("%v", err)
		} else {
			fmt.Println(title)
		}
	}
}

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic
		case bailout{}:
			// expected panic
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p) // unexpected panic
		}
	}()
	links.ForEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}
