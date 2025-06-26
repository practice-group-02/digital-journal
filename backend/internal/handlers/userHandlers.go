package handlers

import (
	"digital-journal/internal/models"
	"digital-journal/internal/service"
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

var usersMu sync.Mutex

func CreateUser(w http.ResponseWriter, r *http.Request) {
	op := "POST /register"

	user := models.NewUser()
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		slog.Error("failed to decode user", "error", err)
		http.Error(w, "Wrong dataType of User", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.PasswordHashed) == "" {
		slog.Error("username, email and password are required", "op", op)
		http.Error(w, "Username, email and password are required!", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHashed), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("failed to hash user password", "error", err)
		http.Error(w, "Failed to regiester", http.StatusInternalServerError)
		return
	}
	user.PasswordHashed = string(hashedPassword)

	usersMu.Lock()
	defer usersMu.Unlock()

	if exists, err := service.UserExists(user.Email, user.Username); err != nil {
		slog.Error("failed to check existance of user", "error", err, "op", op)
		http.Error(w, "Failed to register!", http.StatusInternalServerError)
		return
	} else if exists {
		slog.Error("user already exists", "op", op)
		http.Error(w, "User already exists!", http.StatusConflict)
		return
	}

	err = service.AddUserToDB(user)
	if err != nil {
		if err == service.ErrEmailExists || err == service.ErrUsernameExists {
			slog.Error("bad request of existing user details", "op", op)
			http.Error(w, err.Error(), http.StatusConflict)
			return
		} else {
			slog.Error("failed to add user to DB", "error", err, "op", op)
			http.Error(w, "Failed to register!", http.StatusInternalServerError)
			return
		}
	}
	slog.Info("created user successfully", "op", op)
	WriteJSONResponse(w, http.StatusCreated, map[string]string{"message": "User created successfully"})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	op := "POST /login"
	
	user := models.NewUser()
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		slog.Error("failed to decode user", "error", err)
		http.Error(w, "Wrong dataType of User", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.PasswordHashed) == "" {
		slog.Error("username, email and password are required", "op", op)
		http.Error(w, "Username, email and password are required!", http.StatusBadRequest)
		return
	}

	usersMu.Lock()
	defer usersMu.Unlock()

	if exists, err := service.UserExists(user.Email, user.Username); err != nil {
		slog.Error("failed to check existance of user", "error", err, "op", op)
		http.Error(w, "Failed to login!", http.StatusInternalServerError)
		return
	} else if !exists {
		slog.Error("user not found", "op", op)
		http.Error(w, "User not exists!", http.StatusNotFound)
		return
	}

	err = service.VerifyPassword(user.Email, user.PasswordHashed)
	if err != nil {
		slog.Error("failed to verify password", "error", err, "op", op)
		http.Error(w, "Wrong credentials!", http.StatusBadRequest)
		return 
	}

	slog.Info("successfully logged in", "op", op)
	WriteJSONResponse(w, http.StatusOK, map[string]string{"message": "Successful login"})
}
