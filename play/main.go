package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type URLMapper struct {
	Path string
	URL  string
}

func main() {

	b, err := ioutil.ReadFile("./test.yaml")
	if err != nil {
		log.Fatal("fail to read", err)
	}
	var res []URLMapper
	if err := yaml.Unmarshal(b, &res); err != nil {
		log.Fatal("fail to unmarshal", err)
	}
	fmt.Printf("%v,%d", res, len(res))
}
