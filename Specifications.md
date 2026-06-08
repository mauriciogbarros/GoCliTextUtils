# Project Spec: **CLI Text Utilities**

## 1. Project Summary

**Project name:** `clitextutils`  
**Type:** Interactive command-line application  
**Goal:** Build a small but well-structured Go CLI that lets a user enter text and run simple analyses/transforms on it.

---

## 2. Target Outcome

By the end of the project, the CLI should let the user:

- enter a piece of text,
- choose one utility from a menu,
- see the result,
- and continue using the program until they choose to exit.

---

## 3. Functional Scope
### 3.1 Counts
#### 3.1.1 Character count
Return the total number of characters in the input text.

**Example:**  
Input: `hello world`  
Output: `11 characters`

---

#### 3.1.2 Word count
Return the number of words in the input text.

**Example:**  
Input: `Go is fun`  
Output: `3 words`

---

#### 3.1.3 Vowel count
Count how many vowels appear in the text.

**Example:**  
Input: `Go is fun`  
Output: `3 vowels`

---

#### 3.1.4 Consonant count
Count how many consonants appear in the text.

**Example:**
Input: `Go is fun`
Output: `4 consonants`

---

#### 3.1.5 Number count
Count how many numbers appear in the text.

**Example:**
Input: `Go is 1-2-3 fun`
Output: `3 numbers`

---

#### 3.1.6 Special character count
Count how many non-alphanumeric characters appear in the text.
|Name|Character|Type|ASCII|
|:---:|:---:|:---:|:---:|
|Space|` `|Whitespace||
|Period/Full stop|`.`|Punctuation||
|Comma|`,`|Punctuation||
|Semicolon|`;`|Punctuation||
|Colon|`:`|Punctuation||
|Exclamation mark|`!`|Punctuation||
|Question mark|`?`|Punctuation||
|Apostrophe|`'`|Quotation||
|Double quote|`"`|Quotation||
|Opening parentheses|`(`|Brackets||
|Closing parentheses|`)`|Brackets||
|Opening square bracket|`[`|Brackets||
|Closing square bracket|`]`|Brackets||
|Opening curly brace|`{`|Brackets||
|Closing curly brace|`}`|Brackets||
|Opening angle bracket|`<`|Brackets||
|Closing angle bracket|`>`|Brackets||
|Hyphen/Minus|`-`|Slashes||
|Underscore|`_`|Slashes||
|Forward slash|`/`|Slashes||
|Backslash|`\`|Slashes||
|Vertical bar/Pipe|`|`|Slashes||
|Plus|`+`|Math||
|Asterisk|`*`|Math||
|Equals|`=`|Math||
|Percent|`%`|Math||
|Caret|`^`|Math||
|At sign|`@`|Special||
|Hash / number sign|`#`|Special||
|Dollar sign|`$`|Special||
|Ampersand|`&`|Special||
|Tilde|`~`|Special||
|Backtick / grave accent|`\``|Special||

**Example:**
Input: `Go is @1-2-3 fun`
Output: `6 special characters`

---

### 3.2 Case detection
Report:
- whether the text is all uppercase,
- all lowercase,
- mixed case,

**Example:**  
Input: `HELLO` → `all uppercase`  
Input: `hello` → `all lowercase`  
Input: `Hello123` → `mixed case`

---

### 3.3 Transformations
https://www.geeksforgeeks.org/computer-networks/encryption-vs-hashing-vs-salting/

#### 3.3.1 Reverse text
Return the input text reversed.

**Example:**  
Input: `gopher`  
Output: `rehpog`

---

#### 3.3.1 Add salt


---

#### 3.3.2 Hash it


---

#### 3.3.3 Encode it


---

### 3.4. Password strength checker
Classify a string as:
- `weak`
- `medium`
- `strong`

Suggested rules:
- **weak**: too short or only one character type
- **medium**: sufficient length with some variety
- **strong**: longer text with uppercase, lowercase, and digits

