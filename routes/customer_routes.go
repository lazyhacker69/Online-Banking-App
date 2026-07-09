package routes

import(
	"github.com/gin-gonic/gin"
	"online-banking/handlers"
)

func RegisterCustomerRoutes(router *gin.Engine){
	
	router.POST("/customers", handlers.CreateCustomer)
	router.GET("/customers/:id", handlers.GetCustomer)
	
}