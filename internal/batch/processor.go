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
    "github.com/matt-dunleavy/json-transform/internal/strip"
    "github.com/matt-dunleavy/json-transform/config"
)

var (
    removeStyles    bool
    removeScripts   bool
    removeMarkup    bool
    removeMarkdown  bool
    removeCode      bool
    removeAll       bool
)

func ProcessBatch(inputDir, outputDir, promptFile, operation, outputFormat string, logPayload bool) error {
    var prompt []byte
    var err error

    // Read the prompt file if specified and required by the operation
    if operationRequiresPrompt(operation) {
        if promptFile == "" {
            return fmt.Errorf("prompt file is required for operation %s", operation)
        }
        prompt, err = ioutil.ReadFile(promptFile)
        if err != nil {
            return fmt.Errorf("failed to read prompt file: %w", err)
        }
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

            // Perform artifact stripping based on flags
            content := string(input)
            if removeAll {
                content = strip.StripAll(content)
            } else {
                if removeStyles {
                    content = strip.StripStyles(content)
                }
                if removeScripts {
                    content = strip.StripScripts(content)
                }
                if removeMarkup {
                    content = strip.StripMarkup(content)
                }
                if removeMarkdown {
                    content = strip.StripMarkdown(content)
                }
                if removeCode {
                    content = strip.StripCode(content)
                }
            }

            var result string
            if operationRequiresPrompt(operation) {
                result, err = api.ProcessJSONWithAPI(content, string(prompt), config.Cfg.APIKey, config.Cfg.APIService, config.Cfg.Model, logPayload)
                if err != nil {
                    return fmt.Errorf("failed to process file %s: %w", path, err)
                }
            } else {
                result = content
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

func operationRequiresPrompt(operation string) bool {
    aiOperations := []string{"api-process"} // Add more AI operations if necessary
    for _, op := range aiOperations {
        if operation == op {
            return true
        }
    }
    return false
}
