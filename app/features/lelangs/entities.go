package lelangs

import "github.com/labstack/echo/v4"

type Core struct {
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
	Bid       []BidCore     `gorm:"foreignKey:LelangID"`
	History   []HistoryCore `gorm:"foreignKey:LelangID"`
}

type BidCore struct {
	ID           uint
	LelangID     uint
	BidPrice     int64
	BidBuyer     string
	BidQuantity  int64
}

type HistoryCore struct {
	ID         uint
	LelangID   uint
	UserID     uint
	Image      string
	Item       string
	StatusItem string
	PriceSold  int64
}

type Repository interface {
	CreateLelangWithBid(lelang Core, userID uint) error
	GetLelangs() ([]Core, error)
	GetLelangsByCategory(category string) ([]Core, error)
	GetLelangsByUserID(userid uint) ([]Core, error)
}

type Service interface {
	CreateLelangWithBid(lelang Core, userID uint) error
	GatLelangs() ([]Core, error)
	GetLelangsByCategory(category string) ([]Core, error)
	GetLelangsByUserID(userid uint) ([]Core, error)
}

type Handler interface {
	CreateLelangWithBid() echo.HandlerFunc
	GetLelangs() echo.HandlerFunc
	GetLelangsByUserID() echo.HandlerFunc
}