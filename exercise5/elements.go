package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("error: at least one argument is required.")
	}

	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("%v", err)
	}

	defer resp.Body.Close()

	node, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}

	tags := os.Args[2:]
	nodeList := ElementsByTagName(node, tags...)

	fmt.Printf("url=%s\ttags=%v\t%d found\n", url, tags, len(nodeList))
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	if &doc == nil {
		return nil
	}

	nodeList := []*html.Node{}
	for n := doc.FirstChild; n != nil; n = n.NextSibling {
		if n.Type == html.ElementNode {
			for _, tag := range name {
				if n.Data == tag {
					nodeList = append(nodeList, n)
				}
			}
		}

		nodeList = append(nodeList, ElementsByTagName(n, name...)...)
	}

	return nodeList
}
