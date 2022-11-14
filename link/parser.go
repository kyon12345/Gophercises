package link

import (
	"io"

	"golang.org/x/net/html"
)

func Parse(raw io.Reader) ([]string, error) {
	doc, err := html.Parse(raw)
	if err != nil {
		panic(err)
	}
	return dfs(doc), nil
}

func dfs(root *html.Node) []string {
	res := []string{}
	if root.Type == html.ElementNode && root.Data == "a" {
		for _, v := range root.Attr {
			if v.Key == "href" {
				res = append(res, v.Val)
				break
			}
		}
	}
	for c := root.FirstChild; c != nil; c = c.NextSibling {
		dfs(c)
	}
	return res
}
