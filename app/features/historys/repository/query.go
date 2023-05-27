package repository

import (
	"BukaLelang/app/features/historys"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

type HistroyRepository struct {
	db *gorm.DB 
}

func (rr *HistroyRepository) CreateHistory(request historys.Core) (historys.Core, error) {
	input := History{
		LelangID   : request.LelangID,
		Item       : request.Item,
		StatusItem : request.StatusItem,
		PriceSold  : request.PriceSold,
		Buyer	   : request.Buyer,
		Seller     : request.Seller,

	}

	err := rr.db.Table("reviews").Create(&input).Error
	if err != nil {
		log.Println("Error creating new historys: ", err.Error())
		return historys.Core{}, err
	}

	createdHistory := historys.Core{
		Buyer: request.Buyer,
		Seller: request.Seller,
		LelangID: input.LelangID,
		StatusItem: input.StatusItem,
		PriceSold: input.PriceSold,
	}
	return createdHistory, nil
} 

func (rr *HistroyRepository) UpdateHistory(request historys.Core) (historys.Core, error) {
	input := History{
		UserID      : request.UserID,
		Item        : request.Item,
		StatusItem  : request.StatusItem,
		PriceSold   : request.PriceSold,
		Buyer       : request.Buyer, 
		Seller      : request.Seller,
		UpdatedAt: time.Now(),
	}
	input.Seller = request.Seller
	input.LelangID = request.LelangID
	input.StatusItem = request.StatusItem
	input.PriceSold = request.PriceSold
	input.Buyer = request.Buyer

	if err := rr.db.Save(&input).Error; err != nil {
		return historys.Core{}, err
	}

	if err := rr.db.Model(&History{}).Where("id = ? AND deleted_at IS NULL", request.LelangID).Updates(History{Buyer: input.Buyer, Seller: input.Seller, StatusItem: input.StatusItem, PriceSold: input.PriceSold, UpdatedAt: time.Now()}).Error; err != nil {
		log.Println("Error updating history: ", err.Error())
		return historys.Core{}, err
	}

	updatedHistory := historys.Core{
		PriceSold: input.PriceSold,
		Seller: request.Seller,
		Buyer: request.Buyer,
		LelangID: request.LelangID,
		StatusItem: input.StatusItem,
	}
	return updatedHistory, nil

} 


func (rr *HistroyRepository) DeleteHistory(id uint) error {
	input := History{}
	if err := rr.db.Where("id = ?", id).Find(&input).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		} 
		return err 
	}

	input.DeletedAt = gorm.DeletedAt{Time:time.Now(), Valid: true}

	if err := rr.db.Save(&input).Error; err != nil {
		return err
	}
	return nil 
}