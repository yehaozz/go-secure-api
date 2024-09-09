package handlers

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yehaozz/go-secure-api/models"
)

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
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("song with id %s not found", id)})
		return
	}

	c.IndentedJSON(http.StatusOK, songs[id])
}

// UpdateSong updates an existing song by ID
func UpdateSong(c *gin.Context) {
	id := c.Param("id")
	var updatedSong models.Song

	// Bind the received JSON to newSong
	if err := c.BindJSON(&updatedSong); err != nil {
		return
	}

	mu.Lock()
	defer mu.Unlock()

	existingSong, exist := songs[id]
	if !exist {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("song with id %s not found", id)})
		return
	}

	// Update the fields of the existing song
	if updatedSong.Title != "" {
		existingSong.Title = updatedSong.Title
	}
	if updatedSong.Artist != "" {
		existingSong.Artist = updatedSong.Artist
	}
	if updatedSong.Rating != nil {
		existingSong.Rating = updatedSong.Rating
	}

	// Save the updated song
	songs[id] = existingSong

	c.IndentedJSON(http.StatusOK, existingSong)
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
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("song with id %s not found", id)})
		return
	}

	delete(songs, id)
	c.Status(http.StatusNoContent)
}
