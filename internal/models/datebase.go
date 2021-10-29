package models

import "time"

type DBChatTelegram struct {
	ID     uint `gorm:"primaryKey,autoIncrement"`
	ChatID int  `gorm:"size:255;index:chat_id,unique"`
}

type DBHomeAIS struct {
	ID          uint `gorm:"primaryKey,autoIncrement"`
	HomeID      int  `gorm:"size:255;index:home_id,unique"`
	FullNameAIS string
}

type DBChatWithHome struct {
	ID     uint `gorm:"primaryKey,autoIncrement"`
	HomeID int  `gorm:"home_id"`
	ChatID int  `gorm:"chat_id"`
}

type DBSendMessageTelegram struct {
	ID       uint      `gorm:"primaryKey,autoIncrement"`
	ChatID   int       `gorm:"chat_id"`
	Text     string    `gorm:"text"`
	DateAdd  time.Time `gorm:"date_add"`
	DateSend time.Time `gorm:"date_send"`
}

type DBClaim struct {
	ID                       uint   `gorm:"primaryKey,autoIncrement"`
	Number                   int    `gorm:"number"`                    // Номер
	NumberRoom               int    `gorm:"number_room"`               // Квартира
	TimeCreate               string `gorm:"time_create"`               // Время создания
	Priority                 string `gorm:"priority"`                  // Приоритет
	Status                   string `gorm:"status"`                    // Статус
	TimeTransfer             string `gorm:"time_transfer"`             // Передана исполнителю
	TimeFinish               string `gorm:"time_finish"`               // Выполнена
	ApplicantAssessment      string `gorm:"applicant_assessment"`      // Оценка заявителя
	Cause                    string `gorm:"cause"`                     // Причина
	ApplicationType          string `gorm:"application_type"`          // Тип заявки
	WorkGroup                string `gorm:"work_group"`                // Группа работ
	ControlClaim             string `gorm:"control_claim"`             // Контрольная заявка
	MES                      string `gorm:"mes"`                       // На контроле в МЧС
	ImplementingOrganization string `gorm:"implementing_organization"` // Организация-исполнитель
	Work                     string `gorm:"work"`                      // Работы
}
