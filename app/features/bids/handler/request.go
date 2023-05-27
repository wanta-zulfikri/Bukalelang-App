package handler 


type RequestCreateBid struct {
	BidPrice       uint		`json:"bid_price"`
	BidBuyer       string  	`json:"bid_buyer"`
	BidQuantity    uint 	`json:"bid_quantity"`
}

type RequestUpdateBid struct {
	BidPrice       uint		`json:"bid_price"`
	BidBuyer       string  	`json:"bid_buyer"`
	BidQuantity    uint 	`json:"bid_quantity"`
}