package services

import (
	"BukaLelang/app/features/bids"
	"errors"

	"gorm.io/gorm"
)

type BidsService struct {
	r bids.Repository
}

func New(r bids.Repository) bids.Service {
	return &BidsService{r: r}
}

func (bs *BidsService) GetBids(id uint) ([]bids.Core, error) {
	bids, err := bs.r.GetBids(id)	
	if err != nil {
		return nil, err
	}
	return bids, nil 
}  

func (bs *BidsService) UpdateBids(lelangid uint, updatedBids []bids.Core ) error {
	if err := bs.r.UpdateBids(lelangid, updatedBids); err != nil {
		return err
	}
	return nil 
}

func (bs *BidsService) DeleteBids(id uint) error {
	err := bs.r.DeleteBids(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err 
	}
	return nil 
}