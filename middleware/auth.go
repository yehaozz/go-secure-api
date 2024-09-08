package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lestrrat-go/jwx/jwk"
)

const (
	jwksURL    = "https://www.googleapis.com/oauth2/v3/certs"
	authPrefix = "Bearer "
)

// JWTClaims represents the JWT claims structure
type JWTClaims struct {
	Audience string `json:"aud"`
	Issuer   string `json:"iss"`
	Subject  string `json:"sub"`
	Expiry   int64  `json:"exp"`
	jwt.RegisteredClaims
}

// FetchPublicKey fetches the public key from the JWKS endpoint
func FetchPublicKey(token *jwt.Token) (any, error) {
	// Fetch the JWKS
	set, err := jwk.Fetch(context.Background(), jwksURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch JWKS: %v", err)
	}

	// Find the key by kid (key id) in the JWT header
	kid, found := token.Header["kid"]
	if !found {
		return nil, fmt.Errorf("no 'kid' in token header")
	}

	key, found := set.LookupKeyID(kid.(string))
	if !found {
		return nil, fmt.Errorf("no matchking key id found in JWKS")
	}

	var publicKey any
	if err := key.Raw(&publicKey); err != nil {
		return nil, fmt.Errorf("failed to parse public key: %v", err)
	}

	return publicKey, nil
}

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

		// Parse the tokenString JWT
		// TODO what is a JWTClaim
		token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, FetchPublicKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
			// token is valid, proceed with further checks (expiry, issuer, etc.)
			if claims.Expiry < time.Now().Unix() {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
				return
			}

			// Add claims to the context, so they can be access in handlers
			c.Set("claims", claims)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}
	}
}
