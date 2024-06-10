// internal/export/exporter.go
package export

import (
    "encoding/csv"
    "encoding/json"
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

// ExportJSONFile exports a JSON file to a specified format (csv, xml).
func ExportJSONFile(inputFile, outputFile, format string) error {
    data, err := ioutil.ReadFile(inputFile)
    if err != nil {
        return fmt.Errorf("failed to read input file: %w", err)
    }

    var jsonData []map[string]interface{}
    if err := json.Unmarshal(data, &jsonData); err != nil {
        return fmt.Errorf("failed to unmarshal JSON data: %w", err)
    }

    switch strings.ToLower(format) {
    case "csv":
        return exportToCSV(jsonData, outputFile)
    case "xml":
        return exportToXML(jsonData, outputFile)
    default:
        return fmt.Errorf("unsupported export format: %s", format)
    }
}

func exportToCSV(jsonData []map[string]interface{}, outputFile string) error {
    file, err := os.Create(outputFile)
    if err != nil {
        return fmt.Errorf("failed to create output file: %w", err)
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    if len(jsonData) == 0 {
        return nil
    }

    // Write header
    header := make([]string, 0, len(jsonData[0]))
    for key := range jsonData[0] {
        header = append(header, key)
    }
    if err := writer.Write(header); err != nil {
        return fmt.Errorf("failed to write CSV header: %w", err)
    }

    // Write rows
    for _, record := range jsonData {
        row := make([]string, len(header))
        for i, key := range header {
            if val, ok := record[key]; ok {
                row[i] = fmt.Sprintf("%v", val)
            }
        }
        if err := writer.Write(row); err != nil {
            return fmt.Errorf("failed to write CSV row: %w", err)
        }
    }

    return nil
}

func exportToXML(jsonData []map[string]interface{}, outputFile string) error {
    file, err := os.Create(outputFile)
    if err != nil {
        return fmt.Errorf("failed to create output file: %w", err)
    }
    defer file.Close()

    xmlData, err := xml.MarshalIndent(jsonData, "", "  ")
    if err != nil {
        return fmt.Errorf("failed to marshal XML data: %w", err)
    }

    if _, err := file.Write(xmlData); err != nil {
        return fmt.Errorf("failed to write XML data: %w", err)
    }

    return nil
}
