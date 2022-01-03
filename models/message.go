package models

type Message struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Payload  string `json:"payload"`
	Date     string `json:"date"`
}
