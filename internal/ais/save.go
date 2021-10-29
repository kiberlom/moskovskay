package ais

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/kiberlom/moskovskay/internal/models"
)

func Last(claims []models.Apartment) error {

	// последний элемент
	// lastClaim := claims[len(claims)-1]

	// file, err := os.OpenFile(path.Join(".", "data", "last_save.json"), os.O_RDWR|os.O_CREATE, 0755)
	// if err != nil {
	// 	return fmt.Errorf("Ошибка открытия файла для сохранения последней заявки: %v", err)
	// }

	// defer file.Close()

	b, err := os.ReadFile(path.Join(".", ".", "data", "last_save.json"))
	if err != nil {
		return fmt.Errorf("Ошибка открытия файла для сохранения последней заявки: %v", err)
	}

	fl := &models.SaveLastClaim{}

	err = json.Unmarshal(b, fl)
	if err != nil {
		return fmt.Errorf("Ошибка unmarshal файла c сохранеными последними заявками: %v", err)
	}
	return nil

}
