package main

import (
	"log"

	// "github.com/kiberlom/moskovskay/internal/ais/get"
	// "github.com/kiberlom/moskovskay/internal/ais/parse"
	"github.com/kiberlom/moskovskay/internal/config"
	"github.com/kiberlom/moskovskay/internal/telegram"
)

func main() {

	conf, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	bot := telegram.NewBot(conf)
	bot.Send("Привет")
	bot.Send("Пока")

	// year, month, _ := time.Now().Date()

	// content, err := get.GetContent(int(month), year)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// parse, err := parse.Parsing(content)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if parse == nil {
	// 	log.Fatal("Ничего не найдено")
	// }

	// fmt.Printf("%+v", parse)

	// b, err := json.Marshal(parse)
	// if err != nil {
	// 	log.Panicln(err)
	// }

	// fmt.Println(string(b))

}
