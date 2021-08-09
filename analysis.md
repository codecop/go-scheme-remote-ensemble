# Analysis

## Data Flow / Algorithm / Steps

- File name
- read file -> string
- string -> scan -> list of tokens (type/value)
  - move input into tokenizer
  - types of tokens: open brackets, close brackets, names (function, assignment), numbers, strings/quotes, #t, #f, symbols
- list of tokens -> tree
  - build up a tree
  - Content: tokens
  - approach: new children if reach open bracket
  - check for enough brackets
- language semantic checks, type checks
  - analyse syntax
  - number of arguments, string function need to be string
- Tree -> parsing/error handling and interpreting
- -> execution -> result
- result -> print

## IO Analysis

- inputs
  - file path/name/extension
  - file
    - what do expect in file
    - what to do about whitespace, delimiters, new line, comments (;)
    - encoding of file (utf8)
- output
  - text of all evaluated expressions in the file
  - one line per expression
  - if error, see line number, description

## Potential components

- file handling
- tokenizing
- parsing
- semantic checks (checks for all levels (e.g., tokens, tree))
- error collector
  - line numbers
- execution
- display
- main component

### Discussion

- checks in component or in each component
- separate into packages

- 2 more scanners: brackets, arbitrary names (UpperCase | LowerCase + Any)
