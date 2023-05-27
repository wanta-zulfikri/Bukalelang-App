package repository

import (
	"BukaLelang/app/features/bids/repository"
	"BukaLelang/app/features/lelangs"
	"errors"
	"time"

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
			BidPrice     : bid.BidPrice,
			BidBuyer     : bid.BidBuyer,
			BidQuantity  : bid.BidQuantity,
			LelangID     : newLelang.ID,
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

func (er *LelangRepository) GetLelangsByCategory(category string) ([]lelangs.Core, error) {
	var cores []lelangs.Core 
	if err := er.db.Table("lelangs").Where("category = ? AND deleted_at IS NULL", category).Find(&cores).Error; err != nil {
		return nil, err
	}
	return cores, nil
}

func (er *LelangRepository) GetLelangsByUserID(userid uint) ([]lelangs.Core, error) {
	var cores []lelangs.Core
	if err := er.db.Table("lelangs").Where("user_id = ? AND deleted_at IS NULL", userid).Find(&cores).Error; err != nil {
		return nil, err 
	}
	return cores, nil 
} 

func (er *LelangRepository) GetLelang(lelangid uint) (lelangs.Core, error) {
	var input Lelang 
	result := er.db.Where("id = ? AND deleted_at IS NULL", lelangid).Find(&input)
	if result.Error != nil {
		return lelangs.Core{}, result.Error	
	}
	if result.RowsAffected == 0 { 
		return lelangs.Core{}, result.Error
	}
		return lelangs.Core{
			ID: input.ID, 
			Item: input.Item,
			Deskripsi: input.Deskripsi,
			Price: input.Price,
			Seller: input.Seller,
			Date: input.Date,
			Status: input.Status,
			Time: input.Time,
			Image: input.Image,
		}, nil
}

func (er *LelangRepository) UpdateLelang(id uint, updatedLelang lelangs.Core) error {
	if err := er.db.Model(&Lelang{}).Where("id = ?", id).Updates(map[string]interface{}{
		"item"			: updatedLelang.Item, 
		"deskripsi"		: updatedLelang.Deskripsi,
		"price"			: updatedLelang.Price,
		"seller"        : updatedLelang.Seller,
		"date"	        : updatedLelang.Date, 
		"status"        : updatedLelang.Status, 
		"time"	        : updatedLelang.Time, 
		"image"         : updatedLelang.Image,
		"updated_at"    : time.Now(),
	}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err 
		}
		return err 
	}
	return nil 
} 

func (er *LelangRepository) DeleteLelang(id uint) error {
	input := Lelang{}
	if err := er.db.Where("id = ?", id).Find(&input).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}
	input.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid:true}

	if err := er.db.Save(&input).Error; err != nil {
		return err
	}
	return nil 
}