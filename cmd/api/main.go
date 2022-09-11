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
	router.Use(CORSMiddleware())

	handlers.API(router, classHandler, studentClassHandler)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
