package models

import "gorm.io/gorm"

type Product struct {
  gorm.Model
  ID          uint     `gorm:"primaryKey"`
  Name        string   `gorm:"not null"`
  Price       float64  `gorm:"not null"`
  ImageURL    string
  CategoryID  uint
  Category    Category
}