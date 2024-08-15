package controllers

import "github.com/gin-gonic/gin"

func HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}
func NotFound(c *gin.Context) {
	c.JSON(404, gin.H{
		"message": "Not Found",
	})
}
