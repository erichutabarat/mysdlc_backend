package model

type CreateProjectRequest struct {
	Name        string `json:"name"        binding:"required,min=3,max=150"`
	Description string `json:"description" binding:"omitempty,max=500"`
	SDLCID      uint   `json:"sdlc_id"     binding:"required"`
}