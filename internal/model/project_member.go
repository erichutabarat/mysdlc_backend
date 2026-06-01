package model

type ProjectMember struct {
	ProjectID uint `gorm:"primaryKey;autoIncrement:false"`
	UserID    uint `gorm:"primaryKey;autoIncrement:false"`

	Role      ProjectRole `gorm:"type:enum('owner', 'contributor', 'viewer');not null"`
	Status MemberStatus `gorm:"type:enum('pending','accepted','rejected');default:'pending'"`

	User      User  `json:"user" gorm:"foreignKey:UserID"`
}

type Invitation struct {
	ProjectID   uint   `json:"project_id"`
	UserID    uint   `json:"user_id"`
	Role        ProjectRole `json:"role"`
	Status      MemberStatus `json:"status"`
}