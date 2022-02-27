package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handler() {
	router := gin.Default()

	router.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		isExist := BoolIsRegisteredUser(username)

		if isExist {
			req := GetUser(username)

			match := CheckPasswordHash(password, req.Password)

			if !match {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": "Login failed!",
				})
			} else {
				// update JSON IsLogin = true
				c.JSON(http.StatusOK, gin.H{
					"status":   match,
					"username": username,
					"message":  "Login Succesful!",
				})

			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Login Failed User doesn't exist!",
			})
		}
	})

	router.POST("/logout", func(c *gin.Context) {
		username := c.PostForm("username")

		req := GetUser(username)

		if req.IsLogin {
			c.JSON(http.StatusOK, gin.H{
				"status":   true,
				"username": username,
				"message":  "Logout Succesful!",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  false,
				"message": "Logout failed!",
			})
		}
	})

	router.POST("/payment", func(c *gin.Context) {
		nominal := c.PostForm("nominal")

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"nominal": nominal,
			"message": "Payment Success!",
		})
	})

	router.Run("localhost:8080")
}
