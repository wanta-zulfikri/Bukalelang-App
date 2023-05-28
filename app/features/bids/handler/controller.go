package handler

import (
	"BukaLelang/app/features/bids"
	"BukaLelang/helper"
	"BukaLelang/middlewares"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BidController struct {
	s bids.Service
}

func New(h bids.Service) bids.Handler {
	return &BidController{s: h}
}

func (tc *BidController) GetBids() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization") 
		_, err := middlewares.ValidateJWT2(tokenString)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized,"Missing or Malformed JWT. "+err.Error(), nil))
		}

		inputID := c.Param("id")
		if inputID == "" {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		} 

		id , err := strconv.ParseUint(inputID, 10, 32)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found,", nil))
		} 

		bids, err := tc.s.GetBids(uint(id)) 
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
		} 

		var response []ResponseGetBids 
		for _, bid := range bids {
			response = append(response, ResponseGetBids{
				LelangID: bid.LelangID,
				BidPrice: uint(bid.BidPrice),
				BidBuyer: bid.BidBuyer,
				BidQuantity: uint(bid.BidQuantity),
			})
		}

		return c.JSON(http.StatusOK, helper.DataResponse{
			Code: http.StatusOK,
			Message: "Successful operation.", 
			Data: response,
		})
	}
} 

func (tc *BidController) UpdateBids()echo.HandlerFunc {
	return func(c echo.Context) error {
		var input []RequestUpdateBid 
		tokenString := c.Request().Header.Get("Authorization")
		_, err := middlewares.ValidateJWT2(tokenString) 
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT. "+err.Error(), nil))
		} 

		lelangid , err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		} 

		if err := c.Bind(&input); err != nil {
			c.Logger().Error("Failed to bind input: ", err) 
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		} 

		var updatedBids []bids.Core 
		for _, input := range input {
			updatedBids = append(updatedBids, bids.Core{
				LelangID:    uint(lelangid),
				BidPrice:    uint(input.BidPrice), 
				BidBuyer:    input.BidBuyer,
				BidQuantity: uint(input.BidQuantity),

			})
		}

		err = tc.s.UpdateBids(uint(lelangid), updatedBids)
		if err != nil {
			c.Logger().Error("Failed to update bids: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		} 
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Success Updated Bids", nil)) 
	}
}

func (tc *BidController) DeleteBids() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization") 
		_, err := middlewares.ValidateJWT2(tokenString) 
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT. " +err.Error(),nil)) 

		}
		inputID := c.Param("id")
		if inputID == "" {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		} 

		id , err := strconv.ParseUint(inputID,10,32) 
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		} 

		err = tc.s.DeleteBids(uint(id)) 
		if err != nil {
			c.Logger().Error("Error deleting profile", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Success deleted a bids", nil))
	}
}