package handlers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func API(router *gin.Engine, classHandler ClassHandler, studentClassHandler StudentClassHandler) {

	classGroup := router.Group("/classes")
	classGroup.POST("/", classHandler.CreateHandler)
	classGroup.PUT("/:id", classHandler.UpdateHandler)
	classGroup.GET("/:id", classHandler.GetByIDHandler)
	classGroup.GET("/", classHandler.GetAllHandler)
	classGroup.DELETE("/:id", classHandler.DeleteHandler)

	studentClassGroup := router.Group("/student-classes")
	studentClassGroup.POST("/", studentClassHandler.CreateHandler)
	studentClassGroup.DELETE("/", studentClassHandler.DeleteHandler)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("api can't run in port 8085")
	}
}
