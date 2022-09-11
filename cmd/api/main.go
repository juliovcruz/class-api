package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	cors "github.com/rs/cors/wrapper/gin"
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
	router.Use(cors.Default())
	//router.Use(cors.New(cors.Options{
	//	AllowedOrigins:         []string{"*"},
	//	AllowOriginFunc:        nil,
	//	AllowOriginRequestFunc: nil,
	//	AllowedMethods:         []string{"POST", "PUT", "DELETE", "PATCH", "GET", "OPTIONS"},
	//	AllowedHeaders:         []string{"*"},
	//	MaxAge:                 10000000000000,
	//	AllowCredentials:       false,
	//	OptionsPassthrough:     false,
	//	OptionsSuccessStatus:   0,
	//	Debug:                  false,
	//}))

	handlers.API(router, classHandler, studentClassHandler)
}
