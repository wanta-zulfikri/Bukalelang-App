package repository

import "gorm.io/gorm"

type Bid struct {
	gorm.Model 
	ID          uint    `gorm:"primaryKey: autoIncrement"`
	LelangID    uint 	`gorm:"reference:ID"`
	Price       int64 
	Buyer       string  `gorm:"type:varchar(20)"`
	Quantity    int64 
}