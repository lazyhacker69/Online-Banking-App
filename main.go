package main

import(
	"online-banking/database"
	"online-banking/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	database.ConnectDB()

	router := gin.Default()

	routes.RegisterRoutes(router)
	
	router.Run(":8080")
}