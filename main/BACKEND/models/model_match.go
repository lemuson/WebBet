package models

type Match struct {
	ID          uint         `json:"id"          gorm:"column:ID_Матч; primaryKey"`
	Дата        string       `json:"date"        gorm:"column:Дата"`
	IDРезультат *uint        `json:"result_id"   gorm:"column:ID_Результат"`
	IDКоманда1  uint         `json:"team1_id"    gorm:"column:ID_Команда_1"`
	IDКоманда2  uint         `json:"team2_id"    gorm:"column:ID_Команда_2"`
	IDВидСпорта uint         `json:"sport_id"    gorm:"column:ID_ВидСпорта"`
	Прогнозы    []Prediction `json:"predictions" gorm:"foreignKey:IDМатч"`
	Результат   Result       `                   gorm:"foreignKey:IDРезультат"`
	Команда1    Team         `                   gorm:"foreignKey:IDКоманда1"`
	Команда2    Team         `                   gorm:"foreignKey:IDКоманда2"`
	ВидСпорта   Sport        `                   gorm:"foreignKey:IDВидСпорта"`
}

// JSON ответа
type MatchResponse struct {
	ID        uint                 `json:"id"`
	Дата      string               `json:"date"`
	Результат string               `json:"result"`
	Команда1  TeamResponse         `json:"team1"`
	Команда2  TeamResponse         `json:"team2"`
	Прогнозы  []PredictionResponse `json:"predictions"`
}

func (Match) TableName() string {
	return "Матчи"
}
