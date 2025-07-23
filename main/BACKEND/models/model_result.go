package models

type Result struct {
	ID          uint         `json:"id"    gorm:"column:ID_Результат; primaryKey"`
	Название    string       `json:"name"  gorm:"column:Название"`
	Matches     []Match      `             gorm:"foreignKey:IDРезультат"`
	Predictions []Prediction `             gorm:"foreignKey:IDРезультат"`
}

type ResultResponse struct {
	ID       uint   `json:"id"`
	Название string `json:"name"`
}

func (Result) TableName() string {
	return "Результаты"
}
