package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type AdventureVM struct {
	Title   string
	Story   []string
	Options []Opt
}

type Opt struct {
	Text string
	Arc  string
}

type Node struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Opt    `json:"options"`
	next    *Node    `json:"-"`
}

func main() {
	var filePath string
	flag.StringVar(&filePath, "f", "./gopher.json", "json location")
	flag.Parse()
	//parse json
	fileByte, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("fail to read file", err)
	}

	var list map[string]Node
	if err := json.Unmarshal(fileByte, &list); err != nil {
		log.Fatal("fail to decode json", err)
	}

	//intro
	intro := list["intro"]
	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		data := AdventureVM{
			Story:   intro.Story,
			Title:   intro.Title,
			Options: intro.Options,
		}
		tmpl.Execute(w, data)
	})

	//parse
	http.HandleFunc("/node/", func(w http.ResponseWriter, r *http.Request) {
		key := strings.TrimPrefix(r.URL.Path, "/node/")
		node := list[key]
		data := AdventureVM{
			Story:   node.Story,
			Title:   node.Title,
			Options: node.Options,
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":9090", nil)
}
