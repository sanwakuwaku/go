package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
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

	forEachNode(doc, startElement, endElement)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	// count := 0
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
		// count++
	}

	// fmt.Printf("count=%d\n", count)
	if post != nil {
		post(n)
	}
}

var depth int
var hasChild bool

// func startElement(n *html.Node) {
// 	if n.Type == html.ElementNode {
// 		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
// 		depth++
// 	}
// }

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		var attrStr string
		for _, attr := range n.Attr {
			attrStr += " " + attr.Key + "=" + "\"" + attr.Val + "\""
		}

		tagEndChar := ">"
		if n.FirstChild != nil {
			hasChild = true
		} else {
			tagEndChar = "/>"
			hasChild = false
		}

		fmt.Printf("%*s<%s%s%s\n", depth*2, "", n.Data, attrStr, tagEndChar)
		depth++
	} else if n.Type == html.TextNode && strings.TrimSpace(n.Data) != "" {
		// テキスト
		fmt.Printf("%*s%s\n", depth*2, "", n.Data)
	} else if n.Type == html.CommentNode {
		// コメント
		fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if hasChild {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}
