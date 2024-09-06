package handlers

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yehaozz/go-secure-api/models"
)

// var songs = []models.Song{
// 	{ID: "1", Title: "September", Artist: "Earth, Wind & Fire", Rating: 4.3},
// 	{ID: "2", Title: "Fun For All", Artist: "Vinida", Rating: 4.7},
// 	{ID: "3", Title: "Coco Elva Tia", Artist: "MaSiWei", Rating: 4.7},
// }

// songs is a map of id to song's data
var songs = make(map[string]models.Song)
var mu sync.Mutex

// Helper function to generate an ID for a song
func generateID() string {
	return uuid.New().String()
}

// GetSongs responds with the list of all songs as JSON
func GetSongs(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	songList := []models.Song{}
	for _, song := range songs {
		songList = append(songList, song)
	}

	c.IndentedJSON(http.StatusOK, songList)
}

// GetSong responds with a particular song by ID
func GetSong(c *gin.Context) {
	id := c.Param("id")

	mu.Lock()
	defer mu.Unlock()

	if _, exist := songs[id]; !exist {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("song with id %s not found", id)})
	}

	c.IndentedJSON(http.StatusOK, songs[id])
}

// PostSong adds a song to the songs slice
func PostSong(c *gin.Context) {
	var newSong models.Song

	// Bind the received JSON to newSong
	if err := c.BindJSON(&newSong); err != nil {
		return
	}

	mu.Lock()
	defer mu.Unlock()
	newSong.ID = generateID()
	songs[newSong.ID] = newSong

	c.IndentedJSON(http.StatusCreated, newSong)
}

// DeleteSong deletes a song by ID
func DeleteSong(c *gin.Context) {
	id := c.Param("id")

	mu.Lock()
	defer mu.Unlock()

	if _, exist := songs[id]; !exist {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("song with id %s not found", id)})
	}

	delete(songs, id)
	c.Status(http.StatusNoContent)
}
