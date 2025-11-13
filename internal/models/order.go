package models

import "gorm.io/gorm"

type Order struct {
  gorm.Model
  ID        uint `gorm:"primaryKey"`
  UserID    uint
  Total     float64
  Status    string
}