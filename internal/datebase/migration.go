package datebase

import (
	"fmt"

	"github.com/kiberlom/moskovskay/internal/models"
	"gorm.io/gorm"
)

func migration(conn *gorm.DB) error {

	// таблица зарегистрированных чатов
	err := conn.Migrator().CreateTable(&models.DBChatTelegram{})
	if err != nil {
		return fmt.Errorf("Ошибка создания ChatTelegram: %v", err)
	}

	// таблица уникальных идентификаторов домов
	err = conn.Migrator().CreateTable(&models.DBHomeAIS{})
	if err != nil {
		return fmt.Errorf("Ошибка создания DBHomeAIS: %v", err)
	}

	// таблица очереди сообщений для телеграмма
	err = conn.Migrator().CreateTable(&models.DBSendMessageTelegram{})
	if err != nil {
		return fmt.Errorf("Ошибка создания DBSendMessageTelegram: %v", err)
	}

	// таблица соответствия в какие чаты отправлять какие дома
	err = conn.Migrator().CreateTable(&models.DBChatWithHome{})
	if err != nil {
		return fmt.Errorf("Ошибка создания DBChatWithHome: %v", err)
	}

	// таблица соответствия в какие чаты отправлять какие дома
	err = conn.Migrator().CreateTable(&models.DBClaim{})
	if err != nil {
		return fmt.Errorf("Ошибка создания DBClaim: %v", err)
	}

	// добавление данных
	if err := migrationDate(conn); err != nil {
		return err
	}

	return nil
}

func migrationDate(conn *gorm.DB) error {

	// добавление чата моего
	result := conn.Create(&models.DBChatTelegram{
		ChatID: -1001748587069,
	})
	if result.Error != nil {
		return fmt.Errorf("Ошибка добавления данных DBChatTelegram: %v", result.Error.Error())
	}

	// добавление домов
	result = conn.Create(&models.DBHomeAIS{
		HomeID: 248030,
	})
	if result.Error != nil {
		return fmt.Errorf("Ошибка добавления данных DBChatTelegram: %v", result.Error.Error())
	}

	return nil
}
