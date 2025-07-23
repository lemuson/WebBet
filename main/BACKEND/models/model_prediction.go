package models

type Prediction struct {
	ID          uint    `json:"id"           gorm:"column:ID_Прогноз; primaryKey"`
	IDРезультат uint    `json:"result_id"    gorm:"column:ID_Результат"`
	IDМатч      uint    `json:"match_id"     gorm:"column:ID_Матч"`
	Коэффициент float64 `json:"coefficient"  gorm:"column:Коэффициент"`
	Результат   Result  `                    gorm:"foreignKey:IDРезультат"`
	Матч        Match   `                    gorm:"foreignKey:IDМатч"`
	Ставки      []Bet   `                    gorm:"foreignKey:IDПрогноз"`
}

// JSON ответа
type PredictionResponse struct {
	ID          uint    `json:"id"`
	Коэффициент float64 `json:"coefficient"`
	Название    string  `json:"name"`
}

func (Prediction) TableName() string {
	return "Прогнозы"
}
