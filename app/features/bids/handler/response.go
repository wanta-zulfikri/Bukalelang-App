package handler 


type ResponseGetBids struct {
	LelangID       uint 	`json:"lelang_id"`
	BidPrice       uint		`json:"bid_price"`
	BidBuyer       string  	`json:"bid_buyer"`
	BidQuantity    uint     `json:"bid_quantity"`
} 

type BidResponse struct {
	Code 		int     			`json:"code"`    
	Message     string  			`json:"message"`
	Data        []ResponseGetBids   `json:"data"`
} 

type Pagination struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
}