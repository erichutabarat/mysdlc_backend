package service

import (
	"fmt"
	"errors"
	"mysdlc_backend/internal/model"
	"mysdlc_backend/internal/repository"
)

type ProjectService struct {
	Repo *repository.ProjectRepository
	SDLCRepo *repository.SDLCRepository
}

func (s *ProjectService) CreateProject(ownerID uint, req *model.CreateProjectRequest) (*model.Project, error) {
    // 1. Check SDLC exists and is active
    sdlc, err := s.SDLCRepo.GetSDLCByID(req.SDLCID)
    if err != nil {
        return nil, errors.New("SDLC not found")
    }
    if !sdlc.IsActive {
        return nil, errors.New("SDLC is inactive")
    }

    // 2. Build and save Project first (no phases yet)
    project := &model.Project{
        Name:        req.Name,
        Description: req.Description,
        OwnerID:     ownerID,
        SDLCID:      req.SDLCID,
        Status: model.ProjectActive,
    }
	fmt.Printf("Creating project: %+v\n", project)

    if err := s.Repo.Create(project); err != nil {
        return nil, err
    }

    // 3. Clone SDLC steps into ProjectPhases using the new project.ID
    var phases []model.ProjectPhase
    for i, step := range sdlc.Steps {
        status := model.PhaseLocked
        if i == 0 {
            status = model.PhaseActive
        }
        stepID := step.ID
        phases = append(phases, model.ProjectPhase{
            Name:                  step.Name,
            Order:                 step.Order,
            Status:                status,
            IsRequired:            step.IsRequired,
            EstimatedDurationDays: step.EstimatedDurationDays,
            AllowedDocTypes:       step.AllowedDocTypes,
            ProjectID:             project.ID,  // ← set now that project.ID exists
            SDLCStepID:            &stepID,
        })
    }

    // 4. Save all phases
    if err := s.Repo.CreatePhases(phases); err != nil {
        return nil, err
    }

    // 5. Set CurrentPhaseID to the first phase
    if len(phases) > 0 {
        firstPhaseID := phases[0].ID
        project.CurrentPhaseID = &firstPhaseID
        if err := s.Repo.Update(project); err != nil {
            return nil, err
        }
    }

    return project, nil
}

func (s *ProjectService) GetAllProjects(ownerID uint) ([]model.Project, error) {
	return s.Repo.GetAllByOwnerID(ownerID)
}

func (s *ProjectService) GetProjectByID(id uint) (*model.Project, error) {
	return s.Repo.GetByID(id)
}

func (s *ProjectService) DeleteProjectByID(id uint) error {
	return s.Repo.Delete(id)
}