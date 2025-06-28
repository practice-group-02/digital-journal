package service

import "digital-journal/internal/models"

type Service interface {
    AddProgramToDB(program *models.Program) error
    GetProgramsFromDB() ([]models.Program, error)
    GetProgramsByTags(tags string) ([]models.Program, error)
    GetProgramsByType(programType string) ([]models.Program, error)
    UpdateProgramInDB(program *models.Program) error
    DeleteProgramFromDB(programID int) error 
}