package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func getApiPort() string {
	return os.Getenv("API_PORT")
}

func main() {
	r := gin.Default()
	PORT := getApiPort()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":" + PORT) // listen and serve on 0.0.0.0:8080
}
