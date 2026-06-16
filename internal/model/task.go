package model

import (
	"time"
	"gorm.io/gorm"
)

type Task struct {
    gorm.Model
    PhaseID        uint         `gorm:"not null;index"              json:"phase_id"`
    ProjectID      uint         `gorm:"not null;index"              json:"project_id"`
    Title          string       `gorm:"type:varchar(200);not null"  json:"title"`
    Description    string       `gorm:"type:text"                   json:"description"`
    AssigneeID     *uint        `gorm:"index"                       json:"assignee_id"`
    Assignee       *User        `gorm:"foreignKey:AssigneeID"       json:"assignee,omitempty"`
    Priority       TaskPriority `gorm:"type:varchar(20);not null;default:'medium'" json:"priority"`
    Status         TaskStatus   `gorm:"type:varchar(20);not null;default:'todo'"   json:"status"`
    DueDate        *time.Time   `json:"due_date"`
    FromTemplateID *uint        `gorm:"index"                       json:"from_template_id"` // nullable
}