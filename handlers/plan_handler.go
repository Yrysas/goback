package handlers

import (
	"finance_tracker/config"
	"finance_tracker/models"

	"github.com/gin-gonic/gin"
)

func GetPlans(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var plans []models.Plan
	config.DB.Where("user_id = ?", userID).Find(&plans)

	type PlanResponse struct {
		Title     string `json:"title"`
		Price     float64 `json:"price"`
		DueDate   string `json:"due_date"`
		CreatedAt string `json:"created_at"`
	}

	var responses []PlanResponse
	for _, p := range plans {
		responses = append(responses, PlanResponse{
			Title:     p.Title,
			Price:     p.Price,
			DueDate:   p.DueDate.Format("2006-01-02T15:04:05Z07:00"),
			CreatedAt: p.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	c.JSON(200, responses)
}

func CreatePlan(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var plan models.Plan
	c.BindJSON(&plan)
	plan.UserID = userID.(uint)

	config.DB.Create(&plan)
	c.JSON(200, gin.H{"message": "Added"})
}

func GetPlan(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := c.Param("id")

	var plan models.Plan
	config.DB.Where("id = ? AND user_id = ?", id, userID).First(&plan)

	c.JSON(200, plan)
}

func UpdatePlan(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := c.Param("id")

	var plan models.Plan
	config.DB.Where("id = ? AND user_id = ?", id, userID).First(&plan)

	c.BindJSON(&plan)
	config.DB.Save(&plan)

	c.JSON(200, plan)
}

func DeletePlan(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := c.Param("id")

	config.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Plan{})

	c.JSON(200, gin.H{"message": "deleted"})
}
