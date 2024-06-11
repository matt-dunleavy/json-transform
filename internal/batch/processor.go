// internal/batch/processor.go
package batch

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"

    "github.com/matt-dunleavy/json-transform/internal/api"
    "github.com/matt-dunleavy/json-transform/config"
)

func ProcessBatch(inputDir, outputDir, promptFile, operation, outputFormat string) error {
    // Read the prompt file
    prompt, err := ioutil.ReadFile(promptFile)
    if err != nil {
        return fmt.Errorf("failed to read prompt file: %w", err)
    }

    // Ensure output directory exists
    err = os.MkdirAll(outputDir, os.ModePerm)
    if err != nil {
        return fmt.Errorf("failed to create output directory: %w", err)
    }

    // Process each file in the input directory
    err = filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
            log.Printf("Processing file: %s", path)
            input, err := ioutil.ReadFile(path)
            if err != nil {
                return fmt.Errorf("failed to read input file: %w", err)
            }

            result, err := api.ProcessJSONWithAPI(string(input), string(prompt), config.Cfg.APIKey, config.Cfg.APIService, config.Cfg.Model)
            if err != nil {
                return fmt.Errorf("failed to process file %s: %w", path, err)
            }

            // Determine the output file name and extension
            outputFileName := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name())) + "." + outputFormat
            outputFile := filepath.Join(outputDir, outputFileName)

            // Write the result to the output file
            err = ioutil.WriteFile(outputFile, []byte(result), 0644)
            if err != nil {
                return fmt.Errorf("failed to write output file: %w", err)
            }

            log.Printf("Successfully processed file: %s", outputFile)
        }
        return nil
    })

    return err
}
