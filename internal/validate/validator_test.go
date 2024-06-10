// internal/validate/validator_test.go
package validate

import (
    "io/ioutil"
    "path/filepath"
    "testing"
)

func TestValidateJSONFile(t *testing.T) {
    tempDir := t.TempDir()
    inputFilePath := filepath.Join(tempDir, "input.json")
    schemaFilePath := filepath.Join(tempDir, "schema.json")

    inputJSON := `{
        "name": "Alice",
        "age": 30
    }`

    schemaJSON := `{
        "$schema": "http://json-schema.org/draft-07/schema#",
        "type": "object",
        "properties": {
            "name": {
                "type": "string"
            },
            "age": {
                "type": "integer"
            }
        },
        "required": ["name", "age"]
    }`

    if err := ioutil.WriteFile(inputFilePath, []byte(inputJSON), 0644); err != nil {
        t.Fatalf("failed to write input file: %v", err)
    }

    if err := ioutil.WriteFile(schemaFilePath, []byte(schemaJSON), 0644); err != nil {
        t.Fatalf("failed to write schema file: %v", err)
    }

    errors, err := ValidateJSONFile(inputFilePath, schemaFilePath)
    if err != nil {
        t.Fatalf("failed to validate JSON file: %v", err)
    }

    if len(errors) != 0 {
        t.Fatalf("expected no validation errors, got %v", errors)
    }
}

func TestValidateJSONFileWithErrors(t *testing.T) {
    tempDir := t.TempDir()
    inputFilePath := filepath.Join(tempDir, "input.json")
    schemaFilePath := filepath.Join(tempDir, "schema.json")

    inputJSON := `{
        "name": "Alice"
    }`

    schemaJSON := `{
        "$schema": "http://json-schema.org/draft-07/schema#",
        "type": "object",
        "properties": {
            "name": {
                "type": "string"
            },
            "age": {
                "type": "integer"
            }
        },
        "required": ["name", "age"]
    }`

    if err := ioutil.WriteFile(inputFilePath, []byte(inputJSON), 0644); err != nil {
        t.Fatalf("failed to write input file: %v", err)
    }

    if err := ioutil.WriteFile(schemaFilePath, []byte(schemaJSON), 0644); err != nil {
        t.Fatalf("failed to write schema file: %v", err)
    }

    errors, err := ValidateJSONFile(inputFilePath, schemaFilePath)
    if err != nil {
        t.Fatalf("failed to validate JSON file: %v", err)
    }

    if len(errors) == 0 {
        t.Fatalf("expected validation errors, got none")
    }

    expectedError := "age: age is required"
    if errors[0].String() != expectedError {
        t.Errorf("unexpected validation error: got %s, want %s", errors[0].String(), expectedError)
    }
}
