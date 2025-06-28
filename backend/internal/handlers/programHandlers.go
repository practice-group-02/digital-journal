package handlers

import (
	"digital-journal/internal/models"
	"digital-journal/internal/service"
	"encoding/json"
	"log/slog"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var svc service.Service = &service.ServiceImpl{}

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

	err = svc.AddProgramToDB(program)
	if err != nil {
		slog.Error("failed to add program to DB", "error", err, "op", op)
		http.Error(w, "Failed to add program", http.StatusInternalServerError)
		return
	}

	slog.Info("created program successfully", "op", op)
	WriteJSONResponse(w, http.StatusCreated, map[string]string{"message": "Program created successfully"})
}

func GetPrograms(w http.ResponseWriter, r *http.Request) {
	programs, err := svc.GetProgramsFromDB()
	if err != nil {
		if err.Error() == "programs not found" {
			slog.Error("failed to get programs", "error", err)
			http.Error(w, "Programs not found", http.StatusNotFound)
			return
		}
		slog.Error("failed to get programs", "error", err)
		http.Error(w, "Failed to fetch programs", http.StatusInternalServerError)
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

	programs, err := svc.GetProgramsByTags(tags)
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

func GetProgramsOfType(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	// Используем регулярное выражение для извлечения параметра "Type"
	re := regexp.MustCompile(`/programs/([^/]+)`)
	matches := re.FindStringSubmatch(path)

	if len(matches) < 2 {
		http.Error(w, "Type is required", http.StatusBadRequest)
		return
	}

	programType := matches[1]

	programType = strings.ToLower(programType)
	programs, err := svc.GetProgramsByType(programType)
	if err != nil {
		if err.Error() == "programs not found" {
			slog.Error("failed to get programs with type", "error", err)
			http.Error(w, "Programs not found", http.StatusNotFound)
			return
		}
		slog.Error("failed to get programs with type", "error", err)
		http.Error(w, "Failed to fetch programs with type", http.StatusInternalServerError)
		return
	}

	WriteJSONResponse(w, http.StatusOK, programs)
}


func UpdateProgram(w http.ResponseWriter, r *http.Request) {
	op := "PUT /program/{Id}"

	idStr := r.PathValue("Id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		slog.Error("invalid program ID", "error", err, "op", op)
		http.Error(w, "Invalid program ID", http.StatusBadRequest)
		return
	}

	program := &models.Program{}
	err = json.NewDecoder(r.Body).Decode(program)
	if err != nil {
		slog.Error("failed to decode program", "error", err, "op", op)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	program.ID = id

	err = svc.UpdateProgramInDB(program)
	if err != nil {
		slog.Error("failed to update program in DB", "error", err, "op", op)
		http.Error(w, "Failed to update program", http.StatusInternalServerError)
		return
	}

	slog.Info("updated program successfully", "op", op)
	WriteJSONResponse(w, http.StatusOK, map[string]string{"message": "Program updated successfully"})
}

func DeleteProgram(w http.ResponseWriter, r *http.Request) {
	op := "DELETE /program/{Id}"

	idStr := r.PathValue("Id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		slog.Error("invalid program ID", "error", err, "op", op)
		http.Error(w, "Invalid program ID", http.StatusBadRequest)
		return
	}

	err = svc.DeleteProgramFromDB(id)
	if err != nil {
		slog.Error("failed to delete program from DB", "error", err, "op", op)
		http.Error(w, "Failed to delete program", http.StatusInternalServerError)
		return
	}

	slog.Info("deleted program successfully", "op", op)
	WriteJSONResponse(w, http.StatusOK, map[string]string{"message": "Program deleted successfully"})
}
