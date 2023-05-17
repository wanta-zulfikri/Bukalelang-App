package routes

import (
	"BukaLelang/app/features/users"
	"BukaLelang/config/common"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(e *echo.Echo, uc users.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger()) 
	//authentication 
	e.POST("/register", uc.Register())
	e.POST("/login", uc.Login()) 
	//users 
	e.GET("/users", uc.GetProfile(), middleware.JWT([]byte(common.JWTSecret)))
	e.PUT("/users", uc.GetProfile(), middleware.JWT([]byte(common.JWTSecret)))
	e.DELETE("/users", uc.DeleteProfile(), middleware.JWT([]byte(common.JWTSecret)))
}