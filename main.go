package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tlksimba519/ptcg_trainerweb/controllers"
	"github.com/tlksimba519/ptcg_trainerweb/database"
	"github.com/tlksimba519/ptcg_trainerweb/middlewares"
)

func main() {
	database.Init_mongo()
	database.Init()
	database.Seed()
	r := gin.Default()

	// CORS
	r.Use(middlewares.CORSMiddleware())

	// index page
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from backend"})
	})

	r.GET("/tournament", controllers.RegisterTournament)

	// Auth Routes
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)

	// Public card API
	r.GET("/cards", controllers.GetAllCards)
	r.GET("/cards/:id", controllers.GetCardByID)

	// Protected Routes
	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/users", controllers.GetUsers)
		protected.GET("/users/:id", controllers.GetUserByID)

		protected.GET("/decks", controllers.GetDecks)
		protected.GET("/decks/:id", controllers.GetDeckByID)
		protected.POST("/decks", controllers.CreateDeck)
		protected.PUT("/decks/:id", controllers.UpdateDeck)
		protected.DELETE("/decks/:id", controllers.DeleteDeck)
	}

	r.Run(":8080")
}
