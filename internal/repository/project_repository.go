package repository

import (
	"database/sql"
	"errors"
	"qa_commander/internal/models"
	"time"
)

type ProjectRepository struct {
	DB *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{DB: db}
}

func (pr *ProjectRepository) CreateProject(project models.Project) error {
	_, err := pr.DB.Exec(`
		INSERT INTO projects (name, description, owner_id, created_at) VALUES ($1, $2, $3, $4)
	`, project.Name, project.Description, project.OwnerID, project.CreatedAt)
	return err
}

func (pr *ProjectRepository) CreateProjectModel(
	project_name string, project_description string,
	project_owner_id uint,
	start_date string,
	end_date string,
) (models.Project, error) {
	var project models.Project = models.Project{
		Name:        project_name,
		Description: project_description,
		OwnerID:     project_owner_id,
		CreatedAt:   time.Now().Format(time.RFC3339),
		StartDate:   start_date,
		EndDate:     end_date,
	}

	return project, nil
}

func (pr *ProjectRepository) GetProjectByID(projectID uint) (models.Project, error) {
	var project models.Project
	err := pr.DB.QueryRow(`
		SELECT project_id, name FROM projects WHERE id = $1
	`, projectID).Scan(&project.ID, &project.Name)
	if err != nil {
		return models.Project{}, errors.New("project not found")
	}
	return project, nil
}

func (pr *ProjectRepository) GetProjects() ([]models.Project, error) {
	rows, err := pr.DB.Query(`
		SELECT project_id, name FROM projects
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var project models.Project
		if err := rows.Scan(&project.ID, &project.Name); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	return projects, nil
}

func (pr *ProjectRepository) UpdateProject(project models.Project) error {
	_, err := pr.DB.Exec(`
		UPDATE projects SET name = $1 WHERE id = $2
	`, project.Name, project.ID)
	return err
}

func (pr *ProjectRepository) DeleteProject(projectID uint) error {
	_, err := pr.DB.Exec(`
		DELETE FROM projects WHERE id = $1
	`, projectID)
	return err
}
