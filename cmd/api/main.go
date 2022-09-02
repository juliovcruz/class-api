package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"main/cmd/api/docs"
	"main/cmd/api/handlers"
	"main/internal/account_service"
	"main/internal/platform"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	scope := platform.Scope(os.Getenv("SCOPE"))
	docs.SwaggerInfo.BasePath = "/"

	db, err := gorm.Open(postgres.Open(scope.GetDBConnectionURL()), &gorm.Config{})
	if err != nil {
		log.Fatal("Can't be connect to database")
	}

	accountService := account_service.NewAccountService(scope.GetAccountServiceURL())
	validatr := validator.New()

	classHandler := handlers.NewClassHandler(db, accountService, validatr)
	studentClassHandler := handlers.NewStudentClassHandler(db, accountService, validatr)

	router := gin.Default()

	handlers.API(router, classHandler, studentClassHandler)
}
