package main

import (
	"fmt"
	"jsflor/messaging-service/controllers"
	"jsflor/messaging-service/lib"
	"jsflor/messaging-service/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()
	router := setUpRouter()

	router.Run(fmt.Sprintf("0.0.0.0:%s", lib.GetEnv("PORT", "8080")))
}

func setUpRouter() *gin.Engine {
	router := gin.Default()
	setUpRoutes(router)
	return router
}

func setUpRoutes(r *gin.Engine) {
	r.GET("/messages", controllers.GetMessages)
	r.GET("/messages/:id", controllers.GetMessageByID)
	r.POST("/messages", controllers.CreateMessage)
	r.PUT("/messages/:id", controllers.UpdateMessage)
	r.DELETE("/messages/:id", controllers.DeleteMessage)
}
