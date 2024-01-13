package handlers

import (
	"fmt"
	"net/http"
	"qa_commander/internal/models"
	"qa_commander/internal/repository"
	validators "qa_commander/internal/validator"
	"time"

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
	userRegister := new(models.UserRegister)
	if err := c.ShouldBindJSON(&userRegister); err != nil {
		c.JSON(400, gin.H{"error": "invalid user register data format"})
		return
	}
	// Check if the username or email already exists
	if uh.UserRepo.IsUsernameOrEmailExists(userRegister.Username, userRegister.Email) {
		c.JSON(400, gin.H{"error": "username or email already exists"})
		return
	}
	// check if password is valid
	if !validators.IsPasswordComplex(userRegister.Password) {
		c.JSON(400, gin.H{"error": "password is not valid"})
		return
	}
	if !validators.IsEmailValid(userRegister.Email) {
		c.JSON(400, gin.H{"error": "email is not valid"})
		return
	}
	// Hash the password with the salt
	hashedPassword, err := uh.UserRepo.HashPassword(userRegister.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// Create a new user
	user := models.User{
		Username:     userRegister.Username,
		PasswordHash: hashedPassword,
		Email:        userRegister.Email,
		CreatedAt:    time.Now().Format(time.RFC3339),
		Active:       false,
	}
	// Save the user to the database
	if err := uh.UserRepo.CreateUser(user); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("user %s created", user.Username)})
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
	var isValidPassword bool = uh.UserRepo.ComparePasswordHash(userLogin.Password, user.PasswordHash)
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
