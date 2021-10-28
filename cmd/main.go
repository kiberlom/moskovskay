package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/kiberlom/moskovskay/internal/get"
	"github.com/kiberlom/moskovskay/internal/parse"
)

func main() {

	year, month, _ := time.Now().Date()

	content, err := get.GetContent(int(month), year)
	if err != nil {
		log.Fatal(err)
	}

	parse, err := parse.Parsing(content)
	if err != nil {
		log.Fatal(err)
	}

	if parse == nil {
		log.Fatal("Ничего не найдено")
	}

	fmt.Printf("%+v", parse)

	b, err := json.Marshal(parse)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(string(b))

}
