package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

	elements := make(map[string]int)
	elements = countElements(elements, doc)
	fmt.Printf("------ number per element ------\n")
	for key, value := range elements {
		fmt.Printf("[%s]=%d\n", key, value)
	}
}

// visitは、n内で見つかったリンクを一つ一つlinksへ追加し、その結果を返します。
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}

	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}

	// for c := n.FirstChild; c != nil; c = c.NextSibling {
	// 	links = visit(links, c)
	// }

	return links
}

func countElements(elements map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		elements[n.Data]++
	}

	if n.FirstChild != nil {
		elements = countElements(elements, n.FirstChild)
	}

	if n.NextSibling != nil {
		elements = countElements(elements, n.NextSibling)
	}

	return elements
}
