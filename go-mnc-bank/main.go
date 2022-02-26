package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsLogin  bool   `json:"isLogin"`
}

var users = []user{
	{ID: "1", Username: "usertest@test.com", Password: "$2a$14$HX09NEZcCCmpyKNlIF71DOxDfuh3y2HREjGETGe3FekgyrA2S//yy", IsLogin: false},
	{ID: "2", Username: "usertest2@test.com", Password: "$2a$14$HX09NEZcCCmpyKNlIF71DOxDfuh3y2HREjGETGe3FekgyrA2S//yy", IsLogin: true},
}

func main() {
	router := gin.Default()

	router.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		for _, a := range users {
			if a.Username == username {
				//hash, _ := HashPassword(password) // ignore error for the sake of simplicity
				match := CheckPasswordHash(password, a.Password)

				if !match {
					c.JSON(http.StatusInternalServerError, gin.H{
						"status": "Login failed!",
					})
				} else {
					a.IsLogin = true
					c.JSON(http.StatusOK, gin.H{
						"status":   match,
						"username": username,
						"message":  "Login Succesful!",
					})

				}
			} else {
				c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
			}
		}

	})

	router.POST("/logout", func(c *gin.Context) {
		username := c.PostForm("username")

		for _, a := range users {
			if a.IsLogin {
				c.JSON(http.StatusOK, gin.H{
					"status":   true,
					"username": username,
					"message":  "Logout Successful!",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status":   false,
					"username": username,
					"message":  "You've already logout!",
				})
			}
		}
	})

	router.POST("/payment", func(c *gin.Context) {
		nominal := c.PostForm("nominal")

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"nominal": nominal,
			"message": "Payment Success!",
		})
		// } else {
		// 	c.JSON(http.StatusOK, gin.H{
		// 		"status":  false,
		// 		"nominal": nominal,
		// 		"message": "Payment Failed! Nominal is empty.",
		// 	})

	})

	router.Run("localhost:8080")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
