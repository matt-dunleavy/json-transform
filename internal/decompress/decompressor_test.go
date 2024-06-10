// internal/decompress/decompressor_test.go
package decompress

import (
    "compress/gzip"
    "io/ioutil"
    "os"
    "path/filepath"
    "testing"
)

func TestDecompressJSONFile(t *testing.T) {
    tempDir := t.TempDir()
    inputFilePath := filepath.Join(tempDir, "input.json.gz")
    outputFilePath := filepath.Join(tempDir, "output.json")

    inputJSON := `{"key": "value"}`

    // Create compressed file
    input, err := os.Create(inputFilePath)
    if err != nil {
        t.Fatalf("failed to create input file: %v", err)
    }
    writer := gzip.NewWriter(input)
    _, err = writer.Write([]byte(inputJSON))
    if err != nil {
        t.Fatalf("failed to write compressed data: %v", err)
    }
    writer.Close()
    input.Close()

    // Test decompression
    if err := DecompressJSONFile(inputFilePath, outputFilePath); err != nil {
        t.Fatalf("failed to decompress JSON file: %v", err)
    }

    result, err := ioutil.ReadFile(outputFilePath)
    if err != nil {
        t.Fatalf("failed to read output file: %v", err)
    }

    if string(result) != inputJSON {
        t.Fatalf("unexpected result: got %s, want %s", string(result), inputJSON)
    }
}
