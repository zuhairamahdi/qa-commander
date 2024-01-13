package middleware

import (
	"fmt"
	"net/http"
	"qa_commander/internal/repository"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type middlewareHandler struct {
	UserRepo *repository.UserRepository
}

func NewMiddlewareHandler(userRepo *repository.UserRepository) *middlewareHandler {
	return &middlewareHandler{
		UserRepo: userRepo,
	}
}

var SecretKey = []byte("your-secret-key")

// GenerateJWT generates a new JWT token.
func GenerateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	// Set claims
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Sign the token with the secret key
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "unable to generate authentication token", err
	}

	return tokenString, nil
}

// RequireAuthentication is a middleware to verify JWT tokens.
func RequireAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix from the token string
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return SecretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Check if the token is expired
		if isTokenExpired(token) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
			c.Abort()
			return
		}

		// Set the user information in the context if needed
		// check if "user" key exists in the context
		// if _, ok := c.Get("user"); !ok {
		// 	//get user from repository
		// 	user, err := ra.UserRepo.GetUserByUsername(token.Claims.(jwt.MapClaims)["username"].(string))
		// 	if err != nil {
		// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		// 		c.Abort()
		// 		return
		// 	}
		// 	c.Set("user", user)
		// }

		// If the token is valid and not expired, proceed to the next middleware
		c.Next()
	}
}

// isTokenExpired checks if the token is expired.
func isTokenExpired(token *jwt.Token) bool {
	claims := token.Claims.(jwt.MapClaims)
	fmt.Println(claims["exp"])
	exp := int64(claims["exp"].(float64)) // Convert exp to int64
	return exp < time.Now().Unix()
}
