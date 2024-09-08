package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yehaozz/go-secure-api/handlers"
	"github.com/yehaozz/go-secure-api/middleware"
)

// RegisterRoutes defines the API routes
func RegisterRoutes(r *gin.Engine) {
	r.GET("/songs", handlers.GetSongs)
	r.GET("/songs/:id", handlers.GetSong)
	r.POST("/songs", middleware.JWTMiddleware(), handlers.PostSong)
	r.PUT("/songs/:id", middleware.JWTMiddleware(), handlers.UpdateSong)
	r.DELETE("/songs/:id", middleware.JWTMiddleware(), handlers.DeleteSong)
}
