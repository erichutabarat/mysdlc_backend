package model

type RespondInvitationRequest struct {
    ProjectID uint `json:"project_id" binding:"required"`
    Status    MemberStatus `json:"status" binding:"required,oneof=pending accepted rejected"`
}

type AddProjectMemberRequest struct {
	Email string `json:"email" binding:"required,email"`
	Role  ProjectRole `json:"role" binding:"required,oneof=owner contributor viewer"`
}

type ProjectMemberDTO struct {
	UserID    uint        `json:"user_id"`
	ProjectID uint        `json:"project_id"`
	Role      ProjectRole `json:"role"`
	Status    MemberStatus `json:"status"`
	
	Email  string      `json:"email"`
}
