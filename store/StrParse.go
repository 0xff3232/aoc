package store

import (
    "os"
    "fmt"
    "strings"
    
)

// Data is a struct that holds a slice of strings, each representing a line from a file.
type Data struct {
	Lines []string
}

// parse reads a file and returns a Data struct containing each line of the file.
func StrParse(fileName string) (*Data, error) {
    // os.ReadFile reads the entire file specified by fileName.
    // If there's an error (like the file doesn't exist), it returns the error.
    buf, err := os.ReadFile(fileName)
    if err != nil {
        // Wraps the error with additional context and returns it.
        return nil, fmt.Errorf("read input file: %w", err)
    }

    // Create a new Data struct.
    d := &Data{}

    // Split the content of the file into lines.
    // The string(buf) converts the byte slice (buf) into a string.
    // strings.Split then divides this string into a slice of strings, split by "\n" (newline character).
    lines := strings.Split(string(buf), "\n")

    // Iterate over each line.
    for _, line := range lines {
        // Skip empty lines.
        if line == "" {
            continue
        }
        // Append non-empty lines to the Lines field of the Data struct.
        d.Lines = append(d.Lines, line)
    }

    // Return the pointer to the Data struct and nil (no error).
    return d, nil
}
