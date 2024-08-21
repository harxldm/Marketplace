package model

import "time"

type User struct {
	UserID       uint      `gorm:"primaryKey"`
	Email        string    `gorm:"unique;not null" json:"email"`
	Password     string    `gorm:"not null" json:"password"`
	Rol          string    `gorm:"not null" json:"rol"`
	Created_date time.Time `json:"created_Date"`
}
