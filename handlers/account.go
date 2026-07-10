package handlers

import (
	"fmt"
	"time"
	"net/http"
	"online-banking/dto/request"
	"online-banking/dto/response"
	"online-banking/models"

	"github.com/gin-gonic/gin"
	"online-banking/database"
)


func CreateAccount(c * gin.Context){

	var req request.CreateAccountRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "invalid account data",
		})
		return
	}

	//checking if customer exist

	var customer models.Customer

	result := database.DB.First(&customer, req.CustomerID)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Customer doesnt exist",
		})

		return
	}


	generatedNumber := fmt.Sprintf("%010d", time.Now().UnixNano()%10000000000)
	account := models.Account{
		CustomerID: req.CustomerID,
		AccountType: req.AccountType,

		AccountNumber: generatedNumber,
		Balance: 0,
		Status: "Active",
	}

	result = database.DB.Create(&account)

	if result.Error != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : result.Error.Error(),
		})

		return
	}

	res := response.CreateAccountResponse{
		Message : "Account Created successfully",
		AccountNumber : account.AccountNumber,
		AccountType: account.AccountType,
		Balance: account.Balance,
	}

	c.JSON(http.StatusCreated, res)

}