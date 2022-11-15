package main

import (
	"fmt"
	"os"
	"task/command"
	"task/db"
)

func main() {
	must(db.Init())
	must(command.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
