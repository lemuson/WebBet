package models

type BetStatus struct {
	ID       uint   `json:"id"   gorm:"column:ID_СтатусСтавки; primaryKey"`
	Название string `json:"name" gorm:"column:Название"`
	Bets     []Bet  `            gorm:"foreignKey:IDСтатусСтавки"`
}

func (BetStatus) TableName() string {
	return "СтатусыСтавок"
}
