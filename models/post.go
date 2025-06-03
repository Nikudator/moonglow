package models

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Lead    string `json:"lead"`
	Body    string `json:"body"`
	Created int64  `json:"created"`
	Updated int64  `json:"updated"`
	Author  string `json:"author"`
}
