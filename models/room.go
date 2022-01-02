package models

type Room struct {
	Id           string   `json:"id"`
	Title        string   `json:"title"`
	Participants []string `json:"participants"`
}