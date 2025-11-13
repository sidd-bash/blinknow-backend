package models

import "gorm.io/gorm"

type CartItem struct {
  gorm.Model
  ID        uint `gorm:"primaryKey"`
  UserID    uint
  ProductID uint
  Quantity  int
}