package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(myMiddleware(1))
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("process ping...")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Use(myMiddleware(2))
	r.GET("/test", func(c *gin.Context) {
		fmt.Println("process test..")
		c.JSON(200, gin.H{
			"message": "result",
		})
	})
	r.Run()
}

func myMiddleware(index int) gin.HandlerFunc {
	return func(g *gin.Context) {
		fmt.Println("begin", index)
		g.Next()
		fmt.Println("after", index)
	}
}
