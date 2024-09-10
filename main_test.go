package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReadDirIgnore(t *testing.T) {
	tmpDir := t.TempDir()
	ignoreFilePath := filepath.Join(tmpDir, ".dirignore")

	ignoredDirs := []string{"ignored1", "ignored2"}
	err := os.WriteFile(ignoreFilePath, []byte(strings.Join(ignoredDirs, "\n")), 0644)
	if err != nil {
		t.Fatalf("Failed to write .dirignore file: %v", err)
	}

	ignoredMap, err := readDirIgnore(tmpDir)
	if err != nil {
		t.Fatalf("Error reading .dirignore: %v", err)
	}

	for _, dir := range ignoredDirs {
		if _, ok := ignoredMap[dir]; !ok {
			t.Errorf("Expected directory %s to be ignored, but it was not", dir)
		}
	}
}

func TestScanDirectory(t *testing.T) {
	tmpDir := t.TempDir()
	os.Mkdir(filepath.Join(tmpDir, "dir1"), 0755)
	os.Mkdir(filepath.Join(tmpDir, "dir2"), 0755)
	os.WriteFile(filepath.Join(tmpDir, "file1.txt"), []byte("content"), 0644)

	ignoredDirs := map[string]struct{}{
		"dir2": {},
	}

	style := ConnectorStyle{
		Intermediate: "├── ",
		Last:         "└── ",
		Prefix:       "    ",
		Branch:       "│   ",
	}

	structure, err := scanDirectory(tmpDir, "", ignoredDirs, style, []string{}, -1, 0)
	if err != nil {
		t.Fatalf("Error scanning directory: %v", err)
	}

	expectedStructure := "├── dir1\n└── file1.txt\n"
	if structure != expectedStructure {
		t.Errorf("Expected structure:\n%s\nGot:\n%s", expectedStructure, structure)
	}
}

func TestScanDirectoryWithExclusions(t *testing.T) {
	tmpDir := t.TempDir()
	os.Mkdir(filepath.Join(tmpDir, "dir1"), 0755)
	os.WriteFile(filepath.Join(tmpDir, "file1.txt"), []byte("content"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "file2.log"), []byte("content"), 0644)

	ignoredDirs := map[string]struct{}{}
	style := ConnectorStyle{
		Intermediate: "├── ",
		Last:         "└── ",
		Prefix:       "    ",
		Branch:       "│   ",
	}

	excludePatterns := []string{"*.log"}
	structure, err := scanDirectory(tmpDir, "", ignoredDirs, style, excludePatterns, -1, 0)
	if err != nil {
		t.Fatalf("Error scanning directory: %v", err)
	}

	expectedStructure := "├── dir1\n└── file1.txt\n"
	if structure != expectedStructure {
		t.Errorf("Expected structure:\n%s\nGot:\n%s", expectedStructure, structure)
	}
}

func TestScanDirectoryWithDepthLimit(t *testing.T) {
	tmpDir := t.TempDir()
	os.Mkdir(filepath.Join(tmpDir, "dir1"), 0755)
	os.Mkdir(filepath.Join(tmpDir, "dir1", "subdir1"), 0755)
	os.WriteFile(filepath.Join(tmpDir, "file1.txt"), []byte("content"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "dir1", "subdir1", "file2.txt"), []byte("content"), 0644)

	ignoredDirs := map[string]struct{}{}
	style := ConnectorStyle{
		Intermediate: "├── ",
		Last:         "└── ",
		Prefix:       "    ",
		Branch:       "│   ",
	}

	structure, err := scanDirectory(tmpDir, "", ignoredDirs, style, []string{}, 1, 0)
	if err != nil {
		t.Fatalf("Error scanning directory: %v", err)
	}

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
