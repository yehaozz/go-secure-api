package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authPrefix   = "Bearer "
	tokenInfoURL = "https://oauth2.googleapis.com/tokeninfo?access_token="
)

// JWTMiddleware validates the JWT and checks its claims
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, authPrefix) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, authPrefix)

		// Call the Google Token Info endpoint to validate the token
		resp, err := http.Get(tokenInfoURL + tokenString)
		if err != nil || resp.StatusCode != http.StatusOK {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		defer resp.Body.Close()

		// Parse the response to extract claims
		var claims map[string]any
		if err := json.NewDecoder(resp.Body).Decode(&claims); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to decode token claims"})
			c.Abort()
			return
		}

		// Optionally: You can perform additional checks on claims if needed
		// Example: if claims["aud"] != "<expected_audience>" { ... }

		// Add claims to the context
		c.Set("claims", claims)
		c.Next()
	}
}
