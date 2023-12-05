package store

import (
    "os"
    "fmt"
    "strings"
    
)

type Data struct {
	Lines []string // lines from input file stored in string array
}

/*
StrParse will take input file and parse string by \n into string array
*/
func StrParse(fileName string) (*Data, error) {
    
    buf, err := os.ReadFile(fileName)
    if err != nil {
        return nil, fmt.Errorf("read input file: %w", err)
    }

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

    return d, nil
}
