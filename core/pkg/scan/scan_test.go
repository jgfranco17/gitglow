package scan

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGetDotFilePath(t *testing.T) {
	dotFilePath := GetDotFilePath()
	if dotFilePath == "" {
		t.Error("Expected a non-empty dot file path, got an empty string.")
	}
}

func TestOpenFile(t *testing.T) {
	filePath := "testfile.txt"

	// Cleanup if the file exists
	defer func() {
		_ = os.Remove(filePath)
	}()

	// Create a test file
	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	file.Close()

	// Open the file
	f := openFile(filePath)
	defer f.Close()

	if f == nil {
		t.Error("Expected a non-nil file object, got nil.")
	}
}

func TestParseFileLinesToSlice(t *testing.T) {
	filePath := "testfile.txt"

	// Cleanup if the file exists
	defer func() {
		_ = os.Remove(filePath)
	}()

	// Create a test file
	content := "line1\nline2\nline3"
	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}

	lines := ParseFileLinesToSlice(filePath)

	expectedLines := []string{"line1", "line2", "line3"}

	if len(lines) != len(expectedLines) {
		t.Errorf("Expected %d lines, got %d", len(expectedLines), len(lines))
	}

	for i, line := range lines {
		if line != expectedLines[i] {
			t.Errorf("Expected line '%s', got '%s'", expectedLines[i], line)
		}
	}
}

// Add more tests for other functions as needed
// ...

func TestScan(t *testing.T) {
	// Create a temporary folder for testing
	tmpFolder := "test_folder"
	err := os.Mkdir(tmpFolder, 0755)
	if err != nil {
		t.Fatalf("Error creating temporary folder: %v", err)
	}
	defer os.RemoveAll(tmpFolder)

	// Create a test git repository
	testRepo := tmpFolder + "/test_repo/.git"
	err = os.MkdirAll(testRepo, 0755)
	if err != nil {
		t.Fatalf("Error creating test repository: %v", err)
	}

	// Run the scan
	Scan(tmpFolder)

	// Check if the dot file was created
	dotFilePath := GetDotFilePath()
	if _, err := os.Stat(dotFilePath); os.IsNotExist(err) {
		t.Errorf("Expected dot file to be created, but it was not.")
	}

	// Clean up
	_ = os.Remove(dotFilePath)
}
