package server

import (
	"database/sql"

	"qa_commander/internal/api/handlers"
	"qa_commander/internal/repository"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures the Gin router with API routes.
func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	// Initialize repository and other dependencies
	projectRepo := repository.NewProjectRepository(db)
	defectRepo := repository.NewDefectRepository(db)
	userRepo := repository.NewUserRepository(db)

	// Create handlers
	projectHandler := handlers.NewProjectHandler(projectRepo)
	defectHandler := handlers.NewDefectHandler(defectRepo)
	userHandler := handlers.NewUserHandler(userRepo)

	// Define API routes
	api := r.Group("/api")
	{
		projects := api.Group("/projects")
		{
			projects.GET("/", projectHandler.GetProjects)
			projects.POST("/create", projectHandler.CreateProject)
		}

		defects := api.Group("/defects")
		{
			defects.GET("/", defectHandler.GetDefects)
			defects.POST("/create", defectHandler.CreateDefect)
		}

		users := api.Group("/users")
		{
			users.POST("/create", userHandler.CreateUser)
			users.POST("/login", userHandler.Login)
			users.POST("/comment", userHandler.WriteComment)
		}
	}

	return r
}
