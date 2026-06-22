package main

import "github.com/gin-gonic/gin"

func getUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"email": "john@doe.com",
	})
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.Run()
}