# Password Utils

A small interactive Go command-line application that lets you enter a string representing a password and run analyses and transformations on it.

This is a portfolio project and a study on password security. The values of password, salt, pepper, and hash are outputs of this program. On a production scenario, the password would not be saved as plain text, and the pepper would be stored outside of the database. Additionally, exporting to an external final is only part of this exercise and it should not be handled like in a real-world scenario.

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

## Improvements
- Input => Random => Allow to choose which character sets the user wants to use.

- tools => Refine SpecialChar into more categories.

## Argon2
### Time cost (t)
- The number of times Argon2 runs over its memory
- What it does:
  - Increases CPU work
  - Makes each password guess slower
  - Roughly linear impact on runtime
- Typical values: 2 - 4, 3 is a default in many system

### Memory cost (m)
- Amount of RAM (in KB) Argon2 uses
- What it does:
  - Main defense against attackers
  - Forces attackers to allocate memory per guess
  - Limits GPU/ASIC parallelism
- Typical values:
  - Minimum secure: ~19 MB (OWASP baseline)
  - Typical login: 32 - 64 MB
  - Strong security: 64 - 128 MB
  - High security/offline: 128 - 256 MB+

### Parallelism (p)
- Number of threads (lanes) used during hashing
- What it does:
  - Utilizes multiple CPU cores
  - Improves performance (speed) more than security
  - Also affects internal memory mixing
- Typical values:
  - 1 - 4 threads
  - Many production setups use p = 1 or match CPU cores available

### Key length (l)
- Length of the resulting hash (in bytes)