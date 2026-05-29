package service

import (
	"mysdlc_backend/internal/model"
	"mysdlc_backend/internal/repository"
)

type SDLCService struct {
	Repo *repository.SDLCRepository
}

func (s *SDLCService) CreateSDLC(sdlc *model.SDLC) error {
	return s.Repo.CreateSDLC(sdlc)
}

func (s *SDLCService) GetAllSDLCs() ([]model.SDLC, error) {
	return s.Repo.GetAllSDLCs()
}

func (s *SDLCService) GetSDLCByID(id uint) (*model.SDLC, error) {
	return s.Repo.GetSDLCByID(id)
}

func (s *SDLCService) UpdateSDLC(sdlc *model.SDLC) error {
	return s.Repo.UpdateSDLC(sdlc)
}

func (s *SDLCService) DeleteSDLC(id uint) error {
	return s.Repo.DeleteSDLC(id)
}

func (s *SDLCService) CreateSDLCStep(step *model.SDLCSteps) error {
	return s.Repo.CreateSDLCStep(step)
}

func (s *SDLCService) UpdateSDLCStep(step *model.SDLCSteps) error {
	return s.Repo.UpdateSDLCStep(step)
}

func (s *SDLCService) DeleteSDLCStep(id uint) error {
	return s.Repo.DeleteSDLCStep(id)
}