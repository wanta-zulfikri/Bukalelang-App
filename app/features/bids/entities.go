package bids



type Core struct {
	ID          uint 
	LelangID    uint 
	Price       int64 
	Buyer       string 
	Quantity    int64 
}


type LelangCore struct { 
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
	Bids         []Core `gorm:"foreignKey:LelangID"`
}