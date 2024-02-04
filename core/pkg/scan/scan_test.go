package scan

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDotFilePath(t *testing.T) {
	dotFilePath := GetDotFilePath()
	if dotFilePath == "" {
		t.Error("Expected a non-empty dot file path, got an empty string.")
	}
}

func TestOpenFile(t *testing.T) {
	filePath := "testfile.txt"
	defer func() {
		_ = os.Remove(filePath)
	}()
	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	file.Close()

	// Open the file
	f := openFile(filePath)
	defer f.Close()
	assert.NotNil(t, f, "Expected a non-nil file object, got nil.")
}

func TestParseFileLinesToSlice(t *testing.T) {
	filePath := "testfile.txt"
	defer func() {
		_ = os.Remove(filePath)
	}()

	content := "line1\nline2\nline3"
	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	assert.NoError(t, err)
	lines := ParseFileLinesToSlice(filePath)
	expectedLines := []string{"line1", "line2", "line3"}
	assert.Equal(t, expectedLines, lines, fmt.Sprintf("Expected %d lines, got %d", len(expectedLines), len(lines)))
	for i, line := range lines {
		assert.Equal(t, expectedLines[i], line, fmt.Sprintf("Expected line '%s', got '%s'", expectedLines[i], line))
	}
}

func TestScan(t *testing.T) {
	// Create a temporary folder and repository for testing
	tmpFolder := "test_folder"
	tmpFolderCreationErr := os.Mkdir(tmpFolder, 0755)
	assert.NoError(t, tmpFolderCreationErr, "Could not create temporary folder")
	defer os.RemoveAll(tmpFolder)
	testRepo := tmpFolder + "/test_repo/.git"
	testRepoCreationErr := os.MkdirAll(testRepo, 0755)
	assert.NoError(t, testRepoCreationErr, "Could not create test repository")

	// Test scan
	Scan(tmpFolder)
	dotFilePath := GetDotFilePath()
	_, filePathErr := os.Stat(dotFilePath)
	assert.False(t, os.IsNotExist(filePathErr))
	_ = os.Remove(dotFilePath)
}
