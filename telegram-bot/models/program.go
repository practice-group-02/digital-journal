package models

type Program struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Country     string `json:"country"`
	Deadline    string `json:"deadline"`
	Description string `json:"description"`
}
