package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yehaozz/go-secure-api/handlers"
)

// RegisterRoutes defines the API routes
func RegisterRoutes(r *gin.Engine) {
	r.GET("/songs", handlers.GetSongs)
	r.POST("/songs", handlers.PostSong)
}
