package repository

import (
	"BukaLelang/app/features/transactions"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4/request"
	"gorm.io/gorm"
)

type TransactionRespository struct {
	db *gorm.DB 
} 

func New(db *gorm.DB) *TransactionRespository {
	return &TransactionRespository{db: db}
}

func (tr *TransactionRespository) GetTransaction(transactionid uint) (transactions.Transaction, error) {
	var transaction transactions.Transaction 
	if err := tr.db. 
		Where("transactions.id = ?", transactionid). 
		Preload("Transaction_Bids"). 
		Joins("JOIN users ON users.id = transactions.user_id").
		Joins("JOIN lelangs ON lelangs_id = transactions.lelang_id"). 
		Select("transactions.*, users.username, users.email, lelangs.title, lelangs.lelang_date, lelangs.lelang_time"). 
		First(&transaction).Error; err != nil { 
		if errors.Is(err, gorm.ErrRecordNotFound) {	
			return transactions.Transaction{}, errors.New("Transaction not found")
		}
		return transactions.Transaction{}, fmt.Errorf("Failed to retrieve transaction from database: %w", err)
    }
	return transaction, nil 
} 

func (tr *TransactionRespository) CreateTransaction(request transactions.Transaction) error {
	var err error 
	tx := tr.db.Begin() 

	if err := tx.Model(&request)  
		Create(map[string]interface{}){
			
		}
}