package handler

import (
	"BukaLelang/app/features/historys"
	"BukaLelang/helper"
	"BukaLelang/middlewares"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HistoryController struct {
	n historys.Service
}

func New(o historys.Service) historys.Handler {
	return &HistoryController{n:o}
} 

func (rc *HistoryController) CreateHistory() echo.HandlerFunc {
	return func (c echo.Context)error  {
		var input RequestCreateHistory 
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middlewares.ValidateJWT2(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT"+err.Error(), nil))
		} 

		userid := claims.ID 
		username := claims.Username 
		Lelangid, err := strconv.ParseUint(c.Param("id"), 10,64) 
		if err != nil {
			c.Logger().Error("Failed to parse ID from URL param:", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		} 

		request := historys.Core{
			UserID: userid,
			Buyer: username, 
			LelangID: uint(Lelangid), 
			Item: input.Item,
			Seller: input.Seller,
			PriceSold: input.PriceSold,
			StatusItem: input.StatusItem,
		}
		_, err = rc.n.CreateHistory(request)
		if err != nil {
			c.Logger().Error("Failed to create a history", err)
			return c.JSON(http.StatusInternalServerError,helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		} 
		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Success Created a history", nil))
	}
}

func (rc *HistoryController) UpdateHistory() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RequestUpdateHistory 
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middlewares.ValidateJWT2(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT"+err.Error(), nil))	
		}
		userid := claims.ID
		username := claims.Username
		lelangid, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.Logger().Error("Failed to parse ID from URL param: ", err) 
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		} 

		if err := c.Bind(&input); err != nil {
			c.Logger().Error("Failed to bind input: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		} 

		request := historys.Core{
			UserID: userid,
			Buyer: username, 
			LelangID: uint(lelangid), 
			Item: input.Item,
			Seller: input.Seller,
			PriceSold: input.PriceSold,
			StatusItem: input.StatusItem,
		} 

		_, err = rc.n.UpdateHistory(request)
		if err != nil {
			c.Logger().Error("Failed to update history: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		} 
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, " Success Updated a History", nil))
	}
} 

func (rc *HistoryController) DeleteHistory() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middlewares.ValidateJWT2(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized,"Missing or Malformed JWT"+err.Error(), nil))
		} 

		id, err := strconv.ParseUint(c.Param("id"),10, 32) 
		if err != nil {
			c.Logger().Error("Failed to parse ID from URL param: ", err) 
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest,"Bad Request", nil))
		}

		if claims.ID != uint(id) {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Token is not valid for this user.", nil))
		}

		err = rc.n.DeleteHistory(uint(id)) 
		if err != nil {
			c.Logger().Error("Failed to delete history :", err) 
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Success Deleted a History", nil))
	}
}