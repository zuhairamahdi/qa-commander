package models

type Defect struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	ProjectID  uint   `json:"project_id"`
	StatusID   int    `json:"status"`
	SeverityID int    `json:"severity"`
	AssignerID uint   `json:"assigner_id"`
	AssigneeID uint   `json:"assignee_id"`
}
