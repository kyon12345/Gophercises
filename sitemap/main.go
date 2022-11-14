package main

import (
	"flag"
	"fmt"
	"io"
	"link"
	"net/http"
	"net/url"
	"strings"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func main() {
	rootUrl := flag.String("url", "https://gophercises.com", "root url")
	depth := flag.Int("depth", 10, "max depth")
	flag.Parse()

	//bfs from root
	pages := bfs(*rootUrl, *depth)
	for _, page := range pages {
		fmt.Println(page)
	}
}

func bfs(rootUrl string, depth int) []string {
	seen := make(map[string]struct{})
	q := []string{rootUrl}
	d := 1
	for i := 0; i < len(q) && d < depth; i++ {
		l := len(q)
		url := q[i]

		if _, ok := seen[url]; ok {
			continue
		}
		links := get(url)
		for _, link := range links {
			seen[link] = struct{}{}
		}
		q = append(q, links...)

		if i == len(q)-1 {
			d++
		}

		q = q[l:]
	}

	ret := make([]string, 0, len(seen))
	for url := range seen {
		ret = append(ret, url)
	}
	return ret
}

func get(urlStr string) []string {
	res, err := http.Get(urlStr)
	if err != nil {
		return []string{}
	}
	defer res.Body.Close()
	reqUrl := res.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()
	return filter(hrefs(res.Body, base), withPrefix(base))
}

func hrefs(r io.Reader, base string) []string {
	links, _ := link.Parse(r)
	var ret []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l, "/"):
			ret = append(ret, base+l)
		case strings.HasPrefix(l, "http"):
			ret = append(ret, l)
		}
	}
	return ret
}

func filter(links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, link := range links {
		if keepFn(link) {
			ret = append(ret, link)
		}
	}
	return ret
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}
