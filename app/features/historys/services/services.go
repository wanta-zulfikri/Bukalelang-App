package services

import (
	"BukaLelang/app/features/historys"
	"errors"

	
	"gorm.io/gorm"
)

type HistoryService struct {
	n historys.Repository
}

func New(o historys.Repository) historys.Service {
	return &HistoryService{n: o}
}

func (rs *HistoryService) CreateHistory(request historys.Core) (historys.Core, error) {
	result, err := rs.n.CreateHistory(request)
	if err != nil {
		return request, errors.New(err.Error())
	}
	return result, nil
}


func (rs *HistoryService) UpdateHistory(request historys.Core) (historys.Core, error) {
	result, err := rs.n.UpdateHistory(request)
	if err != nil {
		return request, errors.New(err.Error())
	}
	return result, nil
}

func (rs *HistoryService) DeleteHistory(id uint) error {
	err := rs.n.DeleteHistory(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}
	return nil
}