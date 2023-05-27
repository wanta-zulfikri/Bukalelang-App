package routes

import (
	"BukaLelang/app/features/lelangs"
	"BukaLelang/app/features/users"
	"BukaLelang/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(e *echo.Echo, uc users.Handler, ec lelangs.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger()) 
	//authentication 
	e.POST("/register", uc.Register())
	e.POST("/login", uc.Login()) 
	//users 
	e.GET("/users", uc.GetProfile(), middleware.JWT([]byte(config.JWTKey)))
	e.PUT("/users", uc.GetProfile(), middleware.JWT([]byte(config.JWTKey)))
	e.DELETE("/users", uc.DeleteProfile(), middleware.JWT([]byte(config.JWTKey)))

	//lelangs 
	e.POST("/lelangs", ec.CreateLelangWithBid(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/lelangs", ec.GetLelangs())
	e.GET("/lelangs/:id", ec.GetLelang()) 
	e.PUT("/lelangs/:id", ec.UpdateLelang(),middleware.JWT([]byte(config.JWTKey))) 
	e.DELETE("/lelangs/:id", ec.DeleteLelang(), middleware.JWT([]byte(config.JWTKey)))
}