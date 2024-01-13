package handlers

import (
	"net/http"
	"qa_commander/internal/models"
	"qa_commander/internal/repository"
	"time"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	ProjectRepo *repository.ProjectRepository
}

func NewProjectHandler(projectRepo *repository.ProjectRepository) *ProjectHandler {
	return &ProjectHandler{
		ProjectRepo: projectRepo,
	}
}

func (ph *ProjectHandler) GetProjects(c *gin.Context) {
	projects, err := ph.ProjectRepo.GetProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, projects)
}

func (ph *ProjectHandler) CreateProject(c *gin.Context) {
	var project models.Project

	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	project.OwnerID = uint(c.MustGet("user_id").(float64))
	//convert date to string
	if project.StartDate != "" {
		_, err := time.Parse("2006-01-02", project.StartDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format yyyy-mm-dd"})
			return
		}
	}

	if project.EndDate != "" {
		_, err := time.Parse("2006-01-02", project.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format yyyy-mm-dd"})
			return
		}
	}
	model, err := ph.ProjectRepo.CreateProjectModel(project.Name, project.Description, project.OwnerID, project.StartDate, project.EndDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := ph.ProjectRepo.CreateProject(model); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, project)
}
