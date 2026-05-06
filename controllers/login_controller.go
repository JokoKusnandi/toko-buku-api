package controllers

import (
	"os"
	"toko-buku-api/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	//dummy user

	if user.Username != "admin" || user.Password != "123456" {
		c.JSON(401, gin.H{"error": "invalid"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user.Username,
	})

	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	c.JSON(200, gin.H{"token": tokenString})

}
