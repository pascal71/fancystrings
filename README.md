# fancystrings

The `fancystrings` package provides utility functions for string manipulations in Go.

## Installation

To install `fancystrings`, use `go get`:

```bash
go get path/to/your/repo/fancystrings
```

Replace `path/to/your/repo` with the actual path to your repository.

## Usage

Import the package and use the `InsertSpaces` function to insert spaces after each character in a string.

```go
import "path/to/your/repo/fancystrings"

...

str := "Hello"
result := fancystrings.InsertSpaces(str)
fmt.Println(result)  // Outputs: H e l l o
```

## Functions

### InsertSpaces(s string) string

This function takes a string as input and returns a new string with spaces inserted after each character in the input string.

## Running Tests

Navigate to the `fancystrings` directory and run:

```bash
go test
```

This will execute the test cases provided in the `fancystrings_test.go` file.

