package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Text string `json:"text"`
	From string `json:"from"`
	To   string `json:"to"`
}

type CreateMessageInput struct {
	Text string `json:"text" binding:"required"`
	From string `json:"from" binding:"required"`
	To   string `json:"to" binding:"required"`
}

type UpdateMessageInput struct {
	Text string `json:"text"`
	From string `json:"from"`
	To   string `json:"to"`
}

func GetMessages(messages *[]Message) *gorm.DB {
	return DB.Find(&messages)
}

func GetMessageByID(message *Message, id string) *gorm.DB {
	return DB.First(&message, id)
}

func CreateMessage(message *Message) *gorm.DB {
	return DB.Create(&message)
}

func UpdateMessage(message *Message, input *UpdateMessageInput) *gorm.DB {
	return DB.Model(&message).Updates(input)
}

func DeleteMessage(message *Message) *gorm.DB {
	return DB.Delete(&message)
}
