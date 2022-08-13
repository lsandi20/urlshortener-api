package models

type Link struct {
	ID       string `json:"id"`
	Original string `json:"original"`
	Short    string `json:"short"`
	Clicks   int64  `json:"clicks"`
}
