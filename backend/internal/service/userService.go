package service

import (
	"digital-journal/internal/dal"
	"digital-journal/internal/models"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUsernameExists = errors.New("Username already exists!")
	ErrEmailExists    = errors.New("Email already exists!")
)

func UserExists(email, username string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1 OR username = $2)`
	err := dal.DB.QueryRow(query, email, username).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func AddUserToDB(user *models.User) error {
	if strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.Password) == "" {
		return errors.New("username, email and password are required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	if user.Role == "" {
		user.Role = "user"
	}

	query := `
			INSERT INTO users (username, email, password_hash, role)
			VALUES ($1, $2, $3, $4)
			RETURNING id, created_at`

	err = dal.DB.QueryRow(query, user.Username, user.Email, string(hashedPassword), user.Role).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			if strings.Contains(err.Error(), "username") {
				return ErrUsernameExists
			}
			if strings.Contains(err.Error(), "email") {
				return ErrEmailExists
			}
		}
		return err
	}

	return nil
}
