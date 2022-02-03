package models

type Room struct {
	Id           string   `json:"id" bson:"id"`
	Title        string   `json:"title" bson:"title"`
	Creator      string   `json:"owner" bson:"owner"`
	Participants []string `json:"participants" bson:"participants"`
	Type         string   `json:"type" bson:"type"` //common, group
}
