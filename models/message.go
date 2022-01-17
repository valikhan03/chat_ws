package models

type Message struct {
	Sender  string `json:"sender"`
	ChatID  string `json:"chat_id"`
	Payload string `json:"payload"`
	Date    string `json:"date"`
}