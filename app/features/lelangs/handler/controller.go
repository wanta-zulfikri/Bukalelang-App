package handler

import (
	"BukaLelang/app/features/lelangs"
	"BukaLelang/helper"
	"BukaLelang/middlewares"
	"math"
	"net/http"
	"strconv"

	"github.com/Kong/go-pdk/service/response"
	"github.com/labstack/echo/v4"
)

type LelangController struct {
	s lelangs.Service 
}

func New(h lelangs.Service) lelangs.Handler {
	return &LelangController{s:h}
} 

func (ec *LelangController) CreateLelangWithBid() echo.HandlerFunc {
	return func (c echo.Context) error {
		var input RequestCreateLelangWithBid
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middlewares.ValidateJWT2(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT"+err.Error(), nil))
		}

		id       := claims.ID
		username := claims.Username 

		if err := c.Bind(&input); err != nil {
			c.Logger().Error("Failed to bind input: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		// masuk service 
		lelangBids := make([]lelangs.BidCore, len(input.Bids))
		for i, bid := range input.Bids {
			lelangBids[i] = lelangs.BidCore{
				BidPrice   : bid.BidPrice,
				BidBuyer   : bid.BidBuyer,
				BidQuantity: bid.BidQuantity,
			}
		}

		newLelang := lelangs.Core{
			Item       : input.Item,       
			Deskripsi  : input.Deskripsi,  
			Price      : input.Price,  
			Seller     : username,  
			Date       : input.Date,  
			Status     : input.Status,  
			Time       : input.Time,  
			Image      : input.Image,     
			Bid       : lelangBids,  
		} 

		err = ec.s.CreateLelangWithBid(newLelang, id) 
		if err != nil {
			c.Logger().Error("Failed to create lelang with bids: ",err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		} 

		response := LelangResponse{
			Code: http.StatusCreated,
			Mesaage: "Success Created an lelang",
			Data: LelangData{
				Item		: newLelang.Item,
				Deskripsi	: newLelang.Deskripsi,
				Price		: newLelang.Price,
				Seller		: newLelang.Seller,
				Date        : newLelang.Date,
				Status		: newLelang.Status,
				Time        : newLelang.Time,
				Image		: newLelang.Image,
				Bids	    : make([]BidResponse, len(newLelang.Bid)),

			},
		}

		for i, bid := range newLelang.Bid {
			response.Data.Bids[i] = BidResponse{
				Price		: bid.BidPrice,
				Buyer		: bid.BidBuyer,
				Quantity	: bid.BidQuantity,
			}
		}

		return c.JSON(http.StatusCreated, response)

	}
}

func (ec *LelangController) GetLelangs() echo.HandlerFunc {
	return func(c echo.Context) error { 
	category := c.QueryParam("category")
	var lelangs []lelangs.Core 
	var err error

	if category == "" {
		lelangs, err = ec.s.GetLelangsByCategory(category)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
		}
	} else {
		lelangs, err = ec.s.GatLelangs()
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
		}
	}

	if len(lelangs) == 0 {
		if err != nil {
			c.Logger().Error(err.Error())	
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
		}
	} 

	formattedLelangs := []ResponseGetLelangs{}
	for _, lelang := range lelangs {
		formattedLelang := ResponseGetLelangs{
			ID           : lelang.ID,
			Item         : lelang.Item,
			Deskripsi    : lelang.Deskripsi,
			Price        : lelang.Price,
			Seller       : lelang.Seller,
			Date         : lelang.Date,
			Status       : lelang.Status,
			Time         : lelang.Time,
			Image        : lelang.Image,
		}
		formattedLelangs = append(formattedLelangs, formattedLelang)
	} 

	page := c.QueryParam("page")
	perPage := c.QueryParam("per_page") 
	if page != "" || perPage == "" {
		perPage = "3"
	}
	pageInt := 1 
	if page != "" {
		pageInt, _ = strconv.Atoi(page)
	}
	perPageInt, _ := strconv.Atoi(perPage)

	total := len(formattedLelangs)
	totalPages := int(math.Ceil(float64(total) / float64(perPageInt))) 

	startIndex := (pageInt - 1) * perPageInt 
	endIndex := startIndex + perPageInt 
	if endIndex > total {
		endIndex = total
	}

	response := formattedLelangs[startIndex:endIndex]

	pages := Pagination{
		Page      : pageInt,
		Perpage   : perPageInt,
		TotalPages: totalPages,
		TotalItems: total,
	}  

	return c.JSON(http.StatusOK, LelangsResponse{
		Code		: http.StatusOK, 
		Message		: "Successful operation",
		Data		:   response,
		Pagination	: pages,
	})
  }
} 


func (ec *LelangController) GetLelangsByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middlewares.ValidateJWT2(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT. "+err.Error(), nil))
		}

		userid := claims.ID
		lelangs, err := ec.s.GetLelangsByUserID(userid)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound,"The requested resource was not found.", nil))
		}

		if len(lelangs) == 0 {
			if err != nil {
				c.Logger().Error(err.Error())
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil)) 
			}
		}

		formattedLelangs := []ResponseGetLelangs{}
		for _, lelang := range lelangs {
			formattedLelang := ResponseGetLelangs{
				ID           : lelang.ID,
				Item         : lelang.Item,
				Deskripsi    : lelang.Deskripsi,
				Price        : lelang.Price,
				Seller       : lelang.Seller,
				Date         : lelang.Date,
				Status       : lelang.Status,
				Time         : lelang.Time,
				Image        : lelang.Image,
			} 
			formattedLelangs = append(formattedLelangs, formattedLelang)
		}

		page := c.QueryParam("page")
		perPage := c.QueryParam("per_page")
		if page != "" || perPage == "" {
			perPage = "3"
		}
		pageInt := 1 
		if page != "" {
			pageInt, _ = strconv.Atoi(page)
		}
		perPageInt, _ := strconv.Atoi(perPage)

		total := len(formattedLelangs) 
		totalPages := int(math.Ceil(float64(total) / float64(perPageInt))) 

		startIndex := (pageInt - 1) * perPageInt 
		endIndex   := startIndex + perPageInt 
		if endIndex > total {
			endIndex = total
		}

		response := formattedLelangs[startIndex:endIndex] 

		pages := Pagination{
			Page         :   pageInt,
			Perpage      :   perPageInt,
			TotalPages   :   totalPages,
			TotalItems   :   total,

		}

		return c.JSON(http.StatusOK, LelangsResponse{
			Code       : http.StatusOK,
			Message    : "Successful operation.",
			Data       : response,
			Pagination : pages,
		})
	}
}

func (ec *LelangController) GetLelang() echo.HandlerFunc {
	return func(c echo.Context) error {
		lelangid , err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.Logger().Error("Failed to parse ID from URL param: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		lelang, err := ec.s.GetLelang(uint(lelangid))
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound,"The requested resourse was not found.", nil ))
		} 

		response := ResponseGetLelangs{
			ID			: 	lelang.ID,
			Item		:	lelang.Item,
			Deskripsi	: 	lelang.Deskripsi,
			Price		: 	lelang.Price,
			Seller		: 	lelang.Seller,
			Date        :   lelang.Date,
			Status		:   lelang.Status,
			Time		:   lelang.Time,
			Image		: 	lelang.Image,
		}

		return c.JSON(http.StatusOK, helper.DataResponse{
			Code		: http.StatusOK,
			Message		: "Successful operation",
			Data		: response,
		})
	}
} 

func (ec *LelangController) UpdateLelang() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RequestUpdateLelang 
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middlewares.ValidateJWT2(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT."+err.Error(), nil))
		}

		username := claims.Username 
		id := claims.ID 
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("Failed to bind input from request body: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		} 

		file, err := c.FormFile("image")
		var event_picture string
	}
}