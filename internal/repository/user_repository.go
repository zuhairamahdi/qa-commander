package repository

import (
	"database/sql"
	"errors"
	"qa_commander/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) CreateUser(user models.User) error {
	_, err := ur.DB.Exec(`
		INSERT INTO users (username, password_hash) VALUES ($1, $2)
	`, user.Username, user.PasswordHash)
	return err
}

func (ur *UserRepository) GetUserByID(userID uint) (models.User, error) {
	var user models.User
	err := ur.DB.QueryRow(`
		SELECT id, username, password_hash FROM users WHERE id = $1
	`, userID).Scan(&user.ID, &user.Username, &user.PasswordHash)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

func (ur *UserRepository) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := ur.DB.QueryRow(`
		SELECT id, username, password_hash FROM users WHERE username = $1
	`, username).Scan(&user.ID, &user.Username, &user.PasswordHash)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

func (ur *UserRepository) HashPassword(password string) (string, error) {
	// Hash the password with the salt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// Compare the password with the hashed password
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (ur *UserRepository) CheckPassword(password string, hashedPassword string) bool {
	// Compare the password with the hashed password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (ur *UserRepository) HashAndComparePassword(password string, userHashedPassword string) bool {
	// Hash the password with the salt

	// Compare the password with the hashed password
	return ur.CheckPassword(password, userHashedPassword)
}

func (ur *UserRepository) GenerateJWT(userID uint) (string, error) {
	// Generate a JWT token
	return "", nil
}
