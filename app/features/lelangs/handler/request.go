package handler 


type RequestCreateLelang struct {
	Item         string     `json:"item"`
	Deskripsi    string     `json:"deskripsi"`
	Price        string     `json:"price"`
	Seller       string     `json:"seller"`
	Date         string     `json:"date"`
	Status       string     `json:"status"`
	Time         string     `json:"time"`
	Image        string 	`json:"image"`
} 


type RequestUpdateLelang struct {
	Item         string     `json:"item"`
	Deskripsi    string     `json:"deskripsi"`
	Price        string     `json:"price"`
	Seller       string     `json:"seller"`
	Date         string     `json:"date"`
	Status       string     `json:"status"`
	Time         string     `json:"time"`
	Image        string 	`json:"image"`
} 

type RequestCreateLelangWithBid struct {
	Item         string             `json:"item"`
	Deskripsi    string             `json:"deskripsi"`
	Price        string             `json:"price"`
	Seller       string             `json:"seller"`
	Date         string             `json:"date"`
	Status       string             `json:"status"`
	Time         string             `json:"time"`
	Image        string 	        `json:"image"`
	Bids         []RequestCreateBid `json:"bids"`
} 

type RequestCreateBid struct {
	BidPrice     uint   `json:"bid_price"`
	BidBuyer     string  `json:"bid_buyer"`
	BidQuantity  uint   `json:"bid_quantity"`
}
