package repository

import (
	"BukaLelang/app/features/bids/repository"
	"time"
)

type Transaction struct {
	ID                uint
	UserID            uint
	LelangID          uint
	Invoice           string
	Username          string
	Email             string
	Item              string
	LelangDate        string
	LelangTime        string
	PurchaseStartDate time.Time
	PurchaseEndDate   time.Time
	Statutime         string
	StatusDate        time.Time
	GrandTotal        uint
	PaymentMethod     string
	Transaction_Bids  []Transaction_Bids
} 

type Transaction_Bids struct { 
	ID                   uint 
	TransactionID        uint 
	Transaction          Transaction
	BidID                uint 
	Bid                  repository.Bid 
	BidPrice       		 uint
	BidBuyer       		 string  
	BidQuantity    		 uint 
	Total                uint 
} 
