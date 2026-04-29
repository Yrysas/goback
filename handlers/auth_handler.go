package handlers

import (
	"net/http"

	"finance_tracker/config"
	"finance_tracker/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret_key")

func Register(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	config.DB.Create(&user)
	c.JSON(200, user)
}

func Login(c *gin.Context) {
	var input models.User
	var user models.User

	c.BindJSON(&input)

	config.DB.Where("username = ?", input.Username).First(&user)

	if user.Password != input.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
	})

	tokenString, _ := token.SignedString(jwtKey)

	c.JSON(200, gin.H{"token": tokenString})
}