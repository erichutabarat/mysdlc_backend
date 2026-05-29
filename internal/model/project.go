package model

import (
	"time"
	"gorm.io/gorm"
) 

type Project struct {
    gorm.Model
    Name        string        `gorm:"type:varchar(150);not null"   json:"name"`
    Description string        `gorm:"type:text"                    json:"description"`
    Status      ProjectStatus `gorm:"type:varchar(20);not null;default:'active'" json:"status"`

    // Who owns this project
    OwnerID uint  `gorm:"not null;index"               json:"owner_id"`
    Owner   *User `gorm:"foreignKey:OwnerID"           json:"owner,omitempty"`

    // Which SDLC methodology was chosen
    SDLCID uint  `gorm:"not null;index"               json:"sdlc_id"`
    SDLC   *SDLC `gorm:"foreignKey:SDLCID"            json:"sdlc,omitempty"`

    // Pointer to whichever ProjectPhase is currently active
    // NULL until project is created and first phase is unlocked
    CurrentPhaseID *uint         `gorm:"index"                        json:"current_phase_id"`
    CurrentPhase   *ProjectPhase `gorm:"foreignKey:CurrentPhaseID"    json:"current_phase,omitempty"`

    // All cloned phases for this project
    Phases []ProjectPhase `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE" json:"phases,omitempty"`
}

// ProjectPhase is a snapshot clone of an SDLCStep at project creation time.
// Admin changes to SDLCStep do NOT affect existing ProjectPhases.
type ProjectPhase struct {
    gorm.Model
    Name                  string      `gorm:"type:varchar(100);not null"  json:"name"`
    Order                 int         `gorm:"not null"                    json:"order"`
    Status                PhaseStatus `gorm:"type:varchar(20);not null;default:'locked'" json:"status"`
    IsRequired            bool        `gorm:"not null;default:true"       json:"is_required"`
    EstimatedDurationDays int         `gorm:"default:0"                   json:"estimated_duration_days"`
    AllowedDocTypes       string      `gorm:"type:varchar(255)"           json:"allowed_doc_types"`
    UnlockedAt            *time.Time  `json:"unlocked_at"`
    CompletedAt           *time.Time  `json:"completed_at"`

    // FK back to the project
    ProjectID uint     `gorm:"not null;index"              json:"project_id"`
    Project   *Project `gorm:"foreignKey:ProjectID"        json:"-"` // avoid circular JSON

    // FK back to the original SDLCStep (for traceability only — read-only reference)
    SDLCStepID *uint     `gorm:"index"                       json:"sdlc_step_id"`
    SDLCStep   *SDLCSteps `gorm:"foreignKey:SDLCStepID"       json:"sdlc_step,omitempty"`
}