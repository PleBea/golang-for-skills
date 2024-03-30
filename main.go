package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func loadEnvironment() (DatabaseConfig) {
	stringConfig := os.Getenv("db")
	var config DatabaseConfig

	err := json.Unmarshal([]byte(stringConfig), &config)
	handleError(err)

	return config 
}

func main() {
	config := loadEnvironment()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	uri := "mongodb://" + config.Username + ":" + config.Password + "@" + config.Host + ":" + strconv.FormatInt(config.Port, 10) + "/?ssl=" + strconv.FormatBool(config.Ssl)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	handleError(err)

	uesrsCollection := client.Database("test").Collection("user")

	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		cursor, _ := uesrsCollection.Find(context.Background(), bson.M{})
		defer cursor.Close(context.Background())

		var result []*bson.M

		for cursor.Next(context.Background()) {
			var elem bson.M
			err := cursor.Decode(&elem)
			handleError(err)
			result = append(result, &elem)
		}
		err := cursor.Err();
		handleError(err)

		c.JSON(http.StatusOK, gin.H{
			"user": result,
		})
	})

	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		_, err := uesrsCollection.InsertOne(context.Background(), bson.M{
			name: name,
		})
		handleError(err)

		c.JSON(http.StatusOK, gin.H{
			"message": name,
		})
	})

	r.Run()
}