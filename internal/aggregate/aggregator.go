// internal/aggregate/aggregator.go
package aggregate

import (
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
)

// AggregateJSONFile performs the specified aggregation operation on a field in the JSON file.
func AggregateJSONFile(inputFile, outputFile, field, operation string) error {
    data, err := ioutil.ReadFile(inputFile)
    if err != nil {
        return fmt.Errorf("failed to read input file: %w", err)
    }

    var jsonData []map[string]interface{}
    if err := json.Unmarshal(data, &jsonData); err != nil {
        return fmt.Errorf("failed to unmarshal JSON data: %w", err)
    }

    var result float64
    var count int

    for _, item := range jsonData {
        value, ok := item[field].(float64)
        if !ok {
            return errors.New("field does not contain numeric values")
        }

        switch operation {
        case "sum":
            result += value
        case "avg":
            result += value
            count++
        case "count":
            count++
        default:
            return fmt.Errorf("unsupported operation: %s", operation)
        }
    }

    if operation == "avg" && count > 0 {
        result /= float64(count)
    } else if operation == "count" {
        result = float64(count)
    }

    resultData := map[string]float64{operation: result}
    resultJSON, err := json.MarshalIndent(resultData, "", "  ")
    if err != nil {
        return fmt.Errorf("failed to marshal result data: %w", err)
    }

    if err := ioutil.WriteFile(outputFile, resultJSON, 0644); err != nil {
        return fmt.Errorf("failed to write output file: %w", err)
    }

    return nil
}
