package models

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Lead    string `json:"lead"`
	Body    string `json:"body"`
	Created int    `json:"created"`
	Updated int    `json:"udated"`
	Author  string `json:"author"`
}
