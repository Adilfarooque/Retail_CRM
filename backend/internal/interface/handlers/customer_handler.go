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
func (h *CustomerHandler) GetCustomers(c *gin.Context) {
	var customers []models.Customer
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	result := h.db.Offset(offset).Limit(limit).Find(&customers)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch customers",
		})
		return
	}

	//Count of the pagination
	var total int64
	h.db.Model(&models.Customer{}).Count(&total)

	c.JSON(http.StatusOK, gin.H{
		"customers": customers,
		"total":     total,
		"page":      page,
		"limit":     limit,
	})
}

// Create new customer
func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.Create(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}
	c.JSON(http.StatusCreated, customer)
}

// Update existing customer
func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer

	if err := h.db.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.Save(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update customer"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer

	if err := h.db.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	if err := h.db.Delete(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete customer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}
