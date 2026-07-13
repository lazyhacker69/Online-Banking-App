package routes

import(
	"online-banking/handlers"
	
	"github.com/gin-gonic/gin"
)
func RegisterLoanPaymentsRoutes(router * gin.Engine){
	router.POST("/loans/:id/repay", handlers.RepayLoan)
}