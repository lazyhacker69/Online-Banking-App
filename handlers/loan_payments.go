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

func RepayLoan(c *gin.Context) {

	var req request.RepayLoanRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	if req.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid amount",
		})
		return
	}

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

	if loan.Status != utils.LoanActive {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Loan is already paid",
		})
		return
	}

	if req.Amount > loan.RemainingAmount {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Repayment exceeds remaining loan amount",
		})
		return
	}

	var account models.Account

	result = database.DB.First(&account, loan.AccountID)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	if account.Balance < req.Amount {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Insufficient balance",
		})
		return
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {

		account.Balance -= req.Amount

		if err := tx.Save(&account).Error; err != nil {
			return err
		}

		loan.RemainingAmount -= req.Amount

		if loan.RemainingAmount == 0 {
			loan.Status = utils.LoanPaid
		}

		if err := tx.Save(&loan).Error; err != nil {
			return err
		}

		payment := models.LoanPayment{
			LoanID: loan.LoanID,
			Amount: req.Amount,
		}

		if err := tx.Create(&payment).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	res := response.RepayLoanResponse{
		Message:           "Loan repayment successful",
		RemainingAmount:   loan.RemainingAmount,
		LoanStatus:        loan.Status,
	}

	c.JSON(http.StatusOK, res)
}