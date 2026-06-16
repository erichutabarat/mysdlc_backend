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
    UserRepo *repository.UserRepository
}

type ProjectServiceInterface interface {
	CreateProject(ownerID uint, req *model.CreateProjectRequest) (*model.Project, error)
	GetAllProjects(ownerID uint) ([]model.ProjectResponse, error)
	GetProjectByID(id uint, ownerID uint) (*model.Project, []model.ProjectPhase, error)
	DeleteProjectByID(id uint, userID uint) (*model.Project, error)
	GetProjectMembers(projectID uint, userID uint) ([]model.ProjectMemberDTO, error)
	AddProjectMember(projectID uint, userID uint, newMemberData *model.AddProjectMemberRequest) (*model.ProjectMember, error)
    GetTasks(projectID uint, userID uint, phaseID uint) ([]model.Task, error)
}

// safety check to ensure ProjectService implements ProjectServiceInterface
var _ ProjectServiceInterface = (*ProjectService)(nil)

var ErrUnauthorized = errors.New("unauthorized: only project owner can view members")

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

func (s *ProjectService) GetAllProjects(ownerID uint) ([]model.ProjectResponse, error) {
    projects, phaseMap, err := s.Repo.GetAllByOwnerID(ownerID)
    if err != nil {
        return nil, err
    }

    var response []model.ProjectResponse
    for _, p := range projects {
        response = append(response, model.ProjectResponse{
            ID:   p.ID,
            Name: p.Name,
            Status: p.Status,
            Email: p.Owner.Email,
            SDLCName: p.SDLC.Name,
            CurrentPhase: phaseMap[p.ID],
        })
    }

    return response, nil
}

func (s *ProjectService) GetProjectByID(
    id uint,
    ownerID uint,
) (*model.Project, []model.ProjectPhase, error) {

    project, err := s.Repo.GetByID(id, ownerID)
    if err != nil {
        return nil, nil, err
    }

    phases, err := s.Repo.GetPhasesByProjectID(id)
    if err != nil {
        return nil, nil, err
    }

    return project, phases, nil
}

func (s *ProjectService) DeleteProjectByID(id uint, userID uint) (*model.Project, error) {
	return s.Repo.Delete(id, userID)
}

// PROJECT MEMBER SERVICE
func (s *ProjectService) GetProjectMembers(projectID uint, userID uint) ([]model.ProjectMemberDTO, error) {
	project, err := s.Repo.GetByID(projectID, userID)
	if err != nil {
		return nil, err
	}

	if project.OwnerID != userID {
		return nil, ErrUnauthorized
	}

	return s.Repo.GetAllMembers(projectID)
}

func (s *ProjectService) AddProjectMember(projectID uint, userID uint, newMemberData *model.AddProjectMemberRequest) (*model.ProjectMember, error) {
    project, err := s.Repo.GetByID(projectID, userID)
    if err != nil {
        return nil, err
    }

    if project.OwnerID != userID {
        return nil, ErrUnauthorized
    }

    // Check and get the user by email
    newMember, err := s.UserRepo.GetUserByEmail(newMemberData.Email)
    if err != nil {
        return nil, errors.New("user with this email does not exist")
    }
    var projectMember *model.ProjectMember
    projectMember, err = s.Repo.AddMember(projectID, newMember.ID, newMemberData)
    if err != nil {
        return nil, err
    }

    return projectMember, nil
}

func (s *ProjectService) GetTasks(projectID uint, userID uint, phaseID uint) ([]model.Task, error) {
    isMember, err := s.Repo.IsUserMember(projectID, userID)
    if err != nil {
        return nil, err
    }
    if !isMember {
        return nil, errors.New("user is not a member of this project")
    }

    var tasks[]model.Task
    tasks, err = s.Repo.GetTasksByPhase(projectID, phaseID)
    return tasks , err
}