package historys

import "github.com/labstack/echo/v4"

type Core struct {
	LelangID   uint
	UserID     uint
	Item       string
	StatusItem string
	PriceSold  int64
	Buyer      string 
	Seller     string
}

type Repository interface {
	CreateHistory(Core) (Core, error)
	UpdateHistory(Core) (Core, error)
	DeleteHistory(id uint) error
}

type Service interface {
	CreateHistory(Core) (Core, error)
	UpdateHistory(Core) (Core, error)
	DeleteHistory(id uint) error
}

type Handler interface {
	CreateHistory() echo.HandlerFunc
	UpdateHistory() echo.HandlerFunc
	DeleteHistory() echo.HandlerFunc
}