package handler 

type RequestCreateHistory struct {
	Item       string    `json:"item"`
	StatusItem string	 `json:"status_item"`
	PriceSold  int64 	 `json:"price_sold"`
	Buyer      string	 `json:"buyer"`
	Seller     string    `json:"seller"`
}

type RequestUpdateHistory struct {
	Item       string    `json:"item"`
	StatusItem string	 `json:"status_item"`
	PriceSold  int64 	 `json:"price_sold"`
	Buyer      string	 `json:"buyer"`
	Seller     string    `json:"seller"`
}