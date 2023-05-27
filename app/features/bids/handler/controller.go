package handler

import (
	"BukaLelang/app/features/bids"
	"BukaLelang/helper"
	"BukaLelang/middlewares"
	"go/token"
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
	}
}