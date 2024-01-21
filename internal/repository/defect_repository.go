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

// CreateDefect creates a new defect in the database.
func (dr *DefectRepository) CreateDefect(defect models.Defect) error {
	_, err := dr.DB.Exec(`
		INSERT INTO defects (title, project_id, status_id, severity_id, assignee_id, assigner_id) VALUES ($1, $2, $3, $4, $5, $6)
	`, defect.Title, defect.ProjectID, defect.StatusID, defect.SeverityID, defect.AssigneeID, defect.AssigneeID)
	return err
}

// GetDefectByID retrieves a defect from the database by its ID.
func (dr *DefectRepository) GetDefectByID(defectID uint) (models.Defect, error) {
	var defect models.Defect
	err := dr.DB.QueryRow(`
		SELECT id, title, project_id, status FROM defects WHERE defect_id = $1
	`, defectID).Scan(&defect.ID, &defect.Title, &defect.ProjectID, &defect.StatusID)
	if err != nil {
		return models.Defect{}, errors.New("defect not found")
	}
	return defect, nil
}

// GetDefectsByProject retrieves all defects from the database for a given project.
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
		if err := rows.Scan(&defect.ID, &defect.Title, &defect.ProjectID, &defect.StatusID); err != nil {
			return nil, err
		}
		defects = append(defects, defect)
	}

	return defects, nil
}

// GetDefectsByStatus retrieves all defects from the database for a given status.
func (dr *DefectRepository) UpdateDefect(defect models.Defect) error {
	_, err := dr.DB.Exec(`
		UPDATE defects SET title = $1, project_id = $2, status = $3 WHERE defect_id = $4
	`, defect.Title, defect.ProjectID, defect.StatusID, defect.ID)
	return err
}

// DeleteDefect deletes a defect from the database by its ID.
func (dr *DefectRepository) DeleteDefect(defectID uint) error {
	_, err := dr.DB.Exec(`
		DELETE FROM defects WHERE id = $1
	`, defectID)
	return err
}

// UpdateStatus updates the status of a defect in the database.
func (dr *DefectRepository) UpdateStatus(defectID uint, statusID string) error {
	_, err := dr.DB.Exec(`
		UPDATE defects SET status_id = $1 WHERE defect_id = $2
	`, statusID, defectID)
	return err
}

// UpdateSeverity updates the severity of a defect in the database.
func (dr *DefectRepository) UpdateSeverity(defectID uint, severityID string) error {
	_, err := dr.DB.Exec(`
		UPDATE defects SET severity_id = $1 WHERE defect_id = $2
	`, severityID, defectID)
	return err
}

func (dr *DefectRepository) AddComment(comment models.Comment) error {
	_, err := dr.DB.Exec(`
		INSERT INTO comments (text, defect_id, user_id) VALUES ($1, $2, $3)
	`, comment.Text, comment.DefectID, comment.UserID)
	return err
}

func (dr *DefectRepository) GetCommentsByDefect(defectID uint) ([]models.Comment, error) {
	rows, err := dr.DB.Query(`
		SELECT id, text, defect_id, user_id FROM comments WHERE defect_id = $1
	`, defectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(&comment.ID, &comment.Text, &comment.DefectID, &comment.UserID); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (dr *DefectRepository) GetCommentByID(commentID uint) (models.Comment, error) {
	var comment models.Comment
	err := dr.DB.QueryRow(`
		SELECT id, text, defect_id, user_id FROM comments WHERE id = $1
	`, commentID).Scan(&comment.ID, &comment.Text, &comment.DefectID, &comment.UserID)
	if err != nil {
		return models.Comment{}, errors.New("comment not found")
	}
	return comment, nil
}

func (dr *DefectRepository) UpdateComment(comment models.Comment) error {
	_, err := dr.DB.Exec(`
		UPDATE comments SET text = $1 WHERE id = $2
	`, comment.Text, comment.ID)
	return err
}
