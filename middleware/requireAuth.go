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
		c.AbortWithStatus(http.StatusUnauthorized)
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
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Check the exp
	if exp, ok := claims["exp"].(float64); ok && float64(time.Now().Unix()) > exp {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Find the user with token sub
	var user models.User
	initializers.DB.First(&user, claims["sub"])

	var post models.Post
	initializers.DB.First(&post, claims["sub"])

	var event models.Event
	initializers.DB.First(&event, claims["sub"])

	if user.ID == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}


	// Attach to req
	c.Set("user", user)

	// Continue
	c.Next()
}
