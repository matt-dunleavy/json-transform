// internal/export/exporter_test.go
package export

import (
    "encoding/csv"
    "encoding/xml"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
    "testing"
)

func TestExportJSONFileToCSV(t *testing.T) {
    tempDir := t.TempDir()
    inputFilePath := filepath.Join(tempDir, "input.json")
    outputFilePath := filepath.Join(tempDir, "output.csv")

    inputJSON := `[
        {"name": "Alice", "age": 30},
        {"name": "Bob", "age": 25}
    ]`

    if err := ioutil.WriteFile(inputFilePath, []byte(inputJSON), 0644); err != nil {
        t.Fatalf("failed to write input file: %v", err)
    }

    if err := ExportJSONFile(inputFilePath, outputFilePath, "csv"); err != nil {
        t.Fatalf("failed to export JSON file to CSV: %v", err)
    }

    file, err := os.Open(outputFilePath)
    if err != nil {
        t.Fatalf("failed to open output file: %v", err)
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        t.Fatalf("failed to read CSV data: %v", err)
    }

    expected := [][]string{
        {"name", "age"},
        {"Alice", "30"},
        {"Bob", "25"},
    }

    for i, row := range records {
        for j, value := range row {
            if value != expected[i][j] {
                t.Errorf("unexpected value at row %d, col %d: got %s, want %s", i, j, value, expected[i][j])
            }
        }
    }
}

func TestExportJSONFileToXML(t *testing.T) {
    tempDir := t.TempDir()
    inputFilePath := filepath.Join(tempDir, "input.json")
    outputFilePath := filepath.Join(tempDir, "output.xml")

    inputJSON := `[
        {"name": "Alice", "age": 30},
        {"name": "Bob", "age": 25}
    ]`

    if err := ioutil.WriteFile(inputFilePath, []byte(inputJSON), 0644); err != nil {
        t.Fatalf("failed to write input file: %v", err)
    }

    if err := ExportJSONFile(inputFilePath, outputFilePath, "xml"); err != nil {
        t.Fatalf("failed to export JSON file to XML: %v", err)
    }

    file, err := os.Open(outputFilePath)
    if err != nil {
        t.Fatalf("failed to open output file: %v", err)
    }
    defer file.Close()

    result, err := ioutil.ReadAll(file)
    if err != nil {
        t.Fatalf("failed to read XML data: %v", err)
    }

    expected := `<ArrayOfMap>
  <Map>
    <age>30</age>
    <name>Alice</name>
  </Map>
  <Map>
    <age>25</age>
    <name>Bob</name>
  </Map>
</ArrayOfMap>`

    if strings.TrimSpace(string(result)) != expected {
        t.Errorf("unexpected result: got %s, want %s", string(result), expected)
    }
}
