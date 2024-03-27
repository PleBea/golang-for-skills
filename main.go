package main

import (
	"context"
	"log"
	"net/http"
	"os"
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

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
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
			if err != nil {
				log.Fatal(err)
			}
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