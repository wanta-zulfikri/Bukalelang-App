package repository

import (
	"BukaLelang/app/features/bids/repository"
	"BukaLelang/app/features/lelangs"

	"gorm.io/gorm"
)

type LelangRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *LelangRepository {
	return &LelangRepository{db: db}
}

func (er *LelangRepository) CreateLelangWithBid(lelang lelangs.Core, userID uint) error {
	tx := er.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}() 

	// create lelang 
	newLelang := Lelang {
	Item     :  lelang.Item,
	Deskripsi:  lelang.Deskripsi,
	Price    :  lelang.Price, 
	Seller   :  lelang.Seller,
	Date     :  lelang.Date,
	Status   :  lelang.Status,
	Time     :  lelang.Time,
	Image    :  lelang.Image,
	UserID   :  userID,
	}
	err := tx.Table("lelangs").Create(&newLelang).Error 
	if err != nil {
		tx.Rollback()
		return err 
	}

	bids := make([]repository.Bid, len(lelang.Bid))
	for i, bid := range lelang.Bid {
		bids[i] = repository.Bid{
			Price: bid.Price,
			Buyer: bid.Buyer,
			Quantity: bid.Quantity,
			LelangID: newLelang.ID,
		}
	}
	err = tx.Table("bids").CreateInBatches(bids, len(bids)).Error 
	if err != nil {
		tx.Rollback()
		return err 
	}

	return tx.Commit().Error
} 

func (er *LelangRepository) GetLelangs() ([]lelangs.Core, error) {
	var cores []lelangs.Core
	if err := er.db.Table("lelangs").Where("deleted_at IS NULL").Find(&cores).Error; err != nil {
		return nil, err 
	}
	return cores, nil 
}