package models

import "time"

type Comment struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	DefectID  uint      `json:"defect_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
