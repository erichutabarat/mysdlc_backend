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