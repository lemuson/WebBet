package models

type Role struct {
	ID       uint   `json:"id"    gorm:"column:ID_Роль; primaryKey"`
	Название string `json:"name"  gorm:"column:Название"`
	Users    []User `             gorm:"foreignKey:IDРоль"`
}

func (Role) TableName() string {
	return "Роли"
}
