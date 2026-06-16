package service

import (
	"mysdlc_backend/internal/model"
	"mysdlc_backend/internal/repository"
)

type SDLCService struct {
	Repo *repository.SDLCRepository
}

type SDLCServiceInterface interface {
	CreateSDLC(sdlc *model.SDLC) error
	GetAllSDLCs() ([]model.SDLC, error)
	GetSDLCByID(id uint) (*model.SDLC, error)
	UpdateSDLC(id uint, req *model.UpdateSDLCRequest) (*model.SDLC, error)
	DeleteSDLC(id uint) error

	CreateSDLCStep(step *model.SDLCSteps) error
	UpdateSDLCStep(step *model.SDLCSteps) error
	DeleteSDLCStep(id uint) error
}

// safety check to ensure SDLCService implements SDLCServiceInterface
var _ SDLCServiceInterface = (*SDLCService)(nil)

func (s *SDLCService) CreateSDLC(sdlc *model.SDLC) error {
	return s.Repo.CreateSDLC(sdlc)
}

func (s *SDLCService) GetAllSDLCs() ([]model.SDLC, error) {
	return s.Repo.GetAllSDLCs()
}

func (s *SDLCService) GetSDLCByID(id uint) (*model.SDLC, error) {
	return s.Repo.GetSDLCByID(id)
}

func (s *SDLCService) UpdateSDLC(id uint, req *model.UpdateSDLCRequest) (*model.SDLC, error) {

	sdlc, err := s.Repo.GetSDLCByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		sdlc.Name = *req.Name
	}

	if req.Description != nil {
		sdlc.Description = *req.Description
	}

	if req.Steps != nil {

		for _, stepReq := range *req.Steps {

			for i, step := range sdlc.Steps {

				if step.ID == stepReq.ID {

					if stepReq.Name != nil {
						sdlc.Steps[i].Name = *stepReq.Name
					}

					if stepReq.Description != nil {
						sdlc.Steps[i].Description = *stepReq.Description
					}

					if stepReq.Order != nil {
						sdlc.Steps[i].Order = *stepReq.Order
					}

					if stepReq.IsRequired != nil {
						sdlc.Steps[i].IsRequired = *stepReq.IsRequired
					}

					err := s.Repo.UpdateSDLCStep(&sdlc.Steps[i])
					if err != nil {
						return nil, err
					}
				}
			}
		}
	}

	err = s.Repo.UpdateSDLC(sdlc)
	if err != nil {
		return nil, err
	}

	return sdlc, nil
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