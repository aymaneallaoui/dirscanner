// main_test.go
package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReadDirIgnore(t *testing.T) {
	// Set up a temporary directory for testing
	tmpDir := t.TempDir()
	ignoreFilePath := filepath.Join(tmpDir, ".dirignore")

	// Write some ignored directories to the .dirignore file
	ignoredDirs := []string{"ignored1", "ignored2"}
	err := os.WriteFile(ignoreFilePath, []byte(strings.Join(ignoredDirs, "\n")), 0644)
	if err != nil {
		t.Fatalf("Failed to write .dirignore file: %v", err)
	}

	// Test reading the .dirignore file
	ignoredMap, err := readDirIgnore(tmpDir)
	if err != nil {
		t.Fatalf("Error reading .dirignore: %v", err)
	}

	// Check that the ignored directories are correctly read
	for _, dir := range ignoredDirs {
		if _, ok := ignoredMap[dir]; !ok {
			t.Errorf("Expected directory %s to be ignored, but it was not", dir)
		}
	}
}

func TestScanDirectory(t *testing.T) {
	// Set up a temporary directory structure for testing
	tmpDir := t.TempDir()
	os.Mkdir(filepath.Join(tmpDir, "dir1"), 0755)
	os.Mkdir(filepath.Join(tmpDir, "dir2"), 0755)
	os.WriteFile(filepath.Join(tmpDir, "file1.txt"), []byte("content"), 0644)

	ignoredDirs := map[string]struct{}{
		"dir2": {},
	}

	// Test scanning the directory
	structure, err := scanDirectory(tmpDir, "", ignoredDirs)
	if err != nil {
		t.Fatalf("Error scanning directory: %v", err)
	}

	// Check that the structure is as expected
	expectedStructure := "├── dir1\n└── file1.txt\n"
	if structure != expectedStructure {
		t.Errorf("Expected structure:\n%s\nGot:\n%s", expectedStructure, structure)
	}
}

func TestEnsureMdExtension(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"output", "output.md"},
		{"output.md", "output.md"},
		{"doc.txt", "doc.txt.md"},
	}

	for _, tt := range tests {
		result := ensureMdExtension(tt.input)
		if result != tt.expected {
			t.Errorf("ensureMdExtension(%s): expected %s, got %s", tt.input, tt.expected, result)
		}
	}
}

func TestWriteToFile(t *testing.T) {
	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "testfile.md")
	content := "This is a test"

	err := writeToFile(filename, content)
	if err != nil {
		t.Fatalf("Error writing to file: %v", err)
	}

	readContent, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Error reading file: %v", err)
	}

	if string(readContent) != content {
		t.Errorf("Expected content: %s, got: %s", content, string(readContent))
	}
}
