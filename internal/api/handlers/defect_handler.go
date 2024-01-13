package handlers

import (
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
}
