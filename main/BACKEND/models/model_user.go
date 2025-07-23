package models

type User struct {
	ID                   uint      `json:"id"            gorm:"column:ID_Пользователь;primaryKey"`
	Логин                string    `json:"login"         gorm:"column:Логин"`
	Пароль               string    `json:"password"      gorm:"column:Пароль"`
	IDРоль               uint      `json:"role_id"       gorm:"column:ID_Роль"`
	IDДанныеПользователя *uint     `json:"user_data_id"  gorm:"column:ID_ДанныеПользователя"`
	ДанныеПользователя   *UserData `json:"userData"      gorm:"foreignKey:IDДанныеПользователя"`
	Роль                 Role      `                     gorm:"foreignKey:IDРоль"`
	Ставки               []Bet     `                     gorm:"foreignKey:IDПользователь"`
}

// JSON ответа
type UserResponse struct {
	ID     uint             `json:"id"`
	Логин  string           `json:"login"`
	Данные UserDataResponse `json:"userData"`
	Ставки []BetResponse    `json:"bets"`
}

type LoginInput struct {
	Логин  string `json:"login"`
	Пароль string `json:"password"`
}

func (User) TableName() string {
	return "Пользователи"
}
