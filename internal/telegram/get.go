package telegram

func (b *botWork) GetNewMessage() {
	// // Get запрос на сервер Telegram
	// resp, err := http.Get("https://api.telegram.org/bot" + b.config.TGToken + "/getUpdates?offset=" + strconv.Itoa(idU))
	// if err != nil {
	// 	return nil, errors.New("Ошибка GET запроса к Telegram: " + err.Error())
	// }
	// // закрываем по завершении
	// defer resp.Body.Close()

	// // читаем json в байтах
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, errors.New("Ошибка чтения ответа Telegram: " + err.Error())

	// }
}
