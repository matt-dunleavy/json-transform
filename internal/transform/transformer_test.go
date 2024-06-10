// internal/transform/transformer_test.go
package transform

import (
    "io/ioutil"
    "os"
    "path/filepath"
    "testing"
)

func TestTransformJSONFile(t *testing.T) {
    tempDir := t.TempDir()
    inputFilePath := filepath.Join(tempDir, "input.json")
    outputFilePath := filepath.Join(tempDir, "output.json")
    rulesFilePath := filepath.Join(tempDir, "rules.json")

    inputJSON := `[
        {"name": "Alice", "age": 30},
        {"name": "Bob", "age": 25}
    ]`

    rulesJSON := `{
        "rename": {
            "old": "name",
            "new": "full_name"
        }
    }`

    if err := ioutil.WriteFile(inputFilePath, []byte(inputJSON), 0644); err != nil {
        t.Fatalf("failed to write input file: %v", err)
    }

    if err := ioutil.WriteFile(rulesFilePath, []byte(rulesJSON), 0644); err != nil {
        t.Fatalf("failed to write rules file: %v", err)
    }

    if err := TransformJSONFile(inputFilePath, outputFilePath, rulesFilePath); err != nil {
        t.Fatalf("failed to transform JSON file: %v", err)
    }

    result, err := ioutil.ReadFile(outputFilePath)
    if err != nil {
        t.Fatalf("failed to read output file: %v", err)
    }

    expected := `[
  {
    "age": 30,
    "full_name": "Alice"
  },
  {
    "age": 25,
    "full_name": "Bob"
  }
]`

    if string(result) != expected {
        t.Errorf("unexpected result: got %s, want %s", string(result), expected)
    }
}
