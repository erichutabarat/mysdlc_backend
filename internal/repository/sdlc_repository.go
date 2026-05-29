package repository

import (
	"mysdlc_backend/internal/model"
	"gorm.io/gorm"
)

type SDLCRepository struct {
	DB *gorm.DB
}

func (r *SDLCRepository) CreateSDLC(sdlc *model.SDLC) error {
	return r.DB.Create(sdlc).Error
}

func (r *SDLCRepository) GetAllSDLCs() ([]model.SDLC, error) {
	var sdlcs []model.SDLC
	err := r.DB.Preload("Steps").Find(&sdlcs).Error
	return sdlcs, err
}

func (r *SDLCRepository) GetSDLCByID(id uint) (*model.SDLC, error) {
	var sdlc model.SDLC
	err := r.DB.Preload("Steps").First(&sdlc, id).Error
	if err != nil {
		return nil, err
	}
	return &sdlc, nil
}

func (r *SDLCRepository) UpdateSDLC(sdlc *model.SDLC) error {
	return r.DB.Save(sdlc).Error
}

func (r *SDLCRepository) DeleteSDLC(id uint) error {
	return r.DB.Delete(&model.SDLC{}, id).Error
}

func (r *SDLCRepository) CreateSDLCStep(step *model.SDLCSteps) error {
	return r.DB.Create(step).Error
}

func (r *SDLCRepository) UpdateSDLCStep(step *model.SDLCSteps) error {
	return r.DB.Save(step).Error
}

func (r *SDLCRepository) DeleteSDLCStep(id uint) error {
	return r.DB.Delete(&model.SDLCSteps{}, id).Error
}