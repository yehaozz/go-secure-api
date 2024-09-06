package models

// song represents data about a music song
type Song struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Rate   float64 `json:"rate"`
}
