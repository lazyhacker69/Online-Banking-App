package handlers

import (
	"fmt"
	"net/http"
	"online-banking/dto/request"
	"online-banking/dto/response"
	"online-banking/models"
	"time"

	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"online-banking/database"
)

func CreateAccount(c *gin.Context) {

	var req request.CreateAccountRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid account data",
		})
		return
	}

	//checking if customer exist

	var customer models.Customer

	result := database.DB.First(&customer, req.CustomerID)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Customer doesnt exist",
		})

		return
	}

	//checking if account already exists

	var existingAccount models.Account

	result = database.DB.Where(
		"customer_id = ? AND account_type = ?",
		req.CustomerID,
		req.AccountType,
	).First(&existingAccount)

	if result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Customer already has this account type",
		})

		return
	}

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})

		return
	}

	generatedNumber := fmt.Sprintf("%010d", time.Now().UnixNano()%10000000000)
	account := models.Account{
		CustomerID:  req.CustomerID,
		AccountType: req.AccountType,

		AccountNumber: generatedNumber,
		Balance:       0,
		Status:        "Active",
	}

	result = database.DB.Create(&account)

	if result.Error != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})

		return
	}

	res := response.CreateAccountResponse{
		Message:       "Account Created successfully",
		AccountNumber: account.AccountNumber,
		AccountType:   account.AccountType,
		Balance:       account.Balance,
	}

	c.JSON(http.StatusCreated, res)

}

func DepositMoney(c * gin.Context){

	var deposit request.DepositRequest

	var account models.Account

	if err := c.BindJSON(&deposit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ 
			"error" : "Invalid request",
		})

		return
	}

	//checking if amount >= 0

	if deposit.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Invalid amount",
		})
	}

	id := c.Param("id")

	result := database.DB.First(&account, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : result.Error.Error(),
		})

		return
	}

	if account.Status != "Active" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Account is closed",
		})
	}
	account.Balance += deposit.Amount

	result = database.DB.Save(&account)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : result.Error.Error(),
		})

		return
	}

	var transaction models.Transaction

	res := response.DepositRespone{
		Message: "Balance Updated!",
		AccountNumber: account.AccountNumber,
		NewBalance: account.Balance,
	}

	c.JSON(http.StatusOK, res)
}