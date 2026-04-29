package main

import (
	"finance_tracker/config"
	"finance_tracker/handlers"
	"finance_tracker/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	// Note: Migrations are now handled via golang-migrate tool
	// Run 'make migrate-up' to apply migrations
	// config.DB.AutoMigrate(
	// 	&models.User{},
	// 	&models.Transaction{},
	// 	&models.Plan{},
	// )

	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())

	auth.GET("/transactions", handlers.GetTransactions)
	auth.POST("/transactions", handlers.CreateTransaction)
	auth.GET("/transactions/:id", handlers.GetTransaction)
	auth.PUT("/transactions/:id", handlers.UpdateTransaction)
	auth.DELETE("/transactions/:id", handlers.DeleteTransaction)

	auth.GET("/plans", handlers.GetPlans)
	auth.POST("/plans", handlers.CreatePlan)
	auth.GET("/plans/:id", handlers.GetPlan)
	auth.PUT("/plans/:id", handlers.UpdatePlan)
	auth.DELETE("/plans/:id", handlers.DeletePlan)

	r.Run(":8080")
}