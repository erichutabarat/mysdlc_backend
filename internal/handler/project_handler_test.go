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
)

// Mock Service Implementation
type MockProjectService struct {
	CreateProjectFunc func(uint, *model.CreateProjectRequest) (*model.Project, error)
}

func (m *MockProjectService) CreateProject(uid uint, req *model.CreateProjectRequest) (*model.Project, error) {
	return m.CreateProjectFunc(uid, req)
}

// ... Implement other methods as returning nil, nil to satisfy the interface ...
func (m *MockProjectService) GetAllProjects(uint) ([]model.ProjectResponse, error) { return nil, nil }
func (m *MockProjectService) GetProjectByID(uint, uint) (*model.Project, []model.ProjectPhase, error) { return nil, nil, nil }
func (m *MockProjectService) DeleteProjectByID(uint, uint) (*model.Project, error) { return nil, nil }
func (m *MockProjectService) GetProjectMembers(uint, uint) ([]model.ProjectMemberDTO, error) { return nil, nil }
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