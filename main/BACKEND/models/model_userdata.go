package models

type UserData struct {
	ID           uint    `json:"id"       gorm:"column:ID_ДанныеПользователя; primaryKey"`
	Имя          string  `json:"name"     gorm:"column:Имя"`
	Телефон      string  `json:"phone"    gorm:"column:Телефон"`
	Баланс       float64 `json:"balance"  gorm:"column:Баланс"`
	Пользователи []User  `                gorm:"foreignKey:IDДанныеПользователя"`
}

// JSON ответа
type UserDataResponse struct {
	Имя     string  `json:"name"`
	Телефон string  `json:"phone"`
	Баланс  float64 `json:"balance"`
}

func (UserData) TableName() string {
	return "ДанныеПользователей"
}
