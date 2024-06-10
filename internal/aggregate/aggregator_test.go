// internal/aggregate/aggregator_test.go
package aggregate

import (
    "io/ioutil"
    "os"
    "path/filepath"
    "testing"
)

func TestAggregateJSONFile(t *testing.T) {
    tempDir := t.TempDir()
    inputFilePath := filepath.Join(tempDir, "input.json")
    outputFilePath := filepath.Join(tempDir, "output.json")

    inputJSON := `[
        {"value": 10},
        {"value": 20},
        {"value": 30}
    ]`

    if err := ioutil.WriteFile(inputFilePath, []byte(inputJSON), 0644); err != nil {
        t.Fatalf("failed to write input file: %v", err)
    }

    tests := []struct {
        operation string
        expected  string
    }{
        {"sum", `{"sum":60}`},
        {"avg", `{"avg":20}`},
        {"count", `{"count":3}`},
    }

    for _, tt := range tests {
        t.Run(tt.operation, func(t *testing.T) {
            if err := AggregateJSONFile(inputFilePath, outputFilePath, "value", tt.operation); err != nil {
                t.Fatalf("failed to aggregate JSON file: %v", err)
            }

            result, err := ioutil.ReadFile(outputFilePath)
            if err != nil {
                t.Fatalf("failed to read output file: %v", err)
            }

            expectedResult := fmt.Sprintf("%s\n", tt.expected)
            if string(result) != expectedResult {
                t.Fatalf("unexpected result: got %s, want %s", result, expectedResult)
            }
        })
    }
}
