package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	// Get the cookie off req
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Anda belum login"})
		return
	}

	// Decode/validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
		return
	}

	// Check the exp
	if exp, ok := claims["exp"].(float64); ok && float64(time.Now().Unix()) > exp {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token kedaluwarsa"})
		return
	}

	// Find the user with token sub
	var user models.User
	initializers.DB.First(&user, claims["sub"])

	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak ditemukan"})
		return
	}

	// Check user's level
	if user.Email != "dave@gmail." {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Anda tidak memiliki akses yang cukup"})
		return
	}

	// Attach to req
	c.Set("user", user)

	// Continue
	c.Next()
}
