package model

type UpdateSDLCRequest struct {
	Name        *string                 `json:"name"`
	Description *string                 `json:"description"`
	Steps       *[]UpdateSDLCStepRequest `json:"steps"`
}

type UpdateSDLCStepRequest struct {
	ID          uint    `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Order       *int    `json:"order"`
	Required    *bool   `json:"required"`
}