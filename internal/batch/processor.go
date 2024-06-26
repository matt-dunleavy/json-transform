// internal/batch/processor.go
package batch

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"

    "github.com/fatih/color"
    "github.com/matt-dunleavy/json-transform/config"
    "github.com/matt-dunleavy/json-transform/internal/api"
    "github.com/matt-dunleavy/json-transform/internal/strip"
)

var (
    removeStyles    bool
    removeScripts   bool
    removeMarkup    bool
    removeMarkdown  bool
    removeCode      bool
    removeAll       bool
)

func init() {
    // Set log flags to include date and time
    log.SetFlags(log.LstdFlags)
}

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

    // Get the AI service configuration
    aiService, err := config.GetAIServiceByName(config.Cfg.AIServices[0].Name) // Assuming the first service for now
    if err != nil {
        return fmt.Errorf("failed to get AI service configuration: %w", err)
    }

    // Process each file in the input directory
    files, err := ioutil.ReadDir(inputDir)
    if err != nil {
        return fmt.Errorf("failed to read input directory: %w", err)
    }

    successColor := color.New(color.FgGreen).SprintFunc()
    fileColor := color.New(color.FgBlue).SprintFunc()
    processColor := color.New(color.FgYellow).SprintFunc()
    errorColor := color.New(color.FgRed).SprintFunc()

    for _, file := range files {
        if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
            log.Printf("%s %s", processColor("Processing file:"), fileColor(file.Name()))
            input, err := ioutil.ReadFile(filepath.Join(inputDir, file.Name()))
            if err != nil {
                log.Printf("Failed to read input file: %s [%s]", errorColor(err), time.Now().Format("2006/01/02 15:04:05"))
                continue
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
                result, err = api.ProcessJSONWithAPI(content, string(prompt), aiService.APIKey, aiService.Name, aiService.Model, logPayload)
                if err != nil {
                    log.Printf("Failed to process file %s: %s [%s]", file.Name(), errorColor(err), time.Now().Format("2006/01/02 15:04:05"))
                    continue
                }
            } else {
                result = content
            }

            // Determine the output file name and extension
            outputFileName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())) + "." + outputFormat
            outputFile := filepath.Join(outputDir, outputFileName)

            // Write the result to the output file
            err = ioutil.WriteFile(outputFile, []byte(result), 0644)
            if err != nil {
                log.Printf("Failed to write output file: %s [%s]", errorColor(err), time.Now().Format("2006/01/02 15:04:05"))
                continue
            }

            log.Printf("%s %s", successColor("Successfully processed file:"), fileColor(outputFile))
        }
    }

    return nil
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
