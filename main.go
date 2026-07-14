package main

import(
	"os"

	"online-banking/database"
	"online-banking/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	database.ConnectDB()

	router := gin.Default()

	routes.RegisterRoutes(router)
	
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}