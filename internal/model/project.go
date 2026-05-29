package model

import (
	"time"
	"gorm.io/gorm"
) 

type Project struct {
    gorm.Model
    Name        string        `gorm:"type:varchar(150);not null"              json:"name"`
    Description string        `gorm:"type:text"                               json:"description"`
    Status      ProjectStatus `gorm:"type:varchar(20);not null;default:'active'" json:"status"`

    OwnerID uint  `gorm:"not null;index"     json:"owner_id"`
    Owner   *User `gorm:"foreignKey:OwnerID" json:"owner,omitempty"`

    SDLCID uint  `gorm:"not null;index"     json:"sdlc_id"`
    SDLC   *SDLC `gorm:"foreignKey:SDLCID"  json:"sdlc,omitempty"`

    // just the FK column — no association field, no constraint problem
    CurrentPhaseID *uint `gorm:"index" json:"current_phase_id"`
}

// ProjectPhase is a snapshot clone of an SDLCStep at project creation time.
type ProjectPhase struct {
    gorm.Model
    Name                  string      `gorm:"type:varchar(100);not null"             json:"name"`
    Order                 int         `gorm:"not null"                               json:"order"`
    Status                PhaseStatus `gorm:"type:varchar(20);not null;default:'locked'" json:"status"`
    IsRequired            bool        `gorm:"column:is_required;not null;default:true" json:"is_required"`
    EstimatedDurationDays int         `gorm:"default:0"                              json:"estimated_duration_days"`
    AllowedDocTypes       string      `gorm:"type:varchar(255)"                      json:"allowed_doc_types"`
    UnlockedAt            *time.Time  `json:"unlocked_at"`
    CompletedAt           *time.Time  `json:"completed_at"`

    ProjectID  uint       `gorm:"not null;index" json:"project_id"`
    SDLCStepID *uint      `gorm:"index"          json:"sdlc_step_id"`
}