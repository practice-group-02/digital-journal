package handlers

import (
	"digital-journal/internal/models"
	"digital-journal/internal/service"
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"
)

func PostProgram(w http.ResponseWriter, r *http.Request) {
	op := "POST /program"

	program := &models.Program{}
	err := json.NewDecoder(r.Body).Decode(program)
	if err != nil {
		slog.Error("failed to decode program", "error", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(program.Title) == "" || strings.TrimSpace(program.Description) == "" || program.Type == "" || strings.TrimSpace(program.Country) == "" || strings.TrimSpace(program.Organization) == "" || program.Username == "" || len(program.Tags) == 0 {
		slog.Error("missing required fields", "op", op)
		http.Error(w, "Title, description, type, country, organization, tags, and username are required", http.StatusBadRequest)
		return
	}

	err = service.AddProgramToDB(program)
	if err != nil {
		slog.Error("failed to add program to DB", "error", err, "op", op)
		http.Error(w, "Failed to add program", http.StatusInternalServerError)
		return
	}

	slog.Info("created program successfully", "op", op)
	WriteJSONResponse(w, http.StatusCreated, map[string]string{"message": "Program created successfully"})
}


func GetPrograms(w http.ResponseWriter, r *http.Request) {
	programs, err := service.GetProgramsFromDB()
	if err != nil {
		if err.Error() == "programs not found" {
			slog.Error("failed to get programs by tags", "error", err)
			http.Error(w, "Programs not found", http.StatusNotFound)
			return
		}
		slog.Error("failed to get programs by tags", "error", err)
		http.Error(w, "Failed to fetch programs with tags", http.StatusInternalServerError)
		return
	}

	WriteJSONResponse(w, http.StatusOK, programs)
}

func GetProgramsWithTags(w http.ResponseWriter, r *http.Request) {
	tags := r.URL.Query().Get("tags")
	if tags == "" {
		http.Error(w, "No tags provided", http.StatusBadRequest)
		return
	}

	programs, err := service.GetProgramsByTags(tags)
	if err != nil {
		if err.Error() == "programs not found" {
			slog.Error("failed to get programs by tags", "error", err)
			http.Error(w, "Programs not found with provided tags", http.StatusNotFound)
			return
		}
		slog.Error("failed to get programs by tags", "error", err)
		http.Error(w, "Failed to fetch programs with tags", http.StatusInternalServerError)
		return
	}

	WriteJSONResponse(w, http.StatusOK, programs)
}
