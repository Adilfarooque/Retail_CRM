package models

import (
	"time"

	"gorm.io/gorm"
)

type CustomerPreferences struct {
	PreferredCategories []string `json:"preferredCategories" gorm:"type:text[]"`
	PreferredBrands     []string `json:"preferredBrands" gorm:"type:text[]"`
	CommunicationPrefs  string   `json:"communicationPrefs" gorm:"type:varchar(50);default:'email'"`
	SpecialOffers       bool     `json:"specialOffers" gorm:"default:true"`
	Newsletter          bool     `json:"newsletter" gorm:"default:true"`
}

type Customer struct {
	gorm.Model
	FirstName   string    `json:"firstName" gorm:"type:varchar(100);not null"`
	LastName    string    `json:"lastName" gorm:"type:varchar(100);not null"`
	Email       string    `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Phone       string    `json:"phone" gorm:"type:varchar(20)"`
	Birthday    time.Time `json:"birthday"`
	Anniversary time.Time `json:"anniversary"`

	// Address
	Address    string `json:"address" gorm:"type:text"`
	City       string `json:"city" gorm:"type:varchar(100)"`
	State      string `json:"state" gorm:"type:varchar(100)"`
	PostalCode string `json:"postalCode" gorm:"type:varchar(20)"`
	Country    string `json:"country" gorm:"type:varchar(100)"`

	// Loyalty program
	LoyaltyPoints    int       `json:"loyaltyPoints" gorm:"default:0"`
	LoyaltyTier      string    `json:"loyaltyTier" gorm:"type:varchar(50);default:'bronze'"`
	PointsExpiry     time.Time `json:"pointsExpiry"`
	LastPointsEarned time.Time `json:"lastPointsEarned"`

	// Customer Status
	Status        string    `json:"status" gorm:"type:varchar(20);default:'active'"`
	LastPurchased time.Time `json:"lastPurchased"`
	LastVisited   time.Time `json:"lastVisited"`

	// Customer Preferences
	Preferences CustomerPreferences `json:"preferences" gorm:"embedded"`

	// Purachase History
	PurchaseHistory []Purchase     `json:"purchaseHistory,omitempty" gorm:"foreignKey:CustomerID"`
	SpecialDates    []SpecialDate  `json:"specialDates,omitempty" gorm:"foreignKey:CustomerID"`
	LoyaltyHistory  []LoyaltyEvent `json:"loyaltyHistory,omitempty" gorm:"foreignKey:CustomerID"`
}

type Purchase struct {
	gorm.Model
	CustomerID    uint           `json:"customerId" gorm:"not null"`
	OrderNumber   string         `json:"orderNumber" gorm:"type:varchar(50);unique"`
	TotalAmount   float64        `json:"totalAmount" gorm:"type:decimal(10,2)"`
	PointsEarned  int            `json:"pointsEarned"`
	PurchaseDate  time.Time      `json:"purchaseDate"`
	Status        string         `json:"status" gorm:"type:varchar(20)"`
	PaymentMethod string         `json:"paymentMethod" gorm:"type:varchar(50)"`
	Items         []PurchaseItem `json:"items,omitempty" gorm:"foreignKey:PurchaseID"`
}

type PurchaseItem struct {
	gorm.Model
	PurchaseID  uint    `json:"purchaseId" gorm:"not null"`
	ProductID   uint    `json:"productId" gorm:"not null"`
	ProductName string  `json:"productName" gorm:"type:varchar(255)"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unitPrice" gorm:"type:decimal(10,2)"`
	Subtotal    float64 `json:"subtotal" gorm:"type:decimal(10,2)"`
}

type SpecialDate struct {
	gorm.Model
	CustomerID  uint      `json:"customerId" gorm:"not null"`
	Occasion    string    `json:"occasion" gorm:"type:varchar(100)"`
	Date        time.Time `json:"date"`
	IsRecurring bool      `json:"isRecurring" gorm:"default:true"`
	Reminder    bool      `json:"reminder" gorm:"default:true"`
	Notes       string    `json:"notes" gorm:"type:text"`
}

type LoyaltyEvent struct {
	gorm.Model
	CustomerID  uint      `json:"customerId" gorm:"not null"`
	EventType   string    `json:"eventType" gorm:"type:varchar(50)"` // earned, redeemed, expired
	Points      int       `json:"points"`
	Description string    `json:"description" gorm:"type:text"`
	EventDate   time.Time `json:"eventDate"`
	ExpiryDate  time.Time `json:"expiryDate"`
}
