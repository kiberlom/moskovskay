package get

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var (
	URL    = "https://ulyanovsk-people.aisgorod.ru/Houses/Requests/248030?month=&year="
	method = "GET"
)

// создание url с нужной датой
func getUrl(month, year int) (*string, error) {

	urlBase, err := url.Parse(URL)
	if err != nil {
		return nil, fmt.Errorf("Ошибка парсинга URL: %v", err)
	}

	q := urlBase.Query()
	q.Set("month", strconv.Itoa(month))
	q.Set("year", strconv.Itoa(year))

	urlBase.RawQuery = q.Encode()

	result := urlBase.String()

	return &result, nil
}

// получение контента с сайта
func GetContent(month, year int) ([]byte, error) {

	url, err := getUrl(month, year)
	if err != nil {
		return nil, err
	}

	client := http.Client{
		Transport: &http.Transport{
			ResponseHeaderTimeout: 5 * time.Second,
		},
	}

	r, err := http.NewRequest(method, *url, nil)
	if err != nil {
		return nil, fmt.Errorf("Ошибка NewRequest: %v", err)
	}

	// r.Header.Set("User-Agent", "bot_owners_0.0.1_alpha")
	r.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36")

	responce, err := client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("Ошибка client.Do: %v", err)
	}

	byteBody, err := io.ReadAll(responce.Body)
	if err != nil {
		return nil, fmt.Errorf("Ошибка чтения тела: %v", err)
	}

	defer responce.Body.Close()

	return byteBody, nil
}
