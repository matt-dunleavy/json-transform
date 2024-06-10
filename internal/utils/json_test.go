// internal/utils/json_test.go
package utils

import (
    "io/ioutil"
    "os"
    "path/filepath"
    "testing"
)

func TestReadJSONFile(t *testing.T) {
    tempDir := t.TempDir()
    inputFilePath := filepath.Join(tempDir, "input.json")

    inputJSON := `{"name": "Alice", "age": 30}`

    if err := ioutil.WriteFile(inputFilePath, []byte(inputJSON), 0644); err != nil {
        t.Fatalf("failed to write input file: %v", err)
    }

    var result map[string]interface{}
    if err := ReadJSONFile(inputFilePath, &result); err != nil {
        t.Fatalf("failed to read JSON file: %v", err)
    }

    expected := map[string]interface{}{
        "name": "Alice",
        "age":  float64(30),
    }

    if result["name"] != expected["name"] || result["age"] != expected["age"] {
        t.Errorf("unexpected result: got %v, want %v", result, expected)
    }
}

func TestWriteJSONFile(t *testing.T) {
    tempDir := t.TempDir()
    outputFilePath := filepath.Join(tempDir, "output.json")

    data := map[string]interface{}{
        "name": "Bob",
        "age":  25,
    }

    if err := WriteJSONFile(outputFilePath, data); err != nil {
        t.Fatalf("failed to write JSON file: %v", err)
    }

    result, err := ioutil.ReadFile(outputFilePath)
    if err != nil {
        t.Fatalf("failed to read output file: %v", err)
    }

    expected := `{
  "age": 25,
  "name": "Bob"
}`

    if string(result) != expected {
        t.Errorf("unexpected result: got %s, want %s", string(result), expected)
    }
}

func TestPrettyPrintJSON(t *testing.T) {
    data := map[string]interface{}{
        "name": "Charlie",
        "age":  35,
    }

    result, err := PrettyPrintJSON(data)
    if err != nil {
        t.Fatalf("failed to pretty print JSON: %v", err)
    }

    expected := `{
  "age": 35,
  "name": "Charlie"
}`

    if result != expected {
        t.Errorf("unexpected result: got %s, want %s", result, expected)
    }
}
