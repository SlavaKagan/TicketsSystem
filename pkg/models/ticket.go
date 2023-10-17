package models

// represents data about a ticket.

type Ticket struct {
	Id           string   `json:"id"`
	Title        string   `json:"title"`
	Content      string   `json:"content"`
	UserEmail    string   `json:"userEmail"`
	CreationTime int      `json:"creationTime"`
	Labels       []string `json:"labels"`
}
