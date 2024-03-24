package main

import (
	"BlogServ/logger"
	"fmt"
)

func main() {
	logger.Info("test")

	fmt.Println("hello golang")
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})
	//
	//r.Run(":9091")
}
