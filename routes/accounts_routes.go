package routes

import(
	"online-banking/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterAccountsRoutes(router * gin.Engine){

	router.POST("/accounts", handlers.CreateAccount)
	router.POST("/account/:id/diposit", handlers.DepositMoney)

}