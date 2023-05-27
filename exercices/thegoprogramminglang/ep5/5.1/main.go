package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	s := os.Args[1]
	resp, err := http.Get(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error encountered: %v \n", err)
	}
	// io.Copy(os.Stdout, resp.Body)
	doc, err := html.Parse(resp.Body)
	// for _, link := range visit(nil, doc) {
	// 	fmt.Println(link)
	// }
	links := []string{}
	links = visit(links, doc)
	for _, v := range links {
		fmt.Println(v)
	}
}
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
