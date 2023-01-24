package main

import (
	"log"
	"time"

	"github.com/felix1369/golang-api/api/config"
	"github.com/felix1369/golang-api/api/handlers"
	"github.com/felix1369/golang-api/repository"
	"github.com/gin-gonic/gin"

	_deliveryMiddleware "github.com/felix1369/golang-api/api/middlewares"
	_roleRepo "github.com/felix1369/golang-api/repository"
	_roleUsecase "github.com/felix1369/golang-api/usecase"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	dbConn, err := repository.OpenDB("postgres", config.Env.DbConn)

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

	r := gin.Default()
	r.Use(_deliveryMiddleware.CORS())
	roleRepo := _roleRepo.NewSqlRole(dbConn)
	timeoutContext := time.Duration(config.Env.Timeout) * time.Second
	au := _roleUsecase.NewRoleUsecase(roleRepo, timeoutContext)
	handlers.NewRoleHandler(r, au)

	r.Run(":5000")
}
