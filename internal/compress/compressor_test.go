// internal/compress/compressor_test.go
package compress

import (
    "compress/gzip"
    "io/ioutil"
    "os"
    "path/filepath"
    "testing"
)

func TestCompressJSONFile(t *testing.T) {
    tempDir := t.TempDir()
    inputFilePath := filepath.Join(tempDir, "input.json")
    outputFilePath := filepath.Join(tempDir, "output.json.gz")

    inputJSON := `{"key": "value"}`

    if err := ioutil.WriteFile(inputFilePath, []byte(inputJSON), 0644); err != nil {
        t.Fatalf("failed to write input file: %v", err)
    }

    if err := CompressJSONFile(inputFilePath, outputFilePath); err != nil {
        t.Fatalf("failed to compress JSON file: %v", err)
    }

    output, err := os.Open(outputFilePath)
    if err != nil {
        t.Fatalf("failed to open output file: %v", err)
    }
    defer output.Close()

    reader, err := gzip.NewReader(output)
    if err != nil {
        t.Fatalf("failed to create gzip reader: %v", err)
    }
    defer reader.Close()

    decompressedData, err := ioutil.ReadAll(reader)
    if err != nil {
        t.Fatalf("failed to read decompressed data: %v", err)
    }

    if string(decompressedData) != inputJSON {
        t.Fatalf("unexpected result: got %s, want %s", string(decompressedData), inputJSON)
    }
}
