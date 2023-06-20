package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// var tag_population map[string]int = make(map[string]int)

func main() {
	s := os.Args[1]
	resp, err := http.Get(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error encountered: %v \n", err)
	}
	doc, err := html.Parse(resp.Body)
	// for _, link := range visit(nil, doc) {
	// 	fmt.Println(link)
	// }
	fs, err := os.Create("test.txt")
	visit2(doc, fs)
	fs.Close()

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

func visit2(n *html.Node, fs *os.File) {
	if n.Type == html.TextNode {
		// fmt.Print(n.Data)
		fs.Write([]byte(n.Data))
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {

		if (c.Type == html.ElementNode || c.Type == html.TextNode) && !(c.Data == "style" || c.Data == "script") {
			visit2(c, fs)
		}
	}
}
