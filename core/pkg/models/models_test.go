package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectInitialization(t *testing.T) {
	// Create a new project
	project := Project{
		Name:   "TestProject",
		Email:  "test@example.com",
		Folder: "/path/to/project",
	}

	// Use assertions to check the values
	assert.Equal(t, "TestProject", project.Name, "Name should be 'TestProject'")
	assert.Equal(t, "test@example.com", project.Email, "Email should be 'test@example.com'")
	assert.Equal(t, "/path/to/project", project.Folder, "Folder should be '/path/to/project'")
}

func TestProjectDefaultInitialization(t *testing.T) {
	// Create a new project with default values
	project := Project{}

	// Use assertions to check the default values
	assert.Equal(t, "", project.Name, "Name should be an empty string")
	assert.Equal(t, "", project.Email, "Email should be an empty string")
	assert.Equal(t, "", project.Folder, "Folder should be an empty string")
}
