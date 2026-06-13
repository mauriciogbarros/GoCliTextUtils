# Password Utils

A small interactive Go command-line application that lets you enter a string representing a password and run simple analyses and transformations on it.

## Table of Contents
- [Features](#features)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Running](#running)
- [Usage](#usage)
- [Notes](#notes)
- [Contributing](#contributing)
- [License](#license)

## Features
- Password length
- Character count (no whitespace)
- Word count
- Vowel count
- Consonant count
- Number count
- Special character count
- Case detection
- Reverse text
- Password strength check
- Password salt
- Password hash

## Project Structure
```text
gopwdutil/
├── main.go
├── ui/
│   └── menu.go
├── analysis/
│   ├── count.go
│   ├── casecheck.go
│   └── password.go
├── transform/
│   ├── salt.go
│   ├── hash.go
│   └── reverse.go
└── input/
    └── prompt.go
```

## Installation
### Prerequisites

- Go installed on your machine
- A terminal such as PowerShell, Command Prompt, Bash, or zsh

### Initialize the Project

If you are starting from scratch, create the folder and initialize a Go module:

```bash
mkdir 
cd gopwdutil
go mod init gopwdutil
```

Then add the files and folders from the project spec.

### Build Instructions

#### MacOS/Linux

From the project root, run:

```bash
go build -o textutils
```

This builds an executable named `textutils` in the current folder.

#### Windows

```powershell
go build -o textutils.exe
```

## Running

### Run directly without building a named executable

```bash
go run .
```

### Run the built executable
#### macOS / Linux

```bash
./textutils
```

#### Windows

```powershell
.\textutils.exe
```

## Usage
When you run the application, it will prompt you to enter a piece of text. After you input the text, it will display a menu of options for analyzing or transforming the text. You can select an option by entering the corresponding number. The application will then display the result of the selected operation. You can continue to select different options or exit the application when you're done.

### Workflow

1. Start the app.
2. Enter a piece of text when prompted.
3. Pick a menu option.
4. Review the result.
5. Choose another option or exit.

## Notes

- This README assumes the application entry point is `main.go` in the project root.
- If you choose a different module name, replace `textutils` in the `go mod init` command.
- If you later add tests, you can run them with:

```bash
go test ./...
```

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements please open an issue or submit a pull request.

1. Fork the repository.
2. Create a new branch for your changes.
3. Make your changes and commit them.
4. Push your changes to your forked repository.
5. Open a pull request.

## License

This project is licensed under the [MIT License](./LICENSE.md)


# Review before finishing

<!-- TODO  File naming: Go convention is snake_case or all lowercase for filenames (e.g., wordcount.go, not WordCount.go). PascalCase filenames are unusual and can confuse tooling.-->

<!-- TODO Return -1 for errors: Go's idiomatic error handling is returning (int, error) tuples, not sentinel values like -1. For example: -->

<!-- TODO Mixing concerns: functions like WordCount both compute a result and call fmt.Printf. In Go it's idiomatic to separate computation from I/O. Return the value and let the caller (menu/UI layer) handle printing. => Address after basic implementation of each count function, allow to pass the value to Strength Analysis-->
