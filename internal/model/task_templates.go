package model

import "gorm.io/gorm"

type TaskTemplate struct {
    gorm.Model
    SDLCStepID  uint         `gorm:"not null;index"              json:"sdlc_step_id"`
    SDLCStep    *SDLCSteps   `gorm:"foreignKey:SDLCStepID"       json:"sdlc_step,omitempty"`
    Title       string       `gorm:"type:varchar(200);not null"  json:"title"`
    Description string       `gorm:"type:text"                   json:"description"`
    Priority    TaskPriority `gorm:"type:varchar(20);default:'medium'" json:"priority"`
    IsDefault   bool         `gorm:"default:true"                json:"is_default"`
    Order       int          `gorm:"default:0"                   json:"order"`
}