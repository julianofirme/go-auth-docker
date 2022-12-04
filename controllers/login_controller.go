package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jfirme-sys/go-auth-docker.git/database"
	"github.com/jfirme-sys/go-auth-docker.git/models"
	"github.com/jfirme-sys/go-auth-docker.git/services"
	"github.com/jfirme-sys/go-auth-docker.git/util"
)

func Login(c *gin.Context) {
	db := database.GetDB()

	var credentials models.Login

	err := c.ShouldBindJSON(&credentials)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	var user models.User
	dbError := db.Where("email = ?", credentials.Email).First(&user).Error
	if dbError != nil {
		c.JSON(400, gin.H{
			"error": "cannot find user",
		})
		return
	}

	if user.Password != util.SHA256Encoder(credentials.Password) {
		c.JSON(400, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	token, err := services.NewJWTService().GenerateToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
