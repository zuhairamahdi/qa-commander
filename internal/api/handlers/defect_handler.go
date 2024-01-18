package handlers

import (
	"net/http"
	"qa_commander/internal/models"
	"qa_commander/internal/repository"

	"github.com/gin-gonic/gin"
)

type DefectHandler struct {
	DefectRepo *repository.DefectRepository
}

func NewDefectHandler(defectRepo *repository.DefectRepository) *DefectHandler {
	return &DefectHandler{
		DefectRepo: defectRepo,
	}
}

func (dh *DefectHandler) GetDefects(c *gin.Context) {
	// Handler logic to retrieve defects...
}

func (dh *DefectHandler) CreateDefect(c *gin.Context) {
	// Handler logic to create a defect...
	var defect models.Defect

	if err := c.ShouldBindJSON(&defect); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defect.AssigneeID = uint(c.MustGet("user_id").(float64))

	err := dh.DefectRepo.CreateDefect(defect)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, defect)

}

func (dh *DefectHandler) UpdateDefect(c *gin.Context) {
	// Handler logic to update a defect...
}

func (dh *DefectHandler) DeleteDefect(c *gin.Context) {
	// Handler logic to delete a defect...
}

func (dh *DefectHandler) AddComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comment.UserID = uint(c.MustGet("user_id").(float64))
	err := dh.DefectRepo.AddComment(comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comment)

}

func (dh *DefectHandler) UpdateComment(c *gin.Context) {
	// Handler logic to update a comment to a defect...
}
