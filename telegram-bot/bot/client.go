package bot

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"telegram-bot/config"
// 	"telegram-bot/models"
// )

// func GetLatestPrograms() ([]models.Program, error) {
// 	url := fmt.Sprintf("%s/programs", config.BackendURL)

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	var programs []models.Program
// 	if err := json.NewDecoder(resp.Body).Decode(&programs); err != nil {
// 		return nil, err
// 	}

// 	return programs, nil
// }


import (
	"telegram-bot/models"
)

func GetLatestPrograms() ([]models.Program, error) {
	return []models.Program{
		{
			ID:          1,
			Title:       "Стипендия Chevening",
			Country:     "Великобритания",
			Deadline:    "2025-09-15",
			Description: "Финансирование магистратуры в UK.",
		},
		{
			ID:          2,
			Title:       "Стипендия Erasmus+",
			Country:     "Европа",
			Deadline:    "2025-10-01",
			Description: "Международный обмен и обучение в ЕС.",
		},
	}, nil
}
