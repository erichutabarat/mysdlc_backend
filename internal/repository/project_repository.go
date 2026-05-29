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

func (r *ProjectRepository) GetByID(id uint) (*model.Project, error) {
    var project model.Project
    err := r.DB.Preload("Owner").
        Preload("SDLC").
        First(&project, id).Error
    if err != nil {
        return nil, err
    }
    return &project, nil
}

func (r *ProjectRepository) GetAllByOwnerID(ownerID uint) ([]model.Project, error) {
    var projects []model.Project
    err := r.DB.Preload("Owner").
        Preload("SDLC").
        Where("owner_id = ?", ownerID).
        Find(&projects).Error
    return projects, err
}

func (r *ProjectRepository) Delete(id uint) error {
	return r.DB.Delete(&model.Project{}, id).Error
}

// repository/project_repository.go
func (r *ProjectRepository) CreatePhases(phases []model.ProjectPhase) error {
    return r.DB.Create(&phases).Error
}