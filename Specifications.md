# Project Spec: **Go Password Utils**

## 1. Project Summary

**Project name:** `gopwdutil`  
**Type:** Interactive command-line application to create and analyse a password  
**Goal:** A menu-driven Go CLI that lets a user input a password, run analyses and transformations on it, and export the results.

---

## 2. Target Outcome

By the end of the project, the CLI should let the user:

- enter a password (manually, from clipboard, or randomly generated),
- choose from analysis, transformation, and output submenus,
- see the result,
- and continue using the program until they choose to exit.

---

## 3. Functional Scope

### 3.1 Counts (`analysis` package)

#### 3.1.1 Character count
Returns the number of non-whitespace characters.

**Example:**  
Input: `hello world`  
Output: `Character count: 10 characters`

---

#### 3.1.2 Word count
Returns the number of words (split by whitespace, ignoring empty segments).

**Example:**  
Input: `Go is fun`  
Output: `Word count: 3 words`

---

#### 3.1.3 Letter count
Returns the number of alphabetic characters (upper + lower).

---

#### 3.1.4 Uppercase count
Returns the number of uppercase letters.

---

#### 3.1.5 Lowercase count
Returns the number of lowercase letters.

---

#### 3.1.6 Numeric count
Returns the number of digit characters.

---

#### 3.1.7 Special character count
Returns the number of non-alphanumeric characters from the defined set:

| Name | Character | Type |
|:---:|:---:|:---:|
| Space | ` ` | Whitespace |
| Period | `.` | Punctuation |
| Comma | `,` | Punctuation |
| Semicolon | `;` | Punctuation |
| Colon | `:` | Punctuation |
| Exclamation mark | `!` | Punctuation |
| Question mark | `?` | Punctuation |
| Apostrophe | `'` | Quotation |
| Double quote | `"` | Quotation |
| Opening parentheses | `(` | Brackets |
| Closing parentheses | `)` | Brackets |
| Opening square bracket | `[` | Brackets |
| Closing square bracket | `]` | Brackets |
| Opening curly brace | `{` | Brackets |
| Closing curly brace | `}` | Brackets |
| Opening angle bracket | `<` | Brackets |
| Closing angle bracket | `>` | Brackets |
| Hyphen/Minus | `-` | Slashes |
| Underscore | `_` | Slashes |
| Forward slash | `/` | Slashes |
| Backslash | `\` | Slashes |
| Vertical bar/Pipe | `\|` | Slashes |
| Plus | `+` | Math |
| Asterisk | `*` | Math |
| Equals | `=` | Math |
| Percent | `%` | Math |
| Caret | `^` | Math |
| At sign | `@` | Special |
| Hash | `#` | Special |
| Dollar sign | `$` | Special |
| Ampersand | `&` | Special |
| Tilde | `~` | Special |
| Backtick | `` ` `` | Special |

---

#### 3.1.8 Repeated character counts
Reports characters that appear more than once and how many times each appears.

---

### 3.2 Strength Analysis (`analysis` package)

A scoring-based system that outputs a label: `Weak`, `Fair`, `Strong`, or `Very Strong`.

**Scoring components:**

- Length score (exponential tiers):
  - `>= 8` → +2
  - `>= 14` → +4
  - `>= 18` → +8
  - `>= 36` → +16
  - `>= 72` → +32
- Uppercase count: +1 per character
- Lowercase count: +1 per character
- Numeric count: +1 per character
- Special character count: +1 per character
- Repeated character penalty (fewer repeats = higher score):
  - 0 repeated types → +8
  - 1 → +4
  - 2 → +2
  - 3 → +1
  - 4+ → +0

**Labels:**
- `score >= 16` → `Very Strong`
- `score >= 10` → `Strong`
- `score >= 4` → `Fair`
- default → `Weak`

---

### 3.3 Transformations (`transform` package)

#### 3.3.1 Reverse
Reverses the password in-place and displays the result.

**Example:**  
Input: `gopher`  
Output: `rehpog`

---

#### 3.3.2 Scramble
Randomly shuffles the characters of the password using a Fisher-Yates-style selection with `crypto/rand`.

---

#### 3.3.3 Pepper
Prompts the user to manually enter a pepper string (16–32 characters). Stored in `Password.Pepper`.

---

#### 3.3.4 Salt
Generates a cryptographically random salt (16–32 characters, all character sets). Stored in `Password.Salt`.

---

#### 3.3.5 Hash
Hashes the password using **Argon2id** (`golang.org/x/crypto/argon2`) with:
- Input: `password + pepper` as the key
- Input: `salt` as the salt
- Parameters: `t=3`, `m=62*1024 KB`, `p=1`, `keyLen=32`
- Output: base64-encoded key stored in `Password.HashedKey`

Warns if salt or pepper is missing before hashing.

---

### 3.4 Input Methods (`input` package)

#### 3.4.1 Manual input
Prompts the user to type a password directly. Length must be 8–72 characters. Whitespace-only input is rejected.

#### 3.4.2 Load from clipboard
Clears the clipboard, waits for the user to copy a value, then reads it. Same length and whitespace validation as manual input. Clears the clipboard after reading.

#### 3.4.3 Generate random
Prompts the user for a desired length (8–72), then generates a random password using `crypto/rand` from the full character set (lower + upper + numeric + special).

---

### 3.5 Output Methods (`output` package)

#### 3.5.1 To screen
Prints password, pepper, salt, and hashed key to stdout.

#### 3.5.2 To clipboard
Prompts the user to choose which value to copy (password, pepper, salt, or hashed key), then writes it to the system clipboard using `golang.design/x/clipboard`.

#### 3.5.3 To file
Writes all four values (password, pepper, salt, hashed key) to `output.txt` with `0600` permissions.

> **Note:** Exporting to an external file is part of this exercise only and should not be handled this way in a production scenario.

---

## 4. Non-Goals (v1)

- Advanced Unicode normalization
- Persistent history or database storage
- Colored terminal output
- Command-line flags parsing

---

## 5. User Experience Specification

### App flow

```text
======= CLI Password Utils =======

=========== Main Menu ============
Password length: 0
1) Input tools
0) Exit
----------------------------------
Choice:
```

Once a password is loaded (length > 0), the full menu appears:

```text
=========== Main Menu ============
Password length: 12
1) Input tools
2) Analysis tools
3) Transformation tools
4) Output options
5) Reset current password
0) Exit
----------------------------------
Choice:
```

Each option opens a submenu that loops until the user enters `0` to return to the Main Menu.

### Submenus

**Input tools:**
```text
========== Input Utils ===========
1) Manual input
2) Load from clipboard
3) Generate random string
0) Return to Main Menu
```

**Analysis tools:**
```text
======== Analysis Utils ==========
1) Character count
2) Word count
3) Letter count
4) Upper count
5) Lower count
6) Numeric count
7) Special count
8) Repeated counts
9) Strength analysis
0) Return to Main Menu
```

**Transformation tools:**
```text
====== Transformation Tools ======
1) Reverse password
2) Scramble password
3) Crack fresh pepper
4) Sprinkle some salt
5) Hash password
0) Return to Main Menu
```

**Output options:**
```text
========== Output Utils ==========
1) To screen
2) To clipboard
3) To file
0) Return to Main Menu
```

### Input behavior
- Invalid menu options show an error and re-prompt.
- Invalid or whitespace-only passwords are rejected with a descriptive error.
- Password length is validated against `MinPassword=8` and `MaxPassword=72`.
- Pepper length is validated against `MinPepper=16` and `MaxPepper=32`.
- Salt length is randomly chosen between `MinSalt=16` and `MaxSalt=32`.

---

## 6. Package Structure

```text
gopwdutil/
├── main.go
├── ui/
│   ├── main_menu.go
│   ├── input_menu.go
│   ├── analysis_menu.go
│   ├── transform_menu.go
│   └── output_menu.go
├── utils/
│   ├── input.go
│   ├── analysis.go
│   ├── transform.go
│   └── output.go
├── analysis/
│   ├── count.go
│   ├── repeated.go
│   └── strength.go
├── transform/
│   ├── reverse.go
│   ├── scramble.go
│   ├── pepper.go
│   ├── salt.go
│   └── hash.go
├── input/
│   ├── manual.go
│   ├── from_clipboard.go
│   └── random.go
├── output/
│   ├── to_screen.go
│   ├── to_clipboard.go
│   └── to_file.go
└── tools/
    ├── tools.go
    ├── get_choice.go
    ├── contains_byte.go
    └── reset.go
