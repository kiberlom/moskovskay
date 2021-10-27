package models

type Apartment struct {
	Number       string     // Номер
	NumberRoom   string     // Квартира
	TimeCreate   string     // Время создания
	Priority     string     // Приоритет
	Status       string     // Статус
	TimeTransfer string     // Передана исполнителю
	TimeFinish   string     // Выполнена
	DetailCase   DetailCase // Детали заявки
}

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
