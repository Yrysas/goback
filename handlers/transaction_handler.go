package handlers

import (
	"finance_tracker/config"
	"finance_tracker/models"

	"github.com/gin-gonic/gin"
)

func GetTransactions(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var transactions []models.Transaction
	config.DB.Where("user_id = ?", userID).Find(&transactions)

	type TransactionResponse struct {
		Title     string `json:"title"`
		Price     float64 `json:"price"`
		CreatedAt string `json:"created_at"`
	}

	var responses []TransactionResponse
	for _, t := range transactions {
		responses = append(responses, TransactionResponse{
			Title:     t.Title,
			Price:     t.Price,
			CreatedAt: t.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	c.JSON(200, responses)
}

func CreateTransaction(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var transaction models.Transaction
	c.BindJSON(&transaction)
	transaction.UserID = userID.(uint)

	config.DB.Create(&transaction)
	c.JSON(200, gin.H{"message": "Added"})
}

func GetTransaction(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := c.Param("id")

	var transaction models.Transaction
	config.DB.Where("id = ? AND user_id = ?", id, userID).First(&transaction)

	c.JSON(200, transaction)
}

func UpdateTransaction(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := c.Param("id")

	var transaction models.Transaction
	config.DB.Where("id = ? AND user_id = ?", id, userID).First(&transaction)

	c.BindJSON(&transaction)
	config.DB.Save(&transaction)

	c.JSON(200, transaction)
}

func DeleteTransaction(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := c.Param("id")

	config.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Transaction{})

	c.JSON(200, gin.H{"message": "deleted"})
}
