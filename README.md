# DirScanner

![image](https://github.com/user-attachments/assets/a0b3abc6-160f-460b-bfbf-ee3cc88e00cc)


DirScanner is a CLI tool written in Go that scans a directory and generates a Markdown file with the directory's structure. It supports custom connector styles, exclusion of certain file types, and limiting directory traversal depth.

## Features

- **Generate Directory Structure**: Easily generate a Markdown file representing the structure of a directory.
- **Custom Connector Styles**: Customize the symbols used to represent the directory tree.
- **Exclude File Types**: Exclude specific file types or directories based on patterns.
- **Limit Directory Depth**: Restrict how deep the tool scans the directory tree.

## Installation

You can install DirScanner directly using:

```sh
go install github.com/aymaneallaoui/dirscanner@latest
```

## Usage

### Basic Usage

To scan a directory and generate a Markdown file:

```sh
dirscanner <directory to scan> <output markdown file>
```

### Exclude File Types or Directories

You can exclude specific file types or directories using the --exclude flag or using the `.dirignore` file.

```sh
dirscanner ./dir structure.md --exclude ".txt" --exclude "node_modules"

```

### Limit Directory Depth

To limit how deep the tool scans the directory, use the `--depth` flag:

```sh
dirscanner ./dir structure.md --depth 2
```

### Customize Connector Styles

You can customize the symbols used to draw the directory tree:

```sh
dirscanner ./dir structure.md --intermediate "+-- " --last "`-- " --prefix "    " --branch "|   "
```

## Contributing

Contributions are welcome! Please submit a pull request or open an issue to discuss any changes or improvements.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
