package parser

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type Node struct {
	tag       string
	innerHtml string
	childern  []*Node
}

func NewTree(r io.Reader) Node {
	z := html.NewTokenizer(r)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			
		}
		
	}
}

func (root *Node) Find(tag string) *Node {

}

func Parse() {

}
