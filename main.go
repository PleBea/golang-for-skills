package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})

	r.GET("/db-test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"DB_HOST": os.Getenv("DB_HOST"), "DB_PORT": os.Getenv("DB_PORT"), "DB_USER": os.Getenv("DB_USER"), "DB_PASSWORD": os.Getenv("DB_PASSWORD"), "DB_NAME": os.Getenv("DB_NAME")})
	})

	r.Run()
}