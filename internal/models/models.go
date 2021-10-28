package models

// одна заявка
type Apartment struct {
	Number       int        // Номер
	NumberRoom   int        // Квартира
	TimeCreate   string     // Время создания
	Priority     string     // Приоритет
	Status       string     // Статус
	TimeTransfer string     // Передана исполнителю
	TimeFinish   string     // Выполнена
	DetailCase   DetailCase // Детали заявки
}

// детали одной заявки
type DetailCase struct {
	ApplicantAssessment      string // Оценка заявителя
	Cause                    string // Причина
	ApplicationType          string // Тип заявки
	WorkGroup                string // Группа работ
	ControlClaim             string // Контрольная заявка
	MES                      string // На контроле в МЧС
	ImplementingOrganization string // Организация-исполнитель
	Work                     string // Работы
}

// сохранение последней заявки
type SaveLastClaim struct {
	HomeID  int // id дома
	ClaimID int // id последней заявки
}
