package transactions

import (
	"time"

	"github.com/labstack/echo/v4"
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
	Status        	  string 
	StatusDate        time.Time
	GrandTotal        uint 
	PaymentMethod     string 
	Transaction_Bids  []Transaction_Bids
} 

type Transaction_Bids struct {
	TransactionID        uint 
	BidID                uint 
	BidPrice       		 uint
	BidBuyer       		 string  
	BidQuantity    		 uint 
	Total                uint
} 

type Repository interface {
	CreateTransaction(Transaction) error 
	GetTransaction(transactionid uint) (Transaction, error) 
} 

type Service interface {
	CreateTransaction(user_id, lelang_id, GrandTotal uint, paymentmethod string, request Transaction) error 
	GetTransaction(transactionid uint) (Transaction, error)	
} 

type Handler interface {
	CreateTransaction() echo.HandlerFunc
	GetTransaction()    echo.HandlerFunc
}