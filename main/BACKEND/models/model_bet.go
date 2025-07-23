package models

type Bet struct {
	ID             uint       `json:"id"            gorm:"column:ID_Ставка; primaryKey"`
	IDСтатусСтавки uint       `json:"status_id"     gorm:"column:ID_СтатусСтавки"`
	IDПрогноз      uint       `json:"prediction_id" gorm:"column:ID_Прогноз"`
	IDПользователь uint       `json:"user_id"       gorm:"column:ID_Пользователь"`
	Размер         float64    `json:"amount"        gorm:"column:Размер"`
	Коэффициент    float64    `json:"coefficient"   gorm:"column:Коэффициент"`
	СтатусСтавки   BetStatus  `                     gorm:"foreignKey:IDСтатусСтавки"`
	Прогноз        Prediction `                     gorm:"foreignKey:IDПрогноз"`
	Пользователь   User       `                     gorm:"foreignKey:IDПользователь"`
}

// JSON ответа
type BetResponse struct {
	IDМатч      uint    `json:"matchID"`
	Матч        string  `json:"match"`
	Прогноз     string  `json:"prediction"`
	Размер      float64 `json:"amount"`
	Коэффициент float64 `json:"coefficient"`
	Статус      string  `json:"status"`
}

func (Bet) TableName() string {
	return "Ставки"
}
