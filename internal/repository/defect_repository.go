package repository

import (
	"database/sql"
	"errors"
	"qa_commander/internal/models"
)

type DefectRepository struct {
	DB *sql.DB
}

func NewDefectRepository(db *sql.DB) *DefectRepository {
	return &DefectRepository{DB: db}
}

func (dr *DefectRepository) CreateDefect(defect models.Defect) error {
	_, err := dr.DB.Exec(`
		INSERT INTO defects (title, project_id, status) VALUES ($1, $2, $3)
	`, defect.Title, defect.ProjectID, defect.Status)
	return err
}

func (dr *DefectRepository) GetDefectByID(defectID uint) (models.Defect, error) {
	var defect models.Defect
	err := dr.DB.QueryRow(`
		SELECT id, title, project_id, status FROM defects WHERE id = $1
	`, defectID).Scan(&defect.ID, &defect.Title, &defect.ProjectID, &defect.Status)
	if err != nil {
		return models.Defect{}, errors.New("defect not found")
	}
	return defect, nil
}

func (dr *DefectRepository) GetDefectsByProject(projectID uint) ([]models.Defect, error) {
	rows, err := dr.DB.Query(`
		SELECT id, title, project_id, status FROM defects WHERE project_id = $1
	`, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var defects []models.Defect
	for rows.Next() {
		var defect models.Defect
		if err := rows.Scan(&defect.ID, &defect.Title, &defect.ProjectID, &defect.Status); err != nil {
			return nil, err
		}
		defects = append(defects, defect)
	}

	return defects, nil
}

func (dr *DefectRepository) UpdateDefect(defect models.Defect) error {
	_, err := dr.DB.Exec(`
		UPDATE defects SET title = $1, project_id = $2, status = $3 WHERE id = $4
	`, defect.Title, defect.ProjectID, defect.Status, defect.ID)
	return err
}

func (dr *DefectRepository) DeleteDefect(defectID uint) error {
	_, err := dr.DB.Exec(`
		DELETE FROM defects WHERE id = $1
	`, defectID)
	return err
}
