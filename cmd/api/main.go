package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"main/cmd/api/docs"
	"main/cmd/api/handlers"
	"main/internal/account_service"
)

func main() {
	docs.SwaggerInfo.BasePath = "/"
	dbConn := "postgres://postgres:postgres@postgres-class-api:5432/postgres" // -- docker
	// dbConn := "postgres://postgres:postgres@localhost:5435/postgres" // -- local
	db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})
	if err != nil {
		log.Fatal("Can't be connect to database")
	}

	accountService := account_service.NewAccountService("http://localhost:5187")
	validatr := validator.New()

	classHandler := handlers.NewClassHandler(db, accountService, validatr)
	studentClassHandler := handlers.NewStudentClassHandler(db, accountService)

	router := gin.Default()

	handlers.API(router, classHandler, studentClassHandler)
}