```

### Package responsibilities

| Package | Responsibility |
|---|---|
| `main` | Entry point; owns the main menu loop; calls `utils` |
| `ui` | Renders all menus; returns the user's choice |
| `utils` | Submenu loop coordinators; bridge between `ui` and logic packages |
| `analysis` | All counting logic and strength scoring |
| `transform` | Reverse, scramble, pepper, salt, and hash operations |
| `input` | Manual, clipboard, and random password input |
| `output` | Screen, clipboard, and file export |
| `tools` | Shared constants, types, character sets, errors, and utilities |

---

## 7. Function-Level Spec

### `tools`

```go
// Shared type
type Password struct {
    Password  []byte
    Salt      []byte
    Pepper    []byte
    HashedKey []byte
}

// Constants
MinPassword, MaxPassword = 8, 72
MinSalt, MaxSalt         = 16, 32
MinPepper, MaxPepper     = 16, 32

// Character sets
LowerChar, UpperChar, NumericChar, SpecialChar []byte

// Utilities
Reset(word *[]byte)                          // zeroes and shrinks slice
ContainsByte(b byte, ref []byte) bool        // membership check
GetChoice(maxOption int) (int, error)        // reads and validates an integer menu choice
```

---

### `analysis`

```go
Count(ppwd *[]byte, choice int) error        // dispatches to one of the count helpers by choice (1–7)
Repeated(ppwd *[]byte) error                 // prints characters appearing more than once
Strength(ppwd *[]byte) error                 // prints scoring breakdown and label
```

Internal helpers (unexported):
```go
countChar, countWord, countLetter, countUpper, countLower, countNumeric, countSpecial
getCount, getCountType, getCountUnit
getRepeated(ppwd *[]byte) (map[byte]int, error)
```

---

### `transform`

```go
Reverse(ppwd *[]byte) error
Scramble(ppwd *[]byte) error
Pepper(ppepper *[]byte) error
Salt(psalt *[]byte) error
Hash(ppwd *Password) error
```

---

### `input`

```go
Manual(ppwd *[]byte) error
FromClipboard(ppwd *[]byte) error
Random(ppwd *[]byte) error
```

---

### `output`

```go
ToScreen(ppwd *Password) error
ToClipboard(ppwd *Password) error
ToFile(ppwd *Password) error
```

---

### `ui`

```go
MainMenu(length int) (int, error)
InputMenu() (int, error)
AnalysisMenu() (int, error)
TransformMenu() (int, error)
OutputMenu() (int, error)
```

---

### `utils`

```go
Input(ppwd *[]byte) error
Analysis(ppwd *[]byte) error
Transform(ppwd *Password) error
Output(ppwd *Password) error
```

---

## 8. Security Notes

- Password, pepper, salt, and hashed key are stored as `[]byte` (not `string`) so they can be zeroed in memory.
- `tools.Reset()` zeros each byte before shrinking the slice, avoiding lingering sensitive data.
- `main.go` calls `Reset` on all four fields before exit.
- The "Reset current password" option in the main menu zeroes all four fields.
- `crypto/rand` is used for all random generation (salt, random password, scramble).
- Hashing uses Argon2id via `golang.org/x/crypto/argon2`.

---

## 9. External Dependencies

| Module | Usage |
|---|---|
| `golang.org/x/crypto/argon2` | Argon2id password hashing |
| `golang.design/x/clipboard` | Clipboard read/write |

---

## 10. Argon2 Parameter Reference

### Time cost (t)
- Number of times Argon2 iterates over its memory
- Increases CPU work; roughly linear impact on runtime
- This project uses: `t = 3`

### Memory cost (m)
- Amount of RAM (in KB) Argon2 uses
- Main defense against GPU/ASIC parallelism
- This project uses: `m = 62 * 1024 KB` (~62 MB)

### Parallelism (p)
- Number of threads used during hashing
- This project uses: `p = 1`

### Key length (l)
- Length of the resulting hash in bytes
- This project uses: `l = 32` (encoded as base64 string)

---

## 11. Acceptance Criteria

- The program compiles and runs from the terminal
- It starts by showing the main menu with password length
- Full menu options appear only after a password is loaded
- Each submenu loops until the user returns to the main menu
- Invalid options are handled without crashing
- All sensitive byte slices are zeroed before the program exits
- Logic is organized into the packages described above
- `main.go` only coordinates the top-level loop

---

## 12. Improvements (Backlog)

- Input → Random → Allow the user to choose which character sets to include
- Tools → Refine `SpecialChar` into subcategories (punctuation, math, brackets, etc.)
- Implement batch processing of multiple passwords
