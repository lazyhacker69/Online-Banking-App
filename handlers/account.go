package handlers

import (
	"fmt"
	"net/http"
	"time"
	"errors"

	"online-banking/utils"
	"online-banking/dto/request"
	"online-banking/dto/response"
	"online-banking/models"
	"online-banking/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

		return
	}

	id := c.Param("id")

	result := database.DB.First(&account, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound){
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "account doesn't exist!",
		})

		return
	}

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

		return
	}

	//creating an atomic process

	err := database.DB.Transaction(func(tx *gorm.DB) error{

		account.Balance += deposit.Amount

		if err := tx.Save(&account).Error; err != nil {
			return err
		}

		transaction := models.Transaction{
			AccountID : account.AccountID,
			TransactionType : utils.TransactionDeposit,
			Amount : deposit.Amount,
		}

		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		return nil 
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})

		return
	}

	res := response.DepositRespone{
		Message: "Balance Updated!",
		AccountNumber: account.AccountNumber,
		NewBalance: account.Balance,
	}

	c.JSON(http.StatusOK, res)
}

func WithdrawMoney(c * gin.Context){

	var withdraw request.WithdrawRequest

	var account models.Account

	if err := c.BindJSON(&withdraw); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ 
			"error" : "Invalid request",
		})

		return
	}

	//checking if amount >= 0

	if withdraw.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Invalid amount",
		})

		return
	}

	id := c.Param("id")

	result := database.DB.First(&account, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound){
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "account doesn't exist!",
		})

		return
	}

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

		return
	}

	//checking if we have suff amount
	
	if withdraw.Amount > account.Balance {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "insufficient balance!",
		})
	}
	//creating an atomic process

	err := database.DB.Transaction(func(tx *gorm.DB) error{

		account.Balance -= withdraw.Amount

		if err := tx.Save(&account).Error; err != nil {
			return err
		}

		transaction := models.Transaction{
			AccountID : account.AccountID,
			TransactionType : utils.TransactionWithdraw,
			Amount : withdraw.Amount,
		}

		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		return nil 
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})

		return
	}

	res := response.WithdrawResponse{
		Message: "Withdrawal Successful!",
		AccountNumber: account.AccountNumber,
		NewBalance: account.Balance,
	}

	c.JSON(http.StatusOK, res)
}

func GetBalance(c *gin.Context){

	var account models.Account
	id := c.Param("id")

	result := database.DB.First(&account, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "Account not found!",
		})

		return
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : result.Error.Error(),
		})

		return
	}

	if account.Status == "Closed" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Account is Closed",
		})

		return
	}

	res := response.BalanceResponse{
		Balance : account.Balance,
		AccountNumber: account.AccountNumber,
	}

	c.JSON(http.StatusOK, res)
}

func GetTransactions(c * gin.Context){

	var transactions []models.Transaction
	var account models.Account

	id := c.Param("id")

	result := database.DB.First(&account, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound){
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "Account not found!",
		})

		return
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : result.Error.Error(),
		})

		return
	}

	if account.Status == utils.AccountClosed {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Account is Closed",
		})

		return
	}

	result = database.DB.Where(
		"account_id = ?",
		id,
	).Order("created_at DESC"). 
	Limit(10).
	Find(&transactions)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : result.Error.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, transactions)
}

func CloseAccount(c * gin.Context){

	id := c.Param("id")

	var account models.Account

	result := database.DB.First(&account, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound){
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "account not found",
		})

		return
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : result.Error.Error(),
		})

		return
	}

	if account.Status == utils.AccountClosed {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "account already closed",
		})

		return
	}

	if account.Balance > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Withdraw the remaining amount from the account!",
		})

		return
	}

	account.Status = utils.AccountClosed

	result = database.DB.Save(&account)

	if result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : result.Error.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "account closed successfully",
	})
}
