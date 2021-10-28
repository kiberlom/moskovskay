package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/kiberlom/moskovskay/internal/config"
)

type botWork struct {
	config *config.Config
}

func requestToJson(text string) ([]byte, error) {
	s := SendMessageRequest{
		ChatID: -1001748587069,
		Text:   text,
	}

	b, err := json.Marshal(s)
	if err != nil {
		return nil, fmt.Errorf("Ошибка подготовки сообщения для отправки (Marshal): %v", err)
	}

	return b, nil

}

func (b *botWork) Send(text string) error {

	messageByte, err := requestToJson(text)
	if err != nil {
		return err
	}

	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	r, err := http.NewRequest("POST", "https://api.telegram.org/bot"+b.config.TGToken+"/sendMessage", bytes.NewBuffer([]byte(messageByte)))
	if err != nil {
		return fmt.Errorf("Ошибка NewRequest: %v", err)
	}
	//r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Content-Type", "application/json; charset=utf-8")

	// делаем запрос
	resp, err := client.Do(r)
	if err != nil {
		return fmt.Errorf("Ошибка client.DO: %v", err)
	}

	// читаем ответ
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Ошибка чтения тела ответа: %v", err)
	}

	fmt.Println("-----------------------------------------------------------------------------")
	fmt.Println(string(body))
	//fmt.Printf("%+v", u)
	fmt.Println("------------------------------------------------------------------------------")

	return nil

}
