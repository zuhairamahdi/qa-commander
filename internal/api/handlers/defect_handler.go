package handlers

import (
	"net/http"
	"qa_commander/internal/models"
	"qa_commander/internal/repository"

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

	if err := ph.ProjectRepo.CreateProject(project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, project)
}
