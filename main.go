package main

import (
	"BukaLelang/config"
	"fmt"
	"BukaLelang/app/routes"
	 userHandler"BukaLelang/app/features/users/handler"
	 userRepo"BukaLelang/app/features/users/repository"
	 userLogic"BukaLelang/app/features/users/services"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.GetConfiguration()
	db, _ := config.GetConnection(*cfg) 
	config.Migrate(db)

	userModel := userRepo.New(db)
	userServices := userLogic.New(userModel)
	userController := userHandler.New(userServices)

	routes.Route(e, userController) 

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}