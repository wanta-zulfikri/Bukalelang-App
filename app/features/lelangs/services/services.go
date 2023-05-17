package services

import "BukaLelang/app/features/lelangs"

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

func (es *LelangService) GetLelangs() ([]lelangs.Core, error) {
	lelangs, err := es.r.GetLelangs()
	if err != nil {
		return nil, err
	}
	return lelangs, nil
}