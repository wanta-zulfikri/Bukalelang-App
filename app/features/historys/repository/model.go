package repository

import (
	"time"

	"gorm.io/gorm"
)

type History struct {
	gorm.Model 
	LelangID   uint
	UserID     uint
	Item       string
	StatusItem string
	PriceSold  int64 
	Buyer      string
	Seller     string
	UpdatedAt time.Time
}