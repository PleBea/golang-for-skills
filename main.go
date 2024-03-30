package main

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	r := gin.Default()

	rdb := redis.NewClient(&redis.Options{
		Addr: "srn-redis-4isy4c.serverless.apn2.cache.amazonaws.com:6379",
		Password: "",
		DB: 0,
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	})


	r.GET("/world", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "world",
		})
	})

	r.GET("/world/cache/:value", func (c *gin.Context)  {
		status := "get cache"

		value := c.Param("value")

		res, err := rdb.Get(ctx, value).Result()
		if err == redis.Nil {
			err := rdb.Set(ctx, value, len(value), 3600*time.Second).Err()
			if err != nil {
				panic(err)
			}
			status = "set cache"
		}  else if err != nil {
			panic(err)
		} 
		
		c.JSON(http.StatusOK, gin.H{
			"message": res,
			"stutus": status,
		})
	})
	r.Run()
}