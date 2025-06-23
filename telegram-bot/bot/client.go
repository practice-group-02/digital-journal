package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"telegram-bot/config"
	"telegram-bot/models"
)

func GetLatestPrograms() ([]models.Program, error) {
	url := fmt.Sprintf("%s/programs", config.BackendURL)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var programs []models.Program
	if err := json.NewDecoder(resp.Body).Decode(&programs); err != nil {
		return nil, err
	}

	return programs, nil
}
