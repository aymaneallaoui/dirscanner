# DIRSCANNER:

ADD IMAGE HERE

## Description:

many times i need to share our file structure so I made this cli tool to make it easier in golang to scan provided directory and make a markdown file of the file structure of the whole directory.

## Usage:

- `dirscanner -dir=<your target dir> -output=<output-file.md>`

### example of the output:

```
├── go.mod
├── main.go
├── main_test.go
├── structure
└── testdir
    ├── dir1
    │   ├── file1.txt
    │   └── subdir1
    │       └── subdir2
    │           ├── file.txt
    │           └── lol
    └── file2.txt
```

## Installation:

- `go get github.com/aymaneallaoui/dirscanner`

## License:

MIT
