package repository

import (
	"gorm.io/gorm"
	"mysdlc_backend/internal/model"
)

type ProjectRepository struct {
	DB *gorm.DB
}

func (r *ProjectRepository) Create(project *model.Project) error {
    return r.DB.Create(project).Error
}

func (r *ProjectRepository) Update(project *model.Project) error {
    return r.DB.Save(project).Error
}

func (r *ProjectRepository) GetByID(id uint, ownerID uint) (*model.Project, error) {
    var project model.Project
    err := r.DB.Preload("Owner").
        Preload("SDLC").
        Where("id = ? AND owner_id = ?", id, ownerID).
        First(&project).Error
        
    if err != nil {
        return nil, err
    }
    
    return &project, nil
}

func (r *ProjectRepository) GetAllByOwnerID(ownerID uint) ([]model.Project, map[uint]string, error) {
    var projects []model.Project
    
    // 1. Get Projects
    err := r.DB.Preload("Owner", func(db *gorm.DB) *gorm.DB { return db.Select("id", "email") }).
        Preload("SDLC", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
        Where("owner_id = ?", ownerID).
        Find(&projects).Error
    if err != nil {
        return nil, nil, err
    }

    // 2. Get Phase Names as a Map
    type Result struct {
        ProjectID uint
        Name      string
    }
    var results []Result
    
    // Join projects with phases to get the name associated with the ID
    r.DB.Table("projects").
        Select("projects.id as project_id, project_phases.name").
        Joins("LEFT JOIN project_phases ON project_phases.id = projects.current_phase_id").
        Where("projects.owner_id = ?", ownerID).
        Scan(&results)

    phaseMap := make(map[uint]string)
    for _, res := range results {
        phaseMap[res.ProjectID] = res.Name
    }

    return projects, phaseMap, nil
}

func (r *ProjectRepository) Delete(id uint, ownerID uint) (*model.Project, error) {
	var project model.Project

	err := r.DB.
		Where("id = ? AND owner_id = ?", id, ownerID).
		First(&project).Error

	if err != nil {
		return nil, err
	}

	err = r.DB.Delete(&project).Error
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (r *ProjectRepository) CreatePhases(phases []model.ProjectPhase) error {
    return r.DB.Create(&phases).Error
}

// PHASE REPO
func (r *ProjectRepository) GetPhasesByProjectID(projectID uint) ([]model.ProjectPhase, error) {
    var phases []model.ProjectPhase
    err := r.DB.Where("project_id = ?", projectID).Find(&phases).Error
    return phases, err
}

// PROJECT MEMBER REPO
func (r *ProjectRepository) GetAllMembers(projectID uint) ([]model.ProjectMemberDTO, error) {
    var dtos []model.ProjectMemberDTO

    err := r.DB.Table("project_members").
        Select("project_members.user_id, project_members.project_id, project_members.role, project_members.status, users.email").
        Joins("JOIN users ON users.id = project_members.user_id").
        Where("project_members.project_id = ?", projectID).
        Scan(&dtos).Error

    if err != nil {
        return nil, err
    }

    return dtos, nil
}

func (r *ProjectRepository) AddMember(projectID uint, userId uint, memberData *model.AddProjectMemberRequest) (*model.ProjectMember, error) {
    member := model.ProjectMember{
        ProjectID: projectID,
        UserID:    userId,
        Role:      memberData.Role,
    }
    err := r.DB.Create(&member).Error
    if err != nil {
        return nil, err
    }
    return &member, nil
}

func (r *ProjectRepository) IsUserMember(projectID uint, userID uint) (bool, error) {
    var exists bool
    
    err := r.DB.Table("projects").
        Select("exists (select 1 from projects where id = ? and owner_id = ?) OR exists (select 1 from project_members where project_id = ? and user_id = ?)", 
            projectID, userID, projectID, userID).
        Scan(&exists).Error
        
    return exists, err
}

func (r *ProjectRepository) GetTasksByPhase(projectID uint, phaseID uint) ([]model.Task, error) {
    var tasks []model.Task

    // Using .Where with struct or chainable methods
    err := r.DB.Model(&model.Task{}).
        Where("project_id = ? AND phase_id = ?", projectID, phaseID).
        Order("created_at DESC").
        Find(&tasks).Error

    if err != nil {
        return nil, err
    }

    return tasks, nil
}