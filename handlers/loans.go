package handlers

import (
	"errors"
	"net/http"

	"online-banking/database"
	"online-banking/dto/request"
	"online-banking/dto/response"
	"online-banking/models"
	"online-banking/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateLoan(c * gin.Context){

	var req request.ApplyLoanRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Invalid JSON",
		})

		return
	}

	if req.Amount < 10000 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Minimum loan amount is 10000!",
		})

		return
	}

	var account models.Account

	result := database.DB.First(&account, req.AccountID)

	if errors.Is(result.Error, gorm.ErrRecordNotFound){
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "account not found!",
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
			"error" : "cannot create loan on closed account",
		})

		return
	}

	interestRate := utils.GetInterestRate(req.LoanType)

	if interestRate == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Invalid Loan type",
		})

		return
	}

	var existingLoan models.Loan

	result = database.DB.Where(
		"account_id = ? AND loan_type = ? AND status = ?",
		req.AccountID,
		req.LoanType,
		utils.LoanActive,
	).First(&existingLoan)

	if result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Active loan of this type already exists",
		})
		return
	}

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	
	loan := models.Loan{
		AccountID: req.AccountID,
		LoanType: req.LoanType,
		PrincipalAmount: req.Amount,
		RemainingAmount: req.Amount,
		InterestRate: interestRate,
		Status: utils.LoanActive,
	}
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		
		if err := tx.Create(&loan).Error; err != nil {
			return err
		}

		account.Balance += loan.PrincipalAmount
		if err := tx.Save(&account).Error; err != nil{
			return err
		}

		transaction := models.Transaction{
			AccountID : loan.AccountID,
			TransactionType : utils.TransactionLoanCredit,
			Amount : loan.PrincipalAmount,
		}

		if err := tx.Create(&transaction).Error; err != nil{
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

	res := response.ApplyLoanResponse{
		Message : "loan created successfully",
		LoanID  : loan.LoanID,
		LoanType : loan.LoanType,
		LoanAmount : loan.PrincipalAmount,
		InterestRate : loan.InterestRate,
	}

	c.JSON(http.StatusCreated, res)
}

func GetLoan(c *gin.Context) {

	id := c.Param("id")

	var loan models.Loan

	result := database.DB.First(&loan, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Loan not found",
		})
		return
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, loan)
}

func GetCustomerLoans(c *gin.Context) {

	id := c.Param("id")

	var accounts []models.Account

	result := database.DB.Where(
		"customer_id = ?",
		id,
	).Find(&accounts)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	if len(accounts) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Customer has no accounts",
		})
		return
	}

	accountIDs := make([]uint, 0)

	for _, account := range accounts {
		accountIDs = append(accountIDs, account.AccountID)
	}

	var loans []models.Loan

	result = database.DB.Where(
		"account_id IN ?",
		accountIDs,
	).Find(&loans)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, loans)
}