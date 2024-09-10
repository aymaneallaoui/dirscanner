package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type ConnectorStyle struct {
	Intermediate string
	Last         string
	Prefix       string
	Branch       string
}

func patternToRegex(pattern string) (string, error) {

	regexPattern := regexp.QuoteMeta(pattern)

	regexPattern = strings.ReplaceAll(regexPattern, `\*`, ".*")
	regexPattern = strings.ReplaceAll(regexPattern, `\?`, ".")

	regexPattern = "^" + regexPattern + "$"

	_, err := regexp.Compile(regexPattern)
	if err != nil {
		return "", fmt.Errorf("error compiling regex pattern: %v", err)
	}

	return regexPattern, nil
}

func scanDirectory(root string, prefix string, ignoredDirs map[string]struct{}, style ConnectorStyle, excludePatterns []string, maxDepth, currentDepth int) (string, error) {
	logrus.Infof("Scanning directory: %s with prefix: %s", root, prefix)
	var result strings.Builder
	entries, err := os.ReadDir(root)
	if err != nil {
		return "", fmt.Errorf("error reading directory %s: %v", root, err)
	}

	
	filteredEntries := []os.DirEntry{}
	for _, entry := range entries {
		if _, ok := ignoredDirs[entry.Name()]; ok {
			logrus.Infof("Skipping ignored directory: %s", entry.Name())
			continue
		}

		
		excluded := false
		for _, pattern := range excludePatterns {
			regexPattern, err := patternToRegex(pattern)
			if err != nil {
				return "", fmt.Errorf("error matching pattern %s: %v", pattern, err)
			}
			matched, err := regexp.MatchString(regexPattern, entry.Name())
			if err != nil {
				return "", fmt.Errorf("error matching pattern %s: %v", regexPattern, err)
			}
			if matched {
				logrus.Infof("Skipping excluded file/directory: %s", entry.Name())
				excluded = true
				break
			}
		}
		if !excluded {
			filteredEntries = append(filteredEntries, entry)
		}
	}

	
	if maxDepth != -1 && currentDepth >= maxDepth {
		return "", nil
	}

	for i, entry := range filteredEntries {
		connector := style.Intermediate
		newPrefix := prefix + style.Branch
		if i == len(filteredEntries)-1 {
			connector = style.Last
			newPrefix = prefix + style.Prefix
		}
		result.WriteString(fmt.Sprintf("%s%s%s\n", prefix, connector, entry.Name()))

		if entry.IsDir() {
			subDir, err := scanDirectory(filepath.Join(root, entry.Name()), newPrefix, ignoredDirs, style, excludePatterns, maxDepth, currentDepth+1)
			if err != nil {
				return "", err
			}
			result.WriteString(subDir)
		}
	}
	return result.String(), nil
}

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
	var (
		intermediate string
		last         string
		prefix       string
		branch       string
		exclude      []string
		maxDepth     int
	)

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

			style := ConnectorStyle{
				Intermediate: intermediate,
				Last:         last,
				Prefix:       prefix,
				Branch:       branch,
			}

			structure, err := scanDirectory(dir, "", ignoredDirs, style, exclude, maxDepth, 0)
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

	rootCmd.Flags().StringVar(&intermediate, "intermediate", "├── ", "Symbol for intermediate nodes")
	rootCmd.Flags().StringVar(&last, "last", "└── ", "Symbol for the last node in a directory")
	rootCmd.Flags().StringVar(&prefix, "prefix", "    ", "Prefix for child nodes")
	rootCmd.Flags().StringVar(&branch, "branch", "│   ", "Branch for intermediate nodes")
	rootCmd.Flags().StringSliceVar(&exclude, "exclude", []string{}, "Exclude files or directories matching these patterns (e.g., '*.txt')")
	rootCmd.Flags().IntVar(&maxDepth, "depth", -1, "Limit the depth of the directory traversal (-1 for no limit)")

	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
