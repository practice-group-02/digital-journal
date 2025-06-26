package service

import (
	"database/sql"
	"digital-journal/internal/dal"
	"digital-journal/internal/models"
	"errors"
	"fmt"
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
        if errors.Is(err, sql.ErrNoRows) {
            return false, nil
        }
        return false, fmt.Errorf("database error checking user existence: %w", err)
    }
    return exists, nil
}

func AddUserToDB(user *models.User) error {
	if user.Role == "" {
		user.Role = "user"
	}

	query := `
			INSERT INTO users (username, email, password_hash, role)
			VALUES ($1, $2, $3, $4)
			RETURNING id, created_at`

	err := dal.DB.QueryRow(query, user.Username, user.Email, user.PasswordHashed, user.Role).Scan(&user.ID, &user.CreatedAt)
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


func VerifyPassword(email, password string) error {
	var hashedPassword string
	
	query := `SELECT password_hash FROM users WHERE email = $1`
	err := dal.DB.QueryRow(query, email).Scan(&hashedPassword)
	
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) { 
			return errors.New("invalid credentials")
		}
		return fmt.Errorf("database error: %w", err)
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return errors.New("invalid password")
		}
		return fmt.Errorf("password comparison error: %w", err)
	}
	
	return nil
}