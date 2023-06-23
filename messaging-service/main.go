package main

import (
	"fmt"
	"jsflor/messaging-service/controllers"
	"jsflor/messaging-service/lib"
	"jsflor/messaging-service/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/messages", controllers.GetMessages)
	router.GET("/messages/:id", controllers.GetMessageByID)
	router.POST("/messages", controllers.CreateMessage)
	router.PUT("/messages/:id", controllers.UpdateMessage)
	router.DELETE("/messages/:id", controllers.DeleteMessage)

	router.Run(fmt.Sprintf("0.0.0.0:%s", lib.GetEnv("PORT", "8080")))
}
