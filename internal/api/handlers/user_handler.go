package handlers

import (
	"fmt"
	"qa_commander/internal/repository"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserRepo *repository.UserRepository
}

func NewUserHandler(userRepo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		UserRepo: userRepo,
	}
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	fmt.Println("CreateUser")
	// Handler logic to create a user...
}

func (uh *UserHandler) Login(c *gin.Context) {
	// Handler logic for user login...
}

func (uh *UserHandler) WriteComment(c *gin.Context) {
	// Handler logic to allow a user to write a comment...
}
