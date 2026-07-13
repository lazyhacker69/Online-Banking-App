package routes

import(
	"github.com/gin-gonic/gin"
	"online-banking/handlers"
)
func RegisterLoanRoutes(router * gin.Engine){
	router.POST("/loans", handlers.CreateLoan)
	router.GET("/loans/:id", handlers.GetLoan)
	router.GET("/customers/:id/loans", handlers.GetCustomerLoans)
}