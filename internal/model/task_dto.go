package model

import (
	"time"
)

type AddTaskRequest struct {
    Title       string        `json:"title" binding:"required"`
    Description string        `json:"description"`
    AssigneeID  *uint         `json:"assignee_id"`
    Priority    *TaskPriority `json:"priority" binding:"omitempty,oneof=low medium high"`
    Status      *TaskStatus   `json:"status" binding:"omitempty,oneof=todo in_progress done"`
    DueDate     *time.Time    `json:"due_date"`
}