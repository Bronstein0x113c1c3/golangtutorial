package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var tag_population map[string]int = make(map[string]int)

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
	visit2(&tag_population, doc)
	for k, v := range tag_population {
		fmt.Printf("%v: %v \n", k, v)
	}
}

// func visit(links []string, n *html.Node) []string {
// 	if n.Type == html.ElementNode && n.Data == "a" {
// 		for _, a := range n.Attr {
// 			if a.Key == "href" {
// 				links = append(links, a.Val)
// 			}
// 		}
// 	}
// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		links = visit(links, c)
// 	}
// 	return links
// }

func visit2(tag_population *map[string]int, n *html.Node) {
	if n.Type == html.TextNode {
		(*tag_population)[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit2(tag_population, c)
	}
}
