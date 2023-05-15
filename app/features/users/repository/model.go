package repository

import "gorm.io/gorm"

type User struct {
	gorm.Model 
	Username 	string		`json:"username" gorm:"type:varchar(100);not null"`
	Email 		string 		`json:"email" gorm:"type:varchar(100);not null"`
	Phone 		string 		`json:"phone" gorm:"type:varchar(15);not null"`
	Password 	string      `json:"password" gorm:"type:varchar(100);not null"`
	Image 		string 		`json:"image" gorm:"type:varchar(100)"`
} 

