package handlers

import (
	"bytes"
	"digital-journal/internal/models"
	"digital-journal/internal/service"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostProgram(t *testing.T) {
	mockSvc := &service.MockService{}
	svc = mockSvc

	program := models.Program{Title: "Test Program", Description: "Test Description", Type: "тип", Country: "Казахстан", Organization: "Test Org", Username: "user", Tags: []models.Tag{{Name: "Tag1"}}}
	body, _ := json.Marshal(program)

	req := httptest.NewRequest("POST", "/program", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(PostProgram)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestGetPrograms(t *testing.T) {
	mockSvc := &service.MockService{}
	svc = mockSvc

	req := httptest.NewRequest("GET", "/programs", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetPrograms)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetProgramsWithTags(t *testing.T) {
	mockSvc := &service.MockService{}
	svc = mockSvc

	req := httptest.NewRequest("GET", "/programs?tags=бакалавриат", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetProgramsWithTags)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetProgramsOfType(t *testing.T) {
	mockSvc := &service.MockService{}
	svc = mockSvc

	req := httptest.NewRequest("GET", "/programs/грант", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetProgramsOfType)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
