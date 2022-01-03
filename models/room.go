package models

type Room struct {
	Id           string   `json:"id"`
	Title        string   `json:"title"`
	Owner        string   `json:"owner"`
	Participants []string `json:"participants"`
}
