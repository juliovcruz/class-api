package main

import (
	"github.com/gin-contrib/cors"
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
	"time"
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
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowMethods:           []string{"GET", "DELETE", "PUT", "PATCH", "OPTIONS"},
		AllowHeaders:           []string{"Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
		AllowCredentials:       true,
		ExposeHeaders:          []string{"Content-Length"},
		MaxAge:                 12 * time.Hour,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
		AllowWebSockets:        true,
		AllowFiles:             true,
	}))

	handlers.API(router, classHandler, studentClassHandler)
}
