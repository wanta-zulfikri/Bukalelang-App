package handler 

type ResponseGetLelangs struct {
	ID           uint       `json:"lelang_id"`
	Item         string     `json:"item"`
	Deskripsi    string     `json:"deskripsi"`
	Price        string     `json:"price"`
	Seller       string     `json:"seller"`
	Date         string     `json:"date"`
	Status       string     `json:"status"`
	Time         string     `json:"time"`
	Image        string     `json:"image"`
}

type LelangResponse struct {
	Code 		int 	   `json:"code"` 
	Mesaage 	string 	   `json:"message"`
	Data        LelangData `json:"data"`
} 

type LelangData struct {
	Item         string     `json:"item"`
	Deskripsi    string     `json:"deskripsi"`
	Price        string     `json:"price"`
	Seller       string     `json:"seller"`
	Date         string     `json:"date"`
	Status       string     `json:"status"`
	Time         string     `json:"time"`
	Image        string     `json:"image"` 
	Bids         []BidResponse `json:"bids"`
}

type BidResponse struct {
	Price     int64   `json:"bid_price"`
	Buyer     string  `json:"bid_buyer"`
	Quantity  int64   `json:"bid_quantity"`
} 

type LelangsResponse struct {
	Code 	   int                  `json:"code"`
	Message    string               `json:"message"`
	Data       []ResponseGetLelangs `json:"data"` 
	Pagination Pagination           `json:"pagination"`
} 

type Pagination struct {
	Page 	   int `json:"page"`
	Perpage    int `json:"perpage"` 
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`

}