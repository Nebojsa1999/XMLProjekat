package model

type Post struct {
	Id    uint   `json:"id"`
	Text  string `json:"text"`
	Link  string `json:"link"`
	Image string `json:"image"`
}
