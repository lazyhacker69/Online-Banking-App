package routes

import(
	"github.com/gin-gonic/gin"
	"online-banking/handlers"
)
func RegisterLoanRoutes(router * gin.Engine){
	router.POST("/account/:id/loan", handlers.CreateLoan)
}