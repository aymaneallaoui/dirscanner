package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScanDirectory(t *testing.T) {

	testDir := t.TempDir()

	os.Mkdir(filepath.Join(testDir, "dir1"), 0755)
	os.Mkdir(filepath.Join(testDir, "dir1", "subdir1"), 0755)
	os.WriteFile(filepath.Join(testDir, "dir1", "file1.txt"), []byte("file1"), 0644)
	os.WriteFile(filepath.Join(testDir, "file2.txt"), []byte("file2"), 0644)

	expected := `├── dir1
│   ├── file1.txt
│   └── subdir1
└── file2.txt
`

	result, err := scanDirectory(testDir, "")
	if err != nil {
		t.Fatalf("Error scanning directory: %v", err)
	}

	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

func TestGenerateMarkdown(t *testing.T) {
	dir := "testdir"
	structure := "├── file1.txt\n└── subdir1\n"
	expected := "# Directory structure of testdir\n\n```\n├── file1.txt\n└── subdir1\n```\n"
	result := generateMarkdown(dir, structure)

	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

func TestWriteToFile(t *testing.T) {
	content := "test content"
	testFile := filepath.Join(t.TempDir(), "testfile.md")

	err := writeToFile(testFile, content)
	if err != nil {
		t.Fatalf("Error writing to file: %v", err)
	}

	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Error reading file: %v", err)
	}

	if string(data) != content {
		t.Errorf("Expected:\n%s\nGot:\n%s", content, string(data))
	}
}
