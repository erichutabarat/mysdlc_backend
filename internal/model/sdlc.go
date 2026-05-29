package model

import "gorm.io/gorm"

type SDLC struct {
    gorm.Model
    Name        string      `gorm:"type:varchar(100);not null" json:"name"`
    Description string      `gorm:"type:text" json:"description"`
    
    Steps       []SDLCSteps `gorm:"foreignKey:SDLCID;constraint:OnDelete:CASCADE" json:"steps"`
}

type SDLCSteps struct {
    gorm.Model
    Name        string `gorm:"type:varchar(100);not null" json:"name"`
    Description string `gorm:"type:text" json:"description"`
    Order       int    `gorm:"not null" json:"order"`
    Required    bool   `gorm:"not null" json:"required"`
    
    // Foreign Key
    SDLCID      uint   `gorm:"not null" json:"sdlc_id"`
}