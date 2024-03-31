package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)


func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type DatabaseConfig struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int64    `json:"port"`
	Ssl                 bool   `json:"ssl"`
}


func main() {
	config := os.Getenv("db")
	config_test := os.Getenv("username")

	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": config,
			"message2": config_test,
		})
	})

	r.Run()
}