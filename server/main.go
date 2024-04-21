package main

import "github.com/gin-gonic/gin"

type User struct {
	Name    string
	Country string
}

func main() {
	r := gin.Default()

	user := &User{Name: "Jazz", Country: "Catland"}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"user":    user,
		})
	})
	r.Run(":2024")
}
