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
		INSERT INTO users (username, password_hash, email, created_at, active) VALUES ($1, $2, $3, $4, $5)
	`, user.Username, user.PasswordHash, user.Email, user.CreatedAt, user.Active)
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
		SELECT user_id, username, password_hash FROM users WHERE username = $1
	`, username).Scan(&user.ID, &user.Username, &user.PasswordHash)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

// HashPassword hashes the given password using bcrypt
func (ur *UserRepository) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePasswordHash compares a plain text password with its hashed value
func (ur *UserRepository) ComparePasswordHash(plainPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

func (ur *UserRepository) GenerateJWT(userID uint) (string, error) {
	// Generate a JWT token
	return "", nil
}

func (ur *UserRepository) IsUsernameOrEmailExists(username string, email string) bool {
	var count int
	err := ur.DB.QueryRow(`
		SELECT COUNT(*) FROM users WHERE username = $1 OR email = $2
	`, username, email).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}
