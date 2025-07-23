package models

type Sport struct {
	ID          uint    `json:"id"     gorm:"column:ID_ВидСпорта; primaryKey"`
	Название    string  `json:"name"   gorm:"column:Название"`
	Изображение string  `json:"image"  gorm:"column:Изображение"`
	Matches     []Match `              gorm:"foreignKey:IDВидСпорта"`
}

// JSON ответа
type SportResponse struct {
	ID          uint   `json:"id"`
	Название    string `json:"name"`
	Изображение string `json:"image"`
}

func (Sport) TableName() string {
	return "ВидыСпорта"
}
