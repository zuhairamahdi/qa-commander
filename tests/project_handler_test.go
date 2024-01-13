// tests/project_handler_test.go

package tests

import (
	"net/http"
	"net/http/httptest"
	"qa_commander/internal/api/handlers"
	"qa_commander/internal/repository"
	"qa_commander/internal/server"
	"testing"
)

func TestGetProjects(t *testing.T) {
	// Initialize a test database
	db := setupTestDB()
	defer db.Close()

	// Initialize a project repository with the test database
	projectRepo := repository.NewProjectRepository(db)

	// Initialize a project handler with the test repository
	projectHandler := handlers.NewProjectHandler(projectRepo)

	// Create a test HTTP request
	req, err := http.NewRequest("GET", "/api/projects", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a test HTTP response recorder
	w := httptest.NewRecorder()

	// Create a test Gin router with the project handler
	r := server.SetupRouter(db)
	r.GET("/api/projects", projectHandler.GetProjects)

	// Serve the HTTP request to the recorder
	r.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check the response body or any other assertions based on your application logic
}

func TestCreateProject(t *testing.T) {
	// Initialize a test database
	db := setupTestDB()
	defer db.Close()

	// Initialize a project repository with the test database
	projectRepo := repository.NewProjectRepository(db)

	// Initialize a project handler with the test repository
	projectHandler := handlers.NewProjectHandler(projectRepo)

	// Create a test HTTP request
	req, err := http.NewRequest("POST", "/api/projects/create", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a test HTTP response recorder
	w := httptest.NewRecorder()

	// Create a test Gin router with the project handler
	r := server.SetupRouter(db)
	r.POST("/api/projects/create", projectHandler.CreateProject)

	// Serve the HTTP request to the recorder
	r.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}

	// Check the response body or any other assertions based on your application logic
}
