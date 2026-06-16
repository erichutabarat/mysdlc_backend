package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"mysdlc_backend/internal/handler"
	"mysdlc_backend/internal/model"
	"gorm.io/gorm"
)

// Mock Service Implementation
type MockProjectService struct {
	CreateProjectFunc func(uint, *model.CreateProjectRequest) (*model.Project, error)
	GetAllProjectsFunc func(uint) ([]model.ProjectResponse, error)
	GetProjectByIDFunc func(uint, uint) (*model.Project, []model.ProjectPhase, error)
	DeleteProjectByIDFunc func(uint, uint) (*model.Project, error)
	GetProjectMembersFunc func(uint, uint) ([]model.ProjectMemberDTO, error)
}

func (m *MockProjectService) CreateProject(uid uint, req *model.CreateProjectRequest) (*model.Project, error) {
	return m.CreateProjectFunc(uid, req)
}

func (m *MockProjectService) GetAllProjects(uid uint) ([]model.ProjectResponse, error) {
	return m.GetAllProjectsFunc(uid)
}

func (m *MockProjectService) GetProjectByID(uid uint, pid uint) (*model.Project, []model.ProjectPhase, error) {
	return m.GetProjectByIDFunc(uid, pid)
}

func (m *MockProjectService) DeleteProjectByID(uid uint, pid uint) (*model.Project, error) {
	return m.DeleteProjectByIDFunc(uid, pid)
}

func (m *MockProjectService) GetProjectMembers(pid uint, uid uint) ([]model.ProjectMemberDTO, error) {
	return m.GetProjectMembersFunc(pid, uid)
}

// ... Implement other methods as returning nil, nil to satisfy the interface ...
func (m *MockProjectService) AddProjectMember(uint, uint, *model.AddProjectMemberRequest) (*model.ProjectMember, error) { return nil, nil }

