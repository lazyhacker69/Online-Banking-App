package routes

import (
	"online-banking/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAccountsRoutes(router *gin.Engine) {

	router.POST("/accounts", handlers.CreateAccount)
	router.POST("/accounts/:id/deposit", handlers.DepositMoney)
	router.POST("/accounts/:id/withdraw", handlers.WithdrawMoney)
	router.GET("accounts/:id/balance", handlers.GetBalance)
	router.GET("/accounts/:id/transations", handlers.GetTransactions)
	router.PATCH("/accounts/:id/close", handlers.CloseAccount)

}
