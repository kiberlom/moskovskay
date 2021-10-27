package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kiberlom/moskovskay/internal/get"
	"github.com/kiberlom/moskovskay/internal/parse"
)

func main() {
	content, err := get.GetContent(10, 2021)
	if err != nil {
		log.Fatal(err)
	}

	parse, err := parse.Parsing(content)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(parse)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(string(b))

}
