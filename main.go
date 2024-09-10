package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func readDirIgnore(root string) (map[string]struct{}, error) {
	ignoredDirs := make(map[string]struct{})
	dirIgnorePath := filepath.Join(root, ".dirignore")

	file, err := os.Open(dirIgnorePath)
	if err != nil {
		if os.IsNotExist(err) {
			return ignoredDirs, nil
		}
		return nil, fmt.Errorf("error opening .dirignore file: %v", err)
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
		return nil, fmt.Errorf("error reading .dirignore file: %v", err)
	}

	return ignoredDirs, nil
}

func scanDirectory(root string, prefix string, ignoredDirs map[string]struct{}) (string, error) {
	//logrus.Infof("Scanning directory: %s with prefix: %s", root, prefix)
	var result strings.Builder
	entries, err := os.ReadDir(root)
	if err != nil {
		return "", fmt.Errorf("error reading directory %s: %v", root, err)
	}

	for i, entry := range entries {
		if _, ok := ignoredDirs[entry.Name()]; ok {
			logrus.Infof("Skipping ignored directory: %s", entry.Name())
			continue
		}

		connector := "├── "
		newPrefix := prefix + "│   "
		if i == len(entries)-1 {
			connector = "└── "
			newPrefix = prefix + "    "
		}
		result.WriteString(fmt.Sprintf("%s%s%s\n", prefix, connector, entry.Name()))

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
	logrus.Infof("Generating Markdown for directory: %s", dir)
	return fmt.Sprintf("# Directory structure of %s\n\n```\n%s```\n", dir, structure)
}

func writeToFile(filename string, content string) error {
	logrus.Infof("Writing to file: %s", filename)
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file %s: %v", filename, err)
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return fmt.Errorf("error writing to file %s: %v", filename, err)
	}
	return nil
}

func ensureMdExtension(filename string) string {
	if filepath.Ext(filename) != ".md" {
		return filename + ".md"
	}
	return filename
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "dirscanner",
		Short: "A CLI tool to scan directories and generate a Markdown file with the structure",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return fmt.Errorf("usage: dirscanner <directory to scan> <output markdown file>")
			}

			dir := args[0]
			output := ensureMdExtension(args[1])

			ignoredDirs, err := readDirIgnore(dir)
			if err != nil {
				return fmt.Errorf("error reading .dirignore file: %v", err)
			}

			structure, err := scanDirectory(dir, "", ignoredDirs)
			if err != nil {
				return fmt.Errorf("error scanning directory: %v", err)
			}

			markdownContent := generateMarkdown(dir, structure)

			if err := writeToFile(output, markdownContent); err != nil {
				return fmt.Errorf("error writing to output file: %v", err)
			}

			fmt.Println("Markdown file created successfully.")
			return nil
		},
	}

	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
