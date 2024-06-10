// internal/search/searcher_test.go
package search

import (
    "io/ioutil"
    "os"
    "path/filepath"
    "testing"
)

func TestSearchJSONFile(t *testing.T) {
    tempDir := t.TempDir()
    inputFilePath := filepath.Join(tempDir, "input.json")
    outputFilePath := filepath.Join(tempDir, "output.json")

    inputJSON := `[
        {"name": "Alice", "age": 30},
        {"name": "Bob", "age": 25},
        {"name": "Charlie", "age": 35},
        {"name": "Alice Smith", "age": 40}
    ]`

    if err := ioutil.WriteFile(inputFilePath, []byte(inputJSON), 0644); err != nil {
        t.Fatalf("failed to write input file: %v", err)
    }

    pattern := "alice"

    if err := SearchJSONFile(inputFilePath, outputFilePath, pattern); err != nil {
        t.Fatalf("failed to search JSON file: %v", err)
    }

    result, err := ioutil.ReadFile(outputFilePath)
    if err != nil {
        t.Fatalf("failed to read output file: %v", err)
    }

    expected := `[
  {
    "age": 30,
    "name": "Alice"
  },
  {
    "age": 40,
    "name": "Alice Smith"
  }
]`

    if string(result) != expected {
        t.Errorf("unexpected result: got %s, want %s", string(result), expected)
    }
}
