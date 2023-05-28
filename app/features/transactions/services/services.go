package services

import (
	"BukaLelang/app/features/transactions"
	"BukaLelang/helper"
	"time"
)

type TransactionService struct {
	r transactions.Repository
}

func New(r transactions.Repository) transactions.Service {
	return &TransactionService{r: r}
} 

func (ts *TransactionService) GetTransaction(transactionid uint) (transactions.Transaction, error) {
	transaction, err := ts.r.GetTransaction(transactionid)
	if err != nil {
		return transactions.Transaction{}, err
	}

	return transaction, nil
}

func (ts *TransactionService) CreateTransaction(user_id, lelang_id, GrandTotal uint, paymentmethod string, request transactions.Transaction) error {
	Transaction := transactions.Transaction{
		UserID:  				user_id,
		LelangID: 				lelang_id,
		Invoice: 				helper.GenerateInvoice(),
		PurchaseStartDate: 		time.Now(),
		PurchaseEndDate: 		time.Now().Add(24 * time.Hour),
		Status: 				"pending",
		StatusDate:             time.Now(),
		Transaction_Bids:       request.Transaction_Bids, 
		GrandTotal:             GrandTotal,
		PaymentMethod:          paymentmethod,
	} 

	err := ts.r.CreateTransaction(Transaction)
	if err != nil {
		return err 
	} 

	return nil 
}