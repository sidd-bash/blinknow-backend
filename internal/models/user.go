package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Phone             string `gorm:"uniqueIndex;not null" json:"phone"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	Address           string `json:"address"`
	IsProfileComplete bool   `gorm:"default:false" json:"is_profile_complete"`
}
