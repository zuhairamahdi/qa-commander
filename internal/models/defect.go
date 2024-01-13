package models

type Defect struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	ProjectID uint   `json:"project_id"`
	Status    string `json:"status"`
}
