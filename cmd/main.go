package main

import (
	"log"

	// "github.com/kiberlom/moskovskay/internal/ais/get"
	// "github.com/kiberlom/moskovskay/internal/ais/parse"

	"github.com/kiberlom/moskovskay/internal/config"
	"github.com/kiberlom/moskovskay/internal/datebase"
)

func main() {

	//wg := &sync.WaitGroup{}

	_, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err = datebase.NewConnect(); err != nil {
		log.Fatal(err)
	}

	// 	bot := telegram.NewBot(conf)

	// 	wg.Add(1)
	// 	ais.NewAISBot(&ais.ConfigAIS{
	// 		Conf:  conf,
	// 		TGBot: bot,
	// 		WG:    wg,
	// 	})
	// wg.Wait()

	// b, err := json.Marshal(parse)
	// if err != nil {
	// 	log.Panicln(err)
	// }

	// fmt.Println(string(b))

}
