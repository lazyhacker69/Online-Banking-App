package handlers

import (
	"net/http"

	"online-banking/database"
	"online-banking/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"errors"
)

func CreateCustomer(c *gin.Context){
	var customer models.Customer

	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Invalid JSON",
		})
		return
	}
	
	result := database.DB.Create(&customer)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, customer)
}

func GetCustomers(c * gin.Context){

	var customers []models.Customer

	result := database.DB.Find(&customers)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, customers)
}

func GetCustomer(c * gin.Context){

	var customer models.Customer;

	id := c.Param("id")

	result := database.DB.First(&customer, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "Customer not found",
		})

		return
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : result.Error.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, customer)

}
