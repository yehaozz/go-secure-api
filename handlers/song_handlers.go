package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yehaozz/go-secure-api/models"
)

var songs = []models.Song{
	{ID: "1", Title: "September", Artist: "Earth, Wind & Fire", Rate: 4.3},
	{ID: "2", Title: "Fun For All", Artist: "Vinida", Rate: 4.7},
	{ID: "3", Title: "Coco Elva Tia", Artist: "MaSiWei", Rate: 4.7},
}

// GetSongs responds with the list of all songs as JSON
func GetSongs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, songs)
}

// PostSong adds a song to the songs slice
func PostSong(c *gin.Context) {
	var newSong models.Song

	// Bind the received JSON to newSong
	if err := c.BindJSON(&newSong); err != nil {
		return
	}

	// Add the newSong to the slice
	songs = append(songs, newSong)
	c.IndentedJSON(http.StatusCreated, newSong)
}
