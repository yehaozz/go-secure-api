package models

// song represents data about a music song
type Song struct {
	ID     string   `json:"id"`
	Title  string   `json:"title"`
	Artist string   `json:"artist"`
	Rating *float64 `json:"rating,omitempty"`
}

func IsSameSong(songOne, songTwo Song) bool {
	return songOne.ID == songTwo.ID &&
		songOne.Title == songTwo.Title &&
		songOne.Artist == songTwo.Artist &&
		*songOne.Rating == *songTwo.Rating
}
