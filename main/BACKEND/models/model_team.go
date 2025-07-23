package models

type Team struct {
	ID          uint    `json:"id"     gorm:"column:ID_Команда; primaryKey"`
	Название    string  `json:"name"   gorm:"column:Название"`
	Изображение string  `json:"image"  gorm:"column:Изображение"`
	HomeMatches []Match `              gorm:"foreignKey:IDКоманда1"`
	AwayMatches []Match `              gorm:"foreignKey:IDКоманда2"`
}

// JSON ответа
type TeamResponse struct {
	ID          uint   `json:"id"`
	Название    string `json:"name"`
	Изображение string `json:"image"`
}

func (Team) TableName() string {
	return "Команды"
}
