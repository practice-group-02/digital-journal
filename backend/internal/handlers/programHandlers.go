package handlers

import (
	"digital-journal/internal/models"
	"digital-journal/internal/service"
	"encoding/json"
	"log/slog"
	"net/http"
	"sync"
)

var usersMu sync.Mutex

func PostProgram(w http.ResponseWriter, r *http.Request) {
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	op := "POST /register"

	user := models.NewUser()
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		slog.Error("failed to decode user", "error", err)
		http.Error(w, "Wrong dataType of User", http.StatusBadRequest)
		return
	}

	usersMu.Unlock()
	defer usersMu.Unlock()

	if exists, err := service.UserExists(user.Email, user.Username); err != nil {
		slog.Error("failed to check existance of user", "error", err, "op", op)
		http.Error(w, "Failed to register!", http.StatusInternalServerError)
		return
	} else if exists {
		slog.Error("failed to check existance of user", "op", op)
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
	WriteJSONResponse(w, http.StatusCreated, map[string]string{"message": "User created successfully"})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
}
