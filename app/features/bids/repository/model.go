package repository

import "gorm.io/gorm"

type Bid struct {
	gorm.Model 
	ID             uint    `gorm:"primaryKey: autoIncrement"`
	LelangID       uint 	`gorm:"reference:ID"`
	BidPrice       uint
	BidBuyer       string  `gorm:"type:varchar(20)"`
	BidQuantity    uint 
}