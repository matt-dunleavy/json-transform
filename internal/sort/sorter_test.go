// internal/sort/sorter_test.go
package sort

import (
    "io/ioutil"
    "os"
    "path/filepath"
    "testing"
)

func TestSortJSONFile(t *testing.T) {
    tempDir := t.TempDir()
    inputFilePath := filepath.Join(tempDir, "input.json")
    outputFilePath := filepath.Join(tempDir, "output.json")

    inputJSON := `[
        {"name": "Alice", "age": 30},
        {"name": "Bob", "age": 25},
        {"name": "Charlie", "age": 35}
    ]`

    if err := ioutil.WriteFile(inputFilePath, []byte(inputJSON), 0644); err != nil {
        t.Fatalf("failed to write input file: %v", err)
    }

    field := "name"

    if err := SortJSONFile(inputFilePath, outputFilePath, field); err != nil {
        t.Fatalf("failed to sort JSON file: %v", err)
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
    "age": 25,
    "name": "Bob"
  },
  {
    "age": 35,
    "name": "Charlie"
  }
]`

    if string(result) != expected {
        t.Errorf("unexpected result: got %s, want %s", string(result), expected)
    }
}
