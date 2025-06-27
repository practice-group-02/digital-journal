package service

import (
	"digital-journal/internal/dal"
	"digital-journal/internal/models"
	"fmt"
)


func getTagsForProgram(programID int) ([]models.Tag, error) {
	var tags []models.Tag
	query := `
		SELECT t.id, t.name
		FROM tags t
		JOIN programs_tags pt ON t.id = pt.tag_id
		WHERE pt.program_id = $1
	`
	rows, err := dal.DB.Query(query, programID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tags for program: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var tag models.Tag
		if err := rows.Scan(&tag.ID, &tag.Name); err != nil {
			return nil, fmt.Errorf("failed to scan tag row: %w", err)
		}
		tags = append(tags, tag)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return tags, nil
}
