// internal/filter/filter_test.go
package filter

import (
    "io/ioutil"
    "os"
    "path/filepath"
    "testing"
)

func TestFilterJSONFile(t *testing.T) {
    tempDir := t.TempDir()
    inputFilePath := filepath.Join(tempDir, "input.json")
    outputFilePath := filepath.Join(tempDir, "output.json")

    inputJSON := `[
        {"name": "Alice", "age": 30},
        {"name": "Bob", "age": 25},
        {"name": "Charlie", "age": 30}
    ]`

    if err := ioutil.WriteFile(inputFilePath, []byte(inputJSON), 0644); err != nil {
        t.Fatalf("failed to write input file: %v", err)
    }

    criteria := "30"

    if err := FilterJSONFile(inputFilePath, outputFilePath, criteria); err != nil {
        t.Fatalf("failed to filter JSON file: %v", err)
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
    "age": 30,
    "name": "Charlie"
  }
]`

    if string(result) != expected {
        t.Errorf("unexpected result: got %s, want %s", string(result), expected)
    }
}
