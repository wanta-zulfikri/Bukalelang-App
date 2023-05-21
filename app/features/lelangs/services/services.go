package services

import (
	"BukaLelang/app/features/lelangs"
	"errors"

	"gorm.io/gorm"
)

type LelangService struct {
	r lelangs.Repository
}

func New(r lelangs.Repository) lelangs.Service {
	return &LelangService{r: r}
}

func (s *LelangService) CreateLelangWithBid(lelang lelangs.Core, userID uint) error {
	err := s.r.CreateLelangWithBid(lelang, userID)
	if err != nil {
		return err
	}
	return nil 
} 

func (es *LelangService) GetLelangs()([]lelangs.Core, error) {
	lelangs, err := es.r.GetLelangs()
	if err != nil {
		return nil, err
	}
	return lelangs, nil
} 

func (es *LelangService) GetLelangsByCategory(category string) ([]lelangs.Core, error){
	lelangs, err := es.r.GetLelangsByCategory(category)
	if err != nil {
		return nil, err
	}
	return lelangs, nil
}

func (es *LelangService) GetLelangsByUserID(userid uint) ([]lelangs.Core, error) {
	lelangs, err := es.r.GetLelangsByUserID(userid)
	if err != nil {
		return nil, err 
	}
	return lelangs, nil
} 

func (es *LelangService) GetLelang(lelangid uint) (lelangs.Core, error) {
	lelang, err := es.r.GetLelang(lelangid)
	if err != nil {
		return lelangs.Core{}, err
	}
	return lelang, nil
}

func (es *LelangService) UpdateLelang(id uint, updatedLelang lelangs.Core) error {
	updatedLelang.ID = id 
	if err := es.r.UpdateLelang(id, updatedLelang); err != nil {
		return nil 
	}
	return nil
}

func (es *LelangService) DeleteLelang(id uint) error {
	err := es.r.DeleteLelang(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err 
		}
		return err 
	}
	return nil
}