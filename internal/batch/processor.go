// internal/batch/processor.go
package batch

import (
    "fmt"
    "io/ioutil"
    "log"
    "path/filepath"

    "github.com/matt-dunleavy/json-transform/internal/api"
)

// ProcessBatchFiles processes multiple JSON files with a given operation and parameters.
func ProcessBatchFiles(inputDir, outputDir, apiKey, promptFile, apiService, model, operation string) error {
    promptContent, err := ioutil.ReadFile(promptFile)
    if err != nil {
        return fmt.Errorf("failed to read prompt file: %w", err)
    }

    files, err := ioutil.ReadDir(inputDir)
    if err != nil {
        return fmt.Errorf("failed to read input directory: %w", err)
    }

    for _, file := range files {
        if file.IsDir() {
            continue
        }

        inputFilePath := filepath.Join(inputDir, file.Name())
        outputFilePath := filepath.Join(outputDir, file.Name())

        content, err := ioutil.ReadFile(inputFilePath)
        if err != nil {
            log.Printf("Failed to read file %s: %v", inputFilePath, err)
            continue
        }

        var response string
        switch operation {
        case "api-process":
            response, err = api.ProcessJSONWithAPI(string(content), string(promptContent), apiKey, apiService, model)
            if err != nil {
                log.Printf("Failed to process file %s: %v", inputFilePath, err)
                continue
            }
        default:
            return fmt.Errorf("unsupported operation: %s", operation)
        }

        err = ioutil.WriteFile(outputFilePath, []byte(response), 0644)
        if err != nil {
            log.Printf("Failed to write output file %s: %v", outputFilePath, err)
            continue
        }

        log.Printf("Processed file %s and saved to %s", inputFilePath, outputFilePath)
    }

    return nil
}
