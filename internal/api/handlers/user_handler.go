package handlers

import (
	"fmt"
	"qa_commander/internal/models"
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
	// Read the username and password from the request body
	userLogin := new(models.UserLogin)
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// Get the user from the database
	user, err := uh.UserRepo.GetUserByUsername(userLogin.Username)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//comapre the password with the hashed password
	var isValidPassword bool = uh.UserRepo.HashAndComparePassword(userLogin.Password, user.PasswordHash)
	if !isValidPassword {
		c.JSON(400, gin.H{"error": "invalid password"})
		return
	}
	// Generate a JWT token
	token, err := uh.UserRepo.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": token})
	return

}

func (uh *UserHandler) WriteComment(c *gin.Context) {
	// Handler logic to allow a user to write a comment...
}
