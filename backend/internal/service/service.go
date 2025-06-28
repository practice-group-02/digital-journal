package service

import "digital-journal/internal/models"

type ServiceImpl struct{}

func (s *ServiceImpl) AddProgramToDB(program *models.Program) error {
	return AddProgramToDB(program)
}

func (s *ServiceImpl) GetProgramsFromDB() ([]models.Program, error) {
	return GetProgramsFromDB()
}

func (s *ServiceImpl) GetProgramsByTags(tags string) ([]models.Program, error) {
	return GetProgramsByTags(tags)
}

func (s *ServiceImpl) GetProgramsByType(programType string) ([]models.Program, error) {
	return GetProgramsByType(programType)
}

func (s *ServiceImpl) UpdateProgramInDB(program *models.Program) error {
	return UpdateProgramInDB(program)
}

func (s *ServiceImpl) DeleteProgramFromDB(programID int) error {
	return DeleteProgramFromDB(programID)
}
