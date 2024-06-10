// internal/validate/validator.go
package validate

import (
    "encoding/json"
    "fmt"
    "io/ioutil"

    "github.com/xeipuuv/gojsonschema"
)

// ValidateJSONFile validates the JSON data in the input file against the provided JSON schema.
func ValidateJSONFile(inputFile, schemaFile string) ([]gojsonschema.ResultError, error) {
    data, err := ioutil.ReadFile(inputFile)
    if err != nil {
        return nil, fmt.Errorf("failed to read input file: %w", err)
    }

    schemaData, err := ioutil.ReadFile(schemaFile)
    if err != nil {
        return nil, fmt.Errorf("failed to read schema file: %w", err)
    }

    schemaLoader := gojsonschema.NewStringLoader(string(schemaData))
    documentLoader := gojsonschema.NewStringLoader(string(data))

    result, err := gojsonschema.Validate(schemaLoader, documentLoader)
    if err != nil {
        return nil, fmt.Errorf("failed to validate JSON: %w", err)
    }

    if result.Valid() {
        return nil, nil
    }

    return result.Errors(), nil
}
