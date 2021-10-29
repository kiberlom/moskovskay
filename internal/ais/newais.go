package ais

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/kiberlom/moskovskay/internal/config"
	"github.com/kiberlom/moskovskay/internal/telegram"
)

type ConfigAIS struct {
	Conf  *config.Config
	TGBot telegram.BotWorker
	WG    *sync.WaitGroup
}

// регулярное получение данных с сайта
func NewAISBot(conf *ConfigAIS) {
	go conf.botAIS()
}

func (s *ConfigAIS) botAIS() {
	for {
		time.Sleep(10 * time.Second)
		year, month, _ := time.Now().Date()

		content, err := GetContent(int(month), year)
		if err != nil {
			log.Print(err)
			continue
		}

		result, err := Parsing(content)
		if err != nil {
			log.Print(err)
			continue
		}

		if result == nil {
			log.Print("На странице ничего не найдено")
			continue
		}

		fmt.Printf("%+v", result)

		if len(result) > 0 {
			j := result[len(result)-1]
			if err := s.TGBot.Send(j.DetailCase.Cause); err != nil {
				log.Print(err)

			}
		}
	}
}
