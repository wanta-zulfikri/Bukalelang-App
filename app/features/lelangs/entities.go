package lelangs 

type Core struct {
	ID           uint 
	Item         string 
	Deskripsi    string 
	Price        string  
	Seller       string 
	Date         string  
	Status       string  
	Time         string  
	Image        string 
	UserID       uint 
	Bid          []BidCore      `gorm:"foreignKey:LelangID"`
	History      []HistoryCore  `gorm:"foreignKey:LelangID"`
}

type BidCore struct {
	ID          uint 
	LelangID    uint
	Price       int64
	Buyer       string 
	Quantity    int64
} 

type HistoryCore struct { 
	ID           uint 
	LelangID     uint
	UserID       uint 
	Image        string 
	Item         string 
	StatusItem   string 
	PriceSold    int64
}
