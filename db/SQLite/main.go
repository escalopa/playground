package main

import (
	"net/http"

	"github.com/escalopa/SQLite/handler"
	"github.com/escalopa/SQLite/service"
	"github.com/escalopa/SQLite/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	database := storage.InitDB()
	defer database.Close()

	// Initialize repositories, services, and handlers
	userRepo := storage.NewUserRepository(database)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Initialize the Gin router
	router := gin.Default()

	// Define routes
	router.GET("/user/", userHandler.GetAllUsers)
	router.POST("/user/", userHandler.CreateUser)

	// Serve the HTML page
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Start the server
	router.Run(":8080")
}