func TestCreateProject(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		mockService := &MockProjectService{
			CreateProjectFunc: func(uid uint, req *model.CreateProjectRequest) (*model.Project, error) {
				return &model.Project{Name: req.Name, Description: req.Description, SDLCID: req.SDLCID}, nil
			},
		}
		h := &handler.ProjectHandler{Service: mockService}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("userID", uint(1)) // Simulate Auth Middleware

		reqBody, _ := json.Marshal(model.CreateProjectRequest{Name: "Test Project", Description: "A test project", SDLCID: 1})
		c.Request = httptest.NewRequest("POST", "/projects", bytes.NewBuffer(reqBody))
		c.Request.Header.Set("Content-Type", "application/json")

		h.CreateProject(c)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("Unauthorized_NoUserID", func(t *testing.T) {
		h := &handler.ProjectHandler{}
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		h.CreateProject(c)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Service_Error", func(t *testing.T) {
		mockService := &MockProjectService{
			CreateProjectFunc: func(uint, *model.CreateProjectRequest) (*model.Project, error) {
				return nil, errors.New("database error")
			},
		}
		h := &handler.ProjectHandler{Service: mockService}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("userID", uint(1))
		
		reqBody, _ := json.Marshal(model.CreateProjectRequest{
            Name: "Fail", 
            Description: "Fail",
        })
		
		// Create the request
		req := httptest.NewRequest(http.MethodPost, "/projects", bytes.NewBuffer(reqBody))
		
		// --- CRITICAL FIX: Add this line ---
		req.Header.Set("Content-Type", "application/json")
		// ------------------------------------
		
		c.Request = req
		
		h.CreateProject(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestGetAllProjects(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		mockService := &MockProjectService{
			GetAllProjectsFunc: func(uid uint) ([]model.ProjectResponse, error) {
				return []model.ProjectResponse{
					{ID: 1, Name: "Project 1", Status: model.ProjectActive, Email: "project1@example.com", SDLCName: "SDLC 1"},
					{ID: 2, Name: "Project 2", Status: model.ProjectArchived, Email: "project2@example.com", SDLCName: "SDLC 2"},
				}, nil
			},
		}
		h := &handler.ProjectHandler{Service: mockService}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("userID", uint(1))

		h.GetAllProjects(c)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Unauthorized_NoUserID", func(t *testing.T) {
		h := &handler.ProjectHandler{}
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		h.GetAllProjects(c)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Service_Error", func(t *testing.T) {
		mockService := &MockProjectService{
			GetAllProjectsFunc: func(uint) ([]model.ProjectResponse, error) {
				return nil, errors.New("database error")
			},
		}
		h := &handler.ProjectHandler{Service: mockService}
		
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("userID", uint(1))
		h.GetAllProjects(c)
		
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestGetProjectByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		mockService := &MockProjectService{
			GetProjectByIDFunc: func(uid uint, pid uint) (*model.Project, []model.ProjectPhase, error) {
				return &model.Project{Model: gorm.Model{ID: pid}, Name: "Project 1", Description: "A test project"}, []model.ProjectPhase{
					{Model: gorm.Model{ID: pid}, Name: "Phase 1", Status: model.PhaseActive},
					{Model: gorm.Model{ID: 2}, Name: "Phase 2", Status: model.PhaseComplete},
				}, nil
			},
		}
		h := &handler.ProjectHandler{Service: mockService}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("userID", uint(1))
		c.Params = gin.Params{{Key: "id", Value: "1"}}

		h.GetProjectByID(c)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Unauthorized_NoUserID", func(t *testing.T) {
		h := &handler.ProjectHandler{}
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{{Key: "id", Value: "1"}}

		h.GetProjectByID(c)
		
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Project_NotFound", func(t *testing.T) {
		mockService := &MockProjectService{
			GetProjectByIDFunc: func(uint, uint) (*model.Project, []model.ProjectPhase, error) {
				return nil, nil, gorm.ErrRecordNotFound
			},
		}
		h := &handler.ProjectHandler{Service: mockService}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("userID", uint(1))
		c.Params = gin.Params{{Key: "id", Value: "999"}}
		
		h.GetProjectByID(c)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestDeleteProjectByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		mockService := &MockProjectService{
			DeleteProjectByIDFunc: func(uid uint, pid uint) (*model.Project, error) {
				return &model.Project{Model: gorm.Model{ID: pid}, Name: "Deleted Project"}, nil
			},
		}
		h := &handler.ProjectHandler{Service: mockService}
		
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("userID", uint(1))
		c.Params = gin.Params{{Key: "id", Value: "1"}}
		h.DeleteProjectByID(c)
		
		assert.Equal(t, http.StatusOK, w.Code)
	})
	
	t.Run("Unauthorized_NoUserID", func(t *testing.T) {
		h := &handler.ProjectHandler{}
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{{Key: "id", Value: "1"}}
		h.DeleteProjectByID(c)
		
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Project_NotFound", func(t *testing.T) {
		mockService := &MockProjectService{
			DeleteProjectByIDFunc: func(uint, uint) (*model.Project, error) {
				return nil, gorm.ErrRecordNotFound
			},
		}
		h := &handler.ProjectHandler{Service: mockService}
		
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("userID", uint(1))
		c.Params = gin.Params{{Key: "id", Value: "999"}}
		h.DeleteProjectByID(c)
		
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestProjectMember(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	t.Run("Success", func(t *testing.T) {
		mockService := &MockProjectService{
			GetProjectMembersFunc: func(pid uint, uid uint) ([]model.ProjectMemberDTO, error) {
				return []model.ProjectMemberDTO{
					{ProjectID: pid, UserID: uid, Role: model.RoleOwner, Status: model.MemberAccepted, Email: "user1@gmail.com"},
					{ProjectID: pid, UserID: 2, Role: model.RoleContributor, Status: model.MemberAccepted, Email: "user2@gmail.com"},
				}, nil
			},
		}
		h := &handler.ProjectHandler{Service: mockService}
		
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("userID", uint(1))
		c.Params = gin.Params{{Key: "id", Value: "1"}}
		h.GetProjectMembers(c)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Unauhtorized UserID", func(t *testing.T) {
		h := &handler.ProjectHandler{}
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{{Key: "id", Value: "1"}}
		h.DeleteProjectByID(c)
		
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}