package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	UserID          uint      `gorm:"primaryKey"`
	Email           string    `gorm:"unique;not null" json:"email"`
	Password        string    `gorm:"not null" json:"password"`
	ConfirmPassword string    `gorm:"not null" json:"confirm_password"`
	Rol             string    `gorm:"not null" json:"rol"`
	Created_date    time.Time `json:"created_Date"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Rol      string `json:"rol"`
	UserID   int    `json:"user_id"`
}

type Claim struct {
	Email  string `json:"email"`
	Rol    string `json:"rol"`
	UserID int    `json:"user_id"`
	jwt.StandardClaims
}
