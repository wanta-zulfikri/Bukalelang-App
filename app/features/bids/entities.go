package bids

import "github.com/labstack/echo/v4"

type Core struct {
	ID          uint
	LelangID    uint
	BidPrice    uint
	BidBuyer    string
	BidQuantity uint
}

type LelangCore struct {
	ID        uint
	Item      string
	Deskripsi string
	Price     string
	Seller    string
	Date      string
	Status    string
	Time      string
	Image     string
	UserID    uint
	Bids      []Core `gorm:"foreignKey:LelangID"`
}

type Repository interface {
	GetBids(id uint) ([]Core, error)
	UpdateBids(lelangid uint, updatedBids []Core) error
	DeleteBids(id uint) error
}

type Service interface {
	GetBids(id uint) ([]Core, error)
	UpdateBids(lelangid uint, updatedBids []Core) error
	DeleteBids(id uint) error
}

type Handler interface {
	GetBids()    echo.HandlerFunc
	UpdateBids() echo.HandlerFunc
	DeleteBids() echo.HandlerFunc
}