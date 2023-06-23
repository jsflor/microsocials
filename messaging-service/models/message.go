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
