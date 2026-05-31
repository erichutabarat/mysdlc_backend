package model

type CreateProjectRequest struct {
	Name        string `json:"name"        binding:"required,min=3,max=150"`
	Description string `json:"description" binding:"omitempty,max=500"`
	SDLCID      uint   `json:"sdlc_id"     binding:"required"`
}

type ProjectResponse struct {
    ID    uint   `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    SDLCName string `json:"sdlc_name"`
	CurrentPhase string `json:"current_phase,omitempty"`
}