package model

type User struct {
	ID       string `json:"id" validate:"required" gorm:"primaryKey"`
	UserName string `json:"username" validate:"required,min=3,max=20" gorm:"not null; unique"`
	Password string `json:"password" validate:"required,min=6,max=20" gorm:"not null"`
	Role     int    `json:"role" gorm:"not null; default:0"`
}
