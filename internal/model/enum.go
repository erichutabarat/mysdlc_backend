package model

// TaskStatus types
type TaskStatus string

const (
	TaskTodo       TaskStatus = "todo"
	TaskInProgress TaskStatus = "in_progress"
	TaskBlocked    TaskStatus = "blocked"
	TaskDone       TaskStatus = "done"
)

// TaskPriority types
type TaskPriority string

const (
	PriorityLow      TaskPriority = "low"
	PriorityMedium   TaskPriority = "medium"
	PriorityHigh     TaskPriority = "high"
	PriorityCritical TaskPriority = "critical"
)

// ProjectStatus types
type ProjectStatus string

const (
	ProjectActive   ProjectStatus = "active"
	ProjectArchived ProjectStatus = "archived"
	ProjectComplete ProjectStatus = "complete"
)

// PhaseStatus types
type PhaseStatus string

const (
	PhaseLocked   PhaseStatus = "locked"
	PhaseActive   PhaseStatus = "active"
	PhaseComplete PhaseStatus = "complete"
)

// ProjectRole types
type ProjectRole string

const (
	RoleOwner       ProjectRole = "owner"
	RoleContributor ProjectRole = "contributor"
	RoleViewer      ProjectRole = "viewer"
)

// UserRole types
type UserRole string

const (
	UserRoleUser  UserRole = "user"
	UserRoleAdmin UserRole = "admin"
)

// DocType types
type DocType string

const (
	DocPRD        DocType = "PRD"
	DocMockup     DocType = "mockup"
	DocTestScript DocType = "test_script"
	DocOther      DocType = "other"
)