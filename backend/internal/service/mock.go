package service

import "digital-journal/internal/models"

type MockService struct{}

func (m *MockService) AddProgramToDB(program *models.Program) error {
	return nil
}

func (m *MockService) GetProgramsFromDB() ([]models.Program, error) {
	return []models.Program{
		{ID: 1, Title: "Test Program", Description: "Test Description", Type: "тип", Country: "Казахстан", Organization: "Test Org", Username: "user"},
	}, nil
}

func (m *MockService) GetProgramsByTags(tags string) ([]models.Program, error) {
	return []models.Program{
		{ID: 1, Title: "Test Program", Description: "Test Description", Type: "тип", Country: "Казахстан", Organization: "Test Org", Username: "user"},
	}, nil
}

func (m *MockService) GetProgramsByType(programType string) ([]models.Program, error) {
	return []models.Program{
		{ID: 1, Title: "Test Program", Description: "Test Description", Type: programType, Country: "Казахстан", Organization: "Test Org", Username: "user"},
	}, nil
}
