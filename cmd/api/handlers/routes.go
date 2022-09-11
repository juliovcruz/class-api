package handlers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
)

func API(classHandler ClassHandler, studentClassHandler StudentClassHandler) {
	router := gin.Default()
	router.Use(CORS())

	router.POST("/classes", classHandler.CreateHandler)
	router.PUT("/classes/:id", classHandler.UpdateHandler)
	router.GET("/classes/:id", classHandler.GetByIDHandler)
	router.GET("/classes", classHandler.GetAllHandler)
	router.GET("/classes/:id/student-classes", classHandler.GetAllStudentClassesByHandler)
	router.DELETE("/classes/:id", classHandler.DeleteHandler)

	router.POST("/student-classes", studentClassHandler.CreateHandler)
	router.DELETE("/student-classes", studentClassHandler.DeleteHandler)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err := router.Run(":8081")
	if err != nil {
		log.Fatal("api can't run in port 8085")
	}
}

func CORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		if method == "OPTIONS" {
			ctx.Header("Access-Control-Max-Age", "1728000")
			ctx.Header("Access-Control-Allow-Credentials", "true")
			ctx.Header("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,PATCH,OPTIONS")
			ctx.Header("Access-Control-Allow-Origin", ctx.Request.Header.Get("Access-Control-Allow-Origin"))
			ctx.Header("Access-Control-Allow-Headers", "Content-Type,Cookie,Authorization,Access-Control-Request-Headers,Access-Control-Request-Method,Origin,Referer,Sec-Fetch-Dest,Accept-Language,Accept-Encoding,Sec-Fetch-Mode,Sec-Fetch-Site,User-Agent,Pragma,Host,Connection,Cache-Control,Accept-Language,Accept-Encoding,X-Requested-With,X-Forwarded-For,X-Forwarded-Host,X-Forwarded-Proto,X-Forwarded-Port,X-Forwarded-Prefix,X-Real-IP,Accept")
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}
		ctx.Header("Access-Control-Allow-Origin", ctx.Request.Header.Get("Access-Control-Allow-Origin"))
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Next()
	}
}
