package main

import (
	"log"
	"time"

	"github.com/felix1369/golang-api/api/config"
	"github.com/felix1369/golang-api/repository"
	"github.com/labstack/echo"

	_articleHttpDelivery "github.com/felix1369/golang-api/api/handlers"
	_articleHttpDeliveryMiddleware "github.com/felix1369/golang-api/api/middlewares"
	_roleRepo "github.com/felix1369/golang-api/repository"
	_roleUsecase "github.com/felix1369/golang-api/usecase"
)

func main() {
	dbConn, err := repository.OpenDB(`postgres`, config.Env.DbConn)

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.DB().Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	middL := _articleHttpDeliveryMiddleware.InitMiddleware()
	e.Use(middL.CORS)
	roleRepo := _roleRepo.NewMysqlRole(dbConn)

	timeoutContext := time.Duration(config.Env.Timeout) * time.Second
	au := _roleUsecase.NewRoleUsecase(roleRepo, timeoutContext)
	_articleHttpDelivery.NewRoleHandler(e, au)
}
