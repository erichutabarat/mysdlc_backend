// model/sdlc.go
package model

import "gorm.io/gorm"

type SDLC struct {
    gorm.Model
    Name        string      `gorm:"type:varchar(100);not null" json:"name"`
    Description string      `gorm:"type:text"                               json:"description"`
    IsActive    bool        `gorm:"not null;default:true"                   json:"is_active"`
    CreatedByID uint        `gorm:"not null"                                json:"created_by_id"`
    CreatedBy   User        `gorm:"foreignKey:CreatedByID"                  json:"created_by,omitempty"`
    Steps       []SDLCSteps  `gorm:"foreignKey:SDLCID;constraint:OnDelete:CASCADE" json:"steps,omitempty"`
}

type SDLCSteps struct {
    gorm.Model
    Name                 string   `gorm:"type:varchar(100);not null"  json:"name"`
    Description          string   `gorm:"type:text"                   json:"description"`
    Order                int      `gorm:"not null"                    json:"order"`
    IsRequired           bool     `gorm:"not null;default:true"       json:"is_required"`
    EstimatedDurationDays int     `gorm:"default:0"                   json:"estimated_duration_days"`
    AllowedDocTypes      string   `gorm:"type:varchar(255)"           json:"allowed_doc_types"` // store as comma-separated or JSON string
    // FK
    SDLCID               uint     `gorm:"not null;index"              json:"sdlc_id"`
    SDLC                 *SDLC    `gorm:"foreignKey:SDLCID"           json:"sdlc,omitempty"`
}