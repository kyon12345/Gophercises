package link

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	filename := flag.String("filename", "ex1.html", "html file to be parse")
	flag.Parse()
	// link := parser.Parse(filename)
	// fmt.Println("%v", link)
	// var link Link
	r, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}

	var f func(*html.Node)
	var findText func(*html.Node)
	link := &Link{"", make([]string, 0)}

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			//get href
			for _, a := range n.Attr {
				if a.Key == "href" {
					link.Href = append(link.Href, a.Val)
					break
				}
			}
			findText(n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	findText = func(n *html.Node) {
		if n.Type == html.TextNode {
			b := []byte(link.Text)
			b = append(b, []byte(strings.TrimSpace(n.Data))...)
			link.Text = string(b)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findText(c)
		}
	}

	f(doc)
	fmt.Printf("%v", link)
}

type Link struct {
	Text string
	Href []string
}
