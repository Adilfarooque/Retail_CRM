package handlers

import (
	"net/http"
	"retail_crm/backend/internal/domain/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CustomerHandler struct {
	db *gorm.DB
}

func NewCustomerHandeler(db *gorm.DB) *CustomerHandler {
	return &CustomerHandler{
		db: db,
	}
}

// Get Customers retrives all customers with pagination
func (cu *CustomerHandler) GetCustomers(c *gin.Context) {
	var customers []models.Customer
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	result := cu.db.Offset(offset).Limit(limit).Find(&customers)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch customers",
		})
		return
	}

	//Count of the pagination
	var total int64
	cu.db.Model(&models.Customer{}).Count(&total)

	c.JSON(http.StatusOK, gin.H{
		"customers": customers,
		"total":     total,
		"page":      page,
		"limit":     limit,
	})
}

// Create new customer
func (cu *CustomerHandler) CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cu.db.Create(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}
	c.JSON(http.StatusCreated, customer)
}
