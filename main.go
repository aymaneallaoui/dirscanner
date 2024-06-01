package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func scanDirectory(root string, prefix string) (string, error) {

	// TODO - adding relative path support

	// TODO - adding .dirignore to ignore dirs

	// TODO - skipping git and node_modules dir by default

	fmt.Printf("Scanning directory: %s with prefix: %s\n", root, prefix)
	var result strings.Builder
	entries, err := os.ReadDir(root)
	if err != nil {
		fmt.Printf("Error reading directory %s: %v\n", root, err)
		return "", err
	}

	for i, entry := range entries {
		connector := "├── "
		newPrefix := prefix + "│   "
		if i == len(entries)-1 {
			connector = "└── "
			newPrefix = prefix + "    "
		}
		result.WriteString(fmt.Sprintf("%s%s%s\n", prefix, connector, entry.Name()))
		fmt.Printf("Added entry: %s%s%s\n", prefix, connector, entry.Name())
		if entry.IsDir() {
			subDir, err := scanDirectory(filepath.Join(root, entry.Name()), newPrefix)
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

func main() {

	dir := flag.String("dir", ".", "The directory to scan")
	output := flag.String("output", "structure.md", "The output markdown file")
	flag.Parse()

	fmt.Printf("Scanning directory: %s\n", *dir)
	fmt.Printf("Output file: %s\n", *output)

	structure, err := scanDirectory(*dir, "")
	if err != nil {
		fmt.Println("Error scanning directory:", err)
		return
	}

	fmt.Println("Direcotory structure:\n", structure)

	markdownContent := generateMarkdown(*dir, structure)

	err = writeToFile(*output, markdownContent)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
	} else {
		fmt.Println("Markdown file created successfully.")
	}
}
