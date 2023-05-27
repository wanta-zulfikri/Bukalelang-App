package repository

import (
	"BukaLelang/app/features/bids"
	"BukaLelang/app/features/lelangs"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BidsRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *BidsRepository {
	return &BidsRepository{db: db}
}

func (tr *BidsRepository) GetBids(id uint) ([]bids.Core, error) {
	var cores []bids.Core 
	if err := tr.db.Table("bids").Where("lelang_id = ? AND deleted_at IS NULL",id).Find(&cores).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []bids.Core{}, fmt.Errorf("bid with id %v not found", id)

		}
		return []bids.Core{}, err
	}
	return cores, nil
} 

func (tr *BidsRepository) UpdateBids(lelangid uint , updatedBids []bids.Core) error {
	for _, updupdatedBid := updatedBids {
		if err := tr.db.Table("bids").Where("lelang_id = ? AND bid_category = ?", lelangid, updupdatedBid.bid_category).Updates(map[string]interface{}{
			"bid_price": updupdatedBid.BidPrice, 
			"bid_buyer": updupdatedBid.BidBuyer, 
			"bid_quantity": updupdatedBid.BidQuantity,
			"updated_at":  time.Now(),
		}).Error; err != nil {
			return err
		} 
		return nil
	}
}

func (tr *BidsRepository) DeleteBids(id uint) error {
	input := Bid{}
	if err := tr.db.Where("id = ?", id).Find(&input).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}
	input.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}

	if err := tr.db.Save(&input).Error; err != nil {
		return err 
	}
	return nil 
}