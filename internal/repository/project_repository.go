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