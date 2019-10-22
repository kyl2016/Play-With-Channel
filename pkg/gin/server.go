package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.Any("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "any path"})
	})

	//router.GET("/", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"message": "use get method"})
	//})
	router.Run()
}
