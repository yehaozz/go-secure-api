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

// getSongs responds with the list of all songs as JSON
func GetSongs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, songs)
}
