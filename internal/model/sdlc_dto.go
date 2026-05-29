package model

type UpdateSDLCRequest struct {
	Name        *string                 `json:"name"`
	Description *string                 `json:"description"`
	Steps       *[]UpdateSDLCStepRequest `json:"steps"`
}

type UpdateSDLCStepRequest struct {
    ID                    uint    `json:"id"`
    Name                  *string `json:"name"`
    Description           *string `json:"description"`
    Order                 *int    `json:"order"`
    IsRequired            *bool   `json:"is_required"`             // was: Required
    EstimatedDurationDays *int    `json:"estimated_duration_days"` // add this too — it's in your model
    AllowedDocTypes       *string `json:"allowed_doc_types"`       // add this too
}

// model/sdlc_dto.go
type SDLCResponse struct {
    ID          uint               `json:"id"`
    Name        string             `json:"name"`
    Description string             `json:"description"`
    IsActive    bool               `json:"is_active"`
    Steps       []SDLCStepResponse `json:"steps,omitempty"`
}

type SDLCStepResponse struct {
    ID                    uint   `json:"id"`
    Name                  string `json:"name"`
    Description           string `json:"description"`
    Order                 int    `json:"order"`
    IsRequired            bool   `json:"is_required"`
    EstimatedDurationDays int    `json:"estimated_duration_days"`
    AllowedDocTypes       string `json:"allowed_doc_types"`
}

func ToSDLCResponse(s SDLC) SDLCResponse {
    resp := SDLCResponse{
        ID:          s.ID,
        Name:        s.Name,
        Description: s.Description,
        IsActive:    s.IsActive,
    }
    for _, step := range s.Steps {
        resp.Steps = append(resp.Steps, SDLCStepResponse{
            ID:                    step.ID,
            Name:                  step.Name,
            Description:           step.Description,
            Order:                 step.Order,
            IsRequired:            step.IsRequired,
            EstimatedDurationDays: step.EstimatedDurationDays,
            AllowedDocTypes:       step.AllowedDocTypes,
        })
    }
    return resp
}