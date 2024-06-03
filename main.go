package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func readDirignore(root string) (map[string]struct{}, error) {
	dirignorePath := filepath.Join(root, ".dirignore")
	ignoredDirs := make(map[string]struct{})

	file, err := os.Open(dirignorePath)
	if err != nil {

		if os.IsNotExist(err) {
			return ignoredDirs, nil
		}
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dir := strings.TrimSpace(scanner.Text())
		if dir != "" {
			ignoredDirs[dir] = struct{}{}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return ignoredDirs, nil
}

func scanDirectory(root string, prefix string, ignoredDirs map[string]struct{}) (string, error) {
	fmt.Printf("Scanning directory: %s with prefix: %s\n", root, prefix)
	var result strings.Builder
	entries, err := os.ReadDir(root)
	if err != nil {
		fmt.Printf("Error reading directory %s: %v\n", root, err)
		return "", err
	}

	for i, entry := range entries {

		if _, ok := ignoredDirs[entry.Name()]; ok {
			fmt.Printf("Skipping ignored directory: %s\n", entry.Name())
			continue
		}

		connector := "├── "
		newPrefix := prefix + "│   "
		if i == len(entries)-1 {
			connector = "└── "
			newPrefix = prefix + "    "
		}
		result.WriteString(fmt.Sprintf("%s%s%s\n", prefix, connector, entry.Name()))
		fmt.Printf("Added entry: %s%s%s\n", prefix, connector, entry.Name())
		if entry.IsDir() {
			subDir, err := scanDirectory(filepath.Join(root, entry.Name()), newPrefix, ignoredDirs)
			if err != nil {
				return "", err
			}
			result.WriteString(subDir)
		}
	}
	return result.String(), nil
}

func generateMarkdown(dir string, structure string) string {
	fmt.Printf("Generating Markdown for directory: %s\n", dir)
	markdown := fmt.Sprintf("# Directory structure of %s\n\n```\n%s```\n", dir, structure)
	fmt.Println("Markdown generated successfully")
	return markdown
}

func writeToFile(filename string, content string) error {
	fmt.Printf("Writing to file: %s\n", filename)
	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", filename, err)
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		fmt.Printf("Error writing to file %s: %v\n", filename, err)
		return err
	}
	fmt.Printf("File %s written successfully\n", filename)
	return nil
}

func ensureMdExtension(filename string) string {
	if filepath.Ext(filename) != ".md" {
		return filename + ".md"
	}
	return filename
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <directory to scan> <output markdown file>")
		return
	}

	dir := os.Args[1]
	output := ensureMdExtension(os.Args[2])

	ignoredDirs, err := readDirignore(dir)
	if err != nil {
		fmt.Println("Error reading .dirignore file:", err)
		return
	}

	fmt.Printf("Scanning directory: %s\n", dir)
	fmt.Printf("Output file: %s\n", output)

	structure, err := scanDirectory(dir, "", ignoredDirs)
	if err != nil {
		fmt.Println("Error scanning directory:", err)
		return
	}

	fmt.Println("Directory structure:\n", structure)

	markdownContent := generateMarkdown(dir, structure)

	err = writeToFile(output, markdownContent)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
	} else {
		fmt.Println("Markdown file created successfully.")
	}
}
