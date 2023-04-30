package router

import "github.com/gin-gonic/gin"

// Funcitons with first letter captalized are automatically exported
func Initializ() {
	router := gin.Default()
	// Defeinido uma Rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Executando a api
	router.Run(":8080") 
}