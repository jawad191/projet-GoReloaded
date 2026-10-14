# Go-reloaded

A command-line text transformation project written in Go.

## Features
- `(up)`, `(low)`, `(cap)` and numbered variants such as `(up, 2)`
- Hexadecimal `(hex)` and binary `(bin)` conversion
- Punctuation and quote spacing fixes
- `a` to `an` before a word beginning with a vowel or `h`
- Preserves input line breaks
- Helpful file and argument errors

## Run
Open a terminal in this folder:

```powershell
go run . sample.txt result.txt
```

The program reads `sample.txt` and writes the transformed text to `result.txt`.

## Test
Run automated tests with:

```powershell
go test ./...
```

## Files
- `main.go`: application
- `main_test.go`: automated tests
- `tests/`: sample test inputs
- `GITHUB.md`: instructions for uploading to GitHub
