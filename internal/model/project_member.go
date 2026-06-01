package model

type ProjectMember struct {
	ProjectID uint `gorm:"primaryKey;autoIncrement:false"`
	UserID    uint `gorm:"primaryKey;autoIncrement:false"`

	Role      ProjectRole `gorm:"type:enum('owner', 'contributor', 'viewer');not null"`
	Status MemberStatus `gorm:"type:enum('pending','accepted','rejected');default:'pending'"`
}

type AddProjectMemberRequest struct {
	Email string `json:"email" binding:"required,email"`
	Role  ProjectRole `json:"role" binding:"required,oneof=owner contributor viewer"`
}

type Invitation struct {
	ProjectID   uint   `json:"project_id"`
	UserID    uint   `json:"user_id"`
	Role        ProjectRole `json:"role"`
	Status      MemberStatus `json:"status"`
}

type RespondInvitationRequest struct {
    ProjectID uint `json:"project_id" binding:"required"`
    Status    MemberStatus `json:"status" binding:"required,oneof=pending accepted rejected"`
}