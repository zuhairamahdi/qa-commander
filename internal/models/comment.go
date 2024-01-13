package models

import "time"

type Comment struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	ProjectID uint      `json:"project_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
