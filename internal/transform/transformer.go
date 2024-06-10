// internal/transform/transformer.go
package transform

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

// TransformJSONFile transforms the JSON data in the input file based on the provided rules and writes the transformed data to the output file.
func TransformJSONFile(inputFile, outputFile, rulesFile string) error {
    inputData, err := ioutil.ReadFile(inputFile)
    if err != nil {
        return fmt.Errorf("failed to read input file: %w", err)
    }

    var jsonData []map[string]interface{}
    if err := json.Unmarshal(inputData, &jsonData); err != nil {
        return fmt.Errorf("failed to unmarshal JSON data: %w", err)
    }

    rulesData, err := ioutil.ReadFile(rulesFile)
    if err != nil {
        return fmt.Errorf("failed to read rules file: %w", err)
    }

    var rules map[string]interface{}
    if err := json.Unmarshal(rulesData, &rules); err != nil {
        return fmt.Errorf("failed to unmarshal rules: %w", err)
    }

    transformedData := applyRules(jsonData, rules)

    transformedJSON, err := json.MarshalIndent(transformedData, "", "  ")
    if err != nil {
        return fmt.Errorf("failed to marshal transformed data: %w", err)
    }

    if err := ioutil.WriteFile(outputFile, transformedJSON, 0644); err != nil {
        return fmt.Errorf("failed to write output file: %w", err)
    }

    return nil
}

// applyRules applies the transformation rules to the JSON data.
func applyRules(data []map[string]interface{}, rules map[string]interface{}) []map[string]interface{} {
    // This is a placeholder implementation. Actual transformation logic will depend on the rules.
    // Example: Renaming a field
    if renameRule, ok := rules["rename"].(map[string]interface{}); ok {
        oldField := renameRule["old"].(string)
        newField := renameRule["new"].(string)
        for _, item := range data {
            if value, exists := item[oldField]; exists {
                item[newField] = value
                delete(item, oldField)
            }
        }
    }
    return data
}