This is a very good fit because it combines **loops**, **boolean flags**, and **conditionals**.

---

## Nice-to-have features (optional, not required)
These are good only if you finish early:

- digit count
- whitespace count
- palindrome check
- title-case conversion
- repeated-character detector

I’d keep these out of the first version so the project stays aligned with Chapters 1–4 rather than growing too wide too early.

---

## 4) Non-Goals

To keep the scope disciplined, I suggest **not** including these in v1:

- file input/output
- command-line flags parsing
- persistent history
- colored terminal output
- external libraries
- advanced Unicode normalization
- heavy use of arrays/slices/maps as central architecture

Those are all valid additions later, but they are not necessary to consolidate the first four chapters.

---

## 5) User Experience Specification

## App flow

When the program starts, it should show:

```text
=== Text Utilities ===
1. Character count
2. Word count
3. Vowel count
4. Detect case
5. Reverse text
6. Check password strength
7. Enter new text
8. Exit
```

### Expected interaction model
1. Prompt the user to enter text when the app starts.
2. Show the menu.
3. Ask for an option.
4. Execute the selected utility.
5. Print the result.
6. Return to the menu.
7. Exit only when the user chooses `8`.

---

## Input behavior
- If the user enters an invalid menu option, show an error and re-prompt.
- If the input text is empty, some features should still work gracefully:
  - character count → `0`
  - word count → `0`
  - reverse → empty output
  - password strength → likely `weak`

---

## 6) Learning Objectives Mapped to Chapters

## Chapter 1: Syntax Basics
This project should exercise:
- variable declarations
- strings and booleans
- printing output
- reading input
- arithmetic comparisons
- package imports such as `fmt` and possibly `strings` if you decide to use standard helpers minimally. The chapter coverage for syntax basics is explicitly listed in the book materials. citeturn1search6turn1search16

## Chapter 2: Conditionals and Loops
This project should use:
- `if` / `else if` / `else`
- menu-selection branching
- loops for repeated app usage
- character-by-character inspection for vowels, digits, and case logic

Conditionals and loops are explicitly the focus of Chapter 2. citeturn1search6

## Chapter 3: Functions
Each utility should be implemented as its own function.
Examples:
- `CountCharacters`
- `CountWords`
- `CountVowels`
- `DetectCase`
- `ReverseText`
- `CheckPasswordStrength`

Functions are explicitly the focus of Chapter 3, including return values. citeturn1search6

## Chapter 4: Packages
The project should be split into multiple packages instead of one oversized `main.go`. Packages are explicitly the focus of Chapter 4. citeturn1search6

---

## 7) Proposed Package Structure

I suggest this layout:

```text
textutils/
├── main.go
├── ui/
│   └── menu.go
├── analysis/
│   ├── count.go
│   ├── casecheck.go
│   └── password.go
├── transform/
│   └── reverse.go
└── input/
    └── prompt.go
```

## Package responsibilities

### `main`
- application entry point
- orchestrates the menu loop
- calls functions from other packages

### `ui`
- prints the menu
- prints section headers or helper prompts

### `input`
- reads user input
- trims trailing newline if needed
- validates menu selections

### `analysis`
- character count
- word count
- vowel count
- case detection
- password strength logic

### `transform`
- reverse text

This structure keeps your code modular without introducing concepts that are too advanced for this stage.

---

## 8) Function-Level Spec

Here is a **clear functional contract** for each part.

## In `input`

```go
ReadText() string
```
Reads a full line of text from the user and returns it.

```go
ReadOption() int
```
Reads a menu choice and returns it as an integer.

---

## In `ui`

```go
PrintWelcome()
PrintMenu()
PrintResult(label string, result string)
PrintError(message string)
```

These are intentionally simple and good practice for keeping presentation separate from logic.

---

## In `analysis`

```go
CountCharacters(text string) int
```
Returns total character count.

