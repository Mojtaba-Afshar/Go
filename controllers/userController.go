package controllers

import (
	"github.com/Mojtaba-Afshar/ekart/services"
	"github.com/gin-gonic/gin"
)

func HandleLogin(c *gin.Context) {
	services.Login
	c.JSON(200, gin.H{
		"data": "Welcome to login",
	})
}

func HandleRegister(c *gin.Context) {
	services.Register()
	c.JSON(200, gin.H{
		"data": "Register",
	})
}

func HandleProfile(c *gin.Context) {
	services.GetProfile()
	c.JSON(200, gin.H{
		"data": "Profile",
	})
}

func HandleLogout(c *gin.Context) {
	services.LogOut()
	c.JSON(200, gin.H{
		"data": "Logged out",
	})
}
