package repository

import (
	bids "BukaLelang/app/features/bids/repository"

	"gorm.io/gorm"
)

type Lelang struct {
	gorm.Model 
	ID           uint       `gorm:"primaryKey; autoIncrement"`
	Item         string     `gorm:"type:varchar(100)"`
	Deskripsi    string     `gorm:"type:varchar(225)"`
	Price        string     `gorm:"type:varchar(50)"`
	Seller       string     `gorm:"type:varchar(20)"`
	Date         string     `gorm:"type:varchar(20)"`
	Status       string     `gorm:"type:varchar(20)"`
	Time         string     `gorm:"type:varchar(50)"`
	Image        string     `gorm:"type:varchar(50)"`
	UserID       uint 
	Bids         []bids.Bid `gorm:"foreignkey:LelangID"`
}