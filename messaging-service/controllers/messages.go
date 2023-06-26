package controllers

import (
	"jsflor/messaging-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /messages
func GetMessages(c *gin.Context) {
	var messages []models.Message

	if err := models.GetMessages(&messages).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, messages)
}

// GET /messages/:id
func GetMessageByID(c *gin.Context) {
	var message models.Message

	if err := models.GetMessageByID(&message, c.Param("id")).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, message)
}

// POST /messages
func CreateMessage(c *gin.Context) {
	var input models.CreateMessageInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := models.Message{Text: input.Text, From: input.From, To: input.To}

	if err := models.CreateMessage(&message).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, message)
}

// PUT /messages/:id
func UpdateMessage(c *gin.Context) {
	var message models.Message

	if err := models.GetMessageByID(&message, c.Param("id")).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input models.UpdateMessageInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateMessage(&message, &input).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}

// DELETE /messages/:id
func DeleteMessage(c *gin.Context) {
	var message models.Message

	if err := models.GetMessageByID(&message, c.Param("id")).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DeleteMessage(&message).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})
}