```go
CountWords(text string) int
```
Returns number of words.

```go
CountVowels(text string) int
```
Returns total vowels (`a, e, i, o, u`, case-insensitive).

```go
DetectCase(text string) string
```
Returns one of:
- `"all uppercase"`
- `"all lowercase"`
- `"mixed case"`
- `"no letters"`

```go
CheckPasswordStrength(text string) string
```
Returns one of:
- `"weak"`
- `"medium"`
- `"strong"`

---

## In `transform`

```go
ReverseText(text string) string
```
Returns the reversed version of the text.

---

## 9) Logic Rules

## Word count rule
Suggested rule:
- split text by spaces after trimming,
- ignore empty segments created by repeated spaces.

If you prefer to keep it simpler in v1, you can define a “word” as any chunk separated by spaces.

## Vowel count rule
Count:
- `a`
- `e`
- `i`
- `o`
- `u`
and uppercase versions too.

## Case detection rule
- if it contains letters and all letters are uppercase → `all uppercase`
- if it contains letters and all letters are lowercase → `all lowercase`
- if it contains a mix → `mixed case`
- if no alphabetic characters are present → `no letters`

## Password strength rule
A simple and chapter-appropriate rule set:

### Weak
- length < 8  
**or**
- only one type among lowercase / uppercase / digits

### Medium
- length >= 8
- at least two of: lowercase / uppercase / digits

### Strong
- length >= 10
- contains lowercase, uppercase, and digits

This rule set is simple enough to implement with boolean flags and conditionals.

---

## 10) Example Session

```text
=== Text Utilities ===
Enter your text:
Hello123

1. Character count
2. Word count
3. Vowel count
4. Detect case
5. Reverse text
6. Check password strength
7. Enter new text
8. Exit

Choose an option: 1
Result: 8 characters

Choose an option: 3
Result: 2 vowels

Choose an option: 4
Result: mixed case

Choose an option: 5
Result: 321olleH

Choose an option: 6
Result: strong
```

---

## 11) Acceptance Criteria

Your v1 is “done” if all of these are true:

- the program compiles and runs from the terminal
- it starts by asking for text
- it shows a looped menu
- each menu option works
- invalid options are handled cleanly
- logic is split into multiple functions
- logic is organized into packages
- `main.go` mostly coordinates instead of containing all logic
- there is no duplicated logic for operations like printing menus or validating options

---

## 12) Suggested Build Plan

## Phase 1 — Skeleton
Build:
- `main.go`
- basic text input
- menu loop
- exit option

## Phase 2 — First utilities
Add:
- character count
- reverse text
- vowel count

These give fast progress and confidence.

## Phase 3 — Decision-heavy utilities
Add:
- case detection
- password strength checker

These deepen the Chapter 2 practice.

## Phase 4 — Refactor
Move code into:
- `analysis`
- `transform`
- `input`
- `ui`

This is where you really consolidate Chapter 4.

---

## 13) Stretch Goal for a “Chapter 4 polished” version

If you want one optional improvement that still feels appropriate, add:

### “Analyze all”
A menu option that prints:
- character count
- word count
- vowel count
- case classification
- password strength

for the current text all at once.

That is still compatible with Chapters 1–4 and makes the tool feel much more complete.

---

## 14) My recommendation on implementation order

If you want the smoothest path, I’d build in this exact order:

1. enter text  
2. menu loop  
3. character count  
4. reverse text  
5. vowel count  
6. word count  
7. case detection  
8. password checker  
9. package refactor

That ordering lets you ship a working version very early, then improve structure afterward.

---

# Final Spec Summary

**CLI Text Utilities** should be a menu-driven Go command-line app that stores one active text input and lets the user run a set of simple analysis and transformation operations on it. It should be deliberately scoped to practice **syntax basics**, **conditionals/loops**, **functions**, and **packages**, which is exactly what Chapters 1–4 cover. citeturn1search6turn1search16turn1search46
