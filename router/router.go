package router

import "github.com/gin-gonic/gin"

// Funcitons with first letter captalized are automatically exported
func Initializ() {
	//  Initialize Router
	router := gin.Default()

	// Initialize Routes
	initializeRoutes(router)
	
	// Run the API
	router.Run(":8080") 
}