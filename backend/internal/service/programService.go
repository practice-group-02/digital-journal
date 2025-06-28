package service

import (
	"digital-journal/internal/dal"
	"digital-journal/internal/models"
	"fmt"
	"strings"
)

func AddProgramToDB(program *models.Program) error {
	var userID int
	err := dal.DB.QueryRow(`SELECT id FROM users WHERE username = $1`, program.Username).Scan(&userID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return fmt.Errorf("user with username %s not found", program.Username)
		}
		return fmt.Errorf("failed to get user ID: %w", err)
	}

	var typeID int
	err = dal.DB.QueryRow(`SELECT id FROM program_types WHERE name = $1`, program.Type).Scan(&typeID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			err := dal.DB.QueryRow(`INSERT INTO program_types (name) VALUES ($1) RETURNING id`, program.Type).Scan(&typeID)
			if err != nil {
				return fmt.Errorf("failed to add program type to DB: %w", err)
			}
		} else {
			return fmt.Errorf("failed to get program type ID: %w", err)
		}
	}

	query := `
		INSERT INTO programs (title, description, type, country, organization, deadline, link, user_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at, updated_at`
	
	err = dal.DB.QueryRow(query, program.Title, program.Description, typeID, program.Country, program.Organization, program.Deadline, program.Link, userID).
		Scan(&program.ID, &program.CreatedAt, &program.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to add program to DB: %w", err)
	}

	tagIDs := make([]int, 0)
	for _, tag := range program.Tags {
		tag.Name = strings.ToLower(tag.Name)
		var tagID int
		err := dal.DB.QueryRow(`SELECT id FROM tags WHERE name = $1`, tag.Name).Scan(&tagID)
		if err == nil {
			tagIDs = append(tagIDs, tagID)
		} else if err.Error() == "sql: no rows in result set" {
			err := dal.DB.QueryRow(`INSERT INTO tags (name) VALUES ($1) RETURNING id`, tag.Name).Scan(&tagID)
			if err != nil {
				return fmt.Errorf("failed to add tag to DB: %w", err)
			}
			tagIDs = append(tagIDs, tagID)
		} else {
			return fmt.Errorf("failed to check or add tag: %w", err)
		}
	}

	for _, tagID := range tagIDs {
		_, err := dal.DB.Exec(`
			INSERT INTO programs_tags (program_id, tag_id)
			VALUES ($1, $2)`,
			program.ID, tagID)
		if err != nil {
			return fmt.Errorf("failed to add tag to program: %w", err)
		}
	}

	return nil
}

func GetProgramsFromDB() ([]models.Program, error) {
	var programs []models.Program
	query := `
		SELECT p.id, p.title, p.description, pt.name as type, p.country, p.organization, p.deadline, p.link, u.username, p.created_at, p.updated_at
		FROM programs p
		JOIN program_types pt ON p.type = pt.id
		JOIN users u ON p.user_id = u.id
	`
	rows, err := dal.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch programs: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var program models.Program
		if err := rows.Scan(&program.ID, &program.Title, &program.Description, &program.Type, &program.Country, &program.Organization, &program.Deadline, &program.Link, &program.Username, &program.CreatedAt, &program.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan program row: %w", err)
		}

		tags, err := getTagsForProgram(program.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch tags for program %d: %w", program.ID, err)
		}
		program.Tags = tags

		programs = append(programs, program)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	if len(programs) == 0 {
		return nil, fmt.Errorf("programs not found")
	}

	return programs, nil
}


func GetProgramsByTags(tags string) ([]models.Program, error) {
	tags = strings.ToLower(tags)
	
	tagsSep := strings.Split(tags, ",")
	placeholders := make([]string, len(tagsSep))
	for i := range tagsSep {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	query := fmt.Sprintf(`
		SELECT DISTINCT p.id, p.title, p.description, pt.name as type, p.country, p.organization, p.deadline, p.link, u.username, p.created_at, p.updated_at
		FROM programs p
		JOIN programs_tags ptg ON p.id = ptg.program_id
		JOIN tags t ON ptg.tag_id = t.id
		JOIN program_types pt ON p.type = pt.id
		JOIN users u ON p.user_id = u.id
		WHERE t.name IN (%s)`, strings.Join(placeholders, ","))

	params := make([]interface{}, len(tagsSep))
	for i, tag := range tagsSep {
		params[i] = tag
	}

	rows, err := dal.DB.Query(query, params...)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch programs by tags: %w", err)
	}
	defer rows.Close()

	var programs []models.Program
	for rows.Next() {
		var program models.Program
		if err := rows.Scan(&program.ID, &program.Title, &program.Description, &program.Type, &program.Country, &program.Organization, &program.Deadline, &program.Link, &program.Username, &program.CreatedAt, &program.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan program row: %w", err)
		}

		tags, err := getTagsForProgram(program.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch tags for program %d: %w", program.ID, err)
		}
		program.Tags = tags

		programs = append(programs, program)
	}

	if len(programs) == 0 {
		return nil, fmt.Errorf("programs not found")
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return programs, nil
}




func GetProgramsByType(programType string) ([]models.Program, error) {
	var programs []models.Program
	query := `
		SELECT p.id, p.title, p.description, pt.name as type, p.country, p.organization, p.deadline, p.link, u.username, p.created_at, p.updated_at
		FROM programs p
		JOIN program_types pt ON p.type = pt.id
		JOIN users u ON p.user_id = u.id
		WHERE pt.name = $1`

	rows, err := dal.DB.Query(query, programType)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch programs by type: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var program models.Program
		if err := rows.Scan(&program.ID, &program.Title, &program.Description, &program.Type, &program.Country, &program.Organization, &program.Deadline, &program.Link, &program.Username, &program.CreatedAt, &program.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan program row: %w", err)
		}

		tags, err := getTagsForProgram(program.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch tags for program %d: %w", program.ID, err)
		}
		program.Tags = tags

		programs = append(programs, program)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	if len(programs) == 0 {
		return nil, fmt.Errorf("programs not found")
	}

	return programs, nil
}