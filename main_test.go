package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReadDirignore(t *testing.T) {
	dir := t.TempDir()
	ignoreFile := filepath.Join(dir, ".dirignore")
	err := ioutil.WriteFile(ignoreFile, []byte("dir1\ndir2\n"), 0644)
	if err != nil {
		t.Fatalf("Failed to write .dirignore file: %v", err)
	}

	ignoredDirs, err := readDirignore(dir)
	if err != nil {
		t.Fatalf("readDirignore returned an error: %v", err)
	}

	if _, exists := ignoredDirs["dir1"]; !exists {
		t.Errorf("Expected 'dir1' to be in ignoredDirs")
	}
	if _, exists := ignoredDirs["dir2"]; !exists {
		t.Errorf("Expected 'dir2' to be in ignoredDirs")
	}
}

func TestScanDirectory(t *testing.T) {
	dir := t.TempDir()
	subdir := filepath.Join(dir, "subdir")
	err := os.Mkdir(subdir, 0755)
	if err != nil {
		t.Fatalf("Failed to create subdir: %v", err)
	}
	file := filepath.Join(subdir, "file.txt")
	err = ioutil.WriteFile(file, []byte("content"), 0644)
	if err != nil {
		t.Fatalf("Failed to write file.txt: %v", err)
	}

	ignoredDirs := map[string]struct{}{
		"ignoredir": {},
	}

	result, err := scanDirectory(dir, "", ignoredDirs)
	if err != nil {
		t.Fatalf("scanDirectory returned an error: %v", err)
	}

	if !strings.Contains(result, "subdir") {
		t.Errorf("Expected 'subdir' to be in the scan result")
	}
	if strings.Contains(result, "ignoredir") {
		t.Errorf("Did not expect 'ignoredir' to be in the scan result")
	}
}

func TestGenerateMarkdown(t *testing.T) {
	dir := "/path/to/dir"
	structure := "├── file1\n└── file2\n"
	expected := "# Directory structure of /path/to/dir\n\n```\n├── file1\n└── file2\n```\n"
	result := generateMarkdown(dir, structure)

	if result != expected {
		t.Errorf("Expected %q, but got %q", expected, result)
	}
}

func TestWriteToFile(t *testing.T) {
	content := "test content"
	file := filepath.Join(t.TempDir(), "test.md")
	err := writeToFile(file, content)
	if err != nil {
		t.Fatalf("writeToFile returned an error: %v", err)
	}

	readContent, err := ioutil.ReadFile(file)
	if err != nil {
		t.Fatalf("Failed to read the file: %v", err)
	}

	if string(readContent) != content {
		t.Errorf("Expected file content to be %q, but got %q", content, string(readContent))
	}
}

func TestEnsureMdExtension(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"file", "file.md"},
		{"file.md", "file.md"},
		{"file.txt", "file.txt.md"},
	}

	for _, test := range tests {
		result := ensureMdExtension(test.input)
		if result != test.expected {
			t.Errorf("Expected %q, but got %q", test.expected, result)
		}
	}
}
