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
		fmt.Println(CountWordsAndImages(url))
	}
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing html: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode {
		if n.Data == "style" || n.Data == "script" {
			return
		} else if n.Data == "img" {
			images++
		} else if n.Type == html.TextNode {
			text := strings.TrimSpace(n.Data)
			for _, line := range strings.Split(text, "\n") {
				if line != "" {
					words += len(strings.Split(line, " "))
				}
			}
		}
		var tempWords, tempImages int
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			tempWords, tempImages = countWordsAndImages(c)
			words += tempWords
			images += tempImages
		}
	}
	return
}
