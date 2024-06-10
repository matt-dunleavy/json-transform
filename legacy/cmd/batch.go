// cmd/batch.go
package cmd

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"

    "github.com/matt-dunleavy/json-transform/internal/api"
    "github.com/spf13/cobra"
)

var batchCmd = &cobra.Command{
    Use:   "batch",
    Short: "Perform batch operations on JSON files",
    Long:  `The batch command allows you to perform various batch operations on JSON files, such as merging, splitting, filtering, and processing with external APIs.`,
    Run: func(cmd *cobra.Command, args []string) {
        inputDir, _ := cmd.Flags().GetString("input")
        outputDir, _ := cmd.Flags().GetString("output")
        operation, _ := cmd.Flags().GetString("operation")
        apiKey, _ := cmd.Flags().GetString("api-key")
        promptFile, _ := cmd.Flags().GetString("prompt")
        apiService, _ := cmd.Flags().GetString("api-service")
        model, _ := cmd.Flags().GetString("model")

        if operation == "api-process" {
            processWithAPI(inputDir, outputDir, apiKey, promptFile, apiService, model)
        } else {
            fmt.Println("Unsupported operation")
        }
    },
}

func init() {
    rootCmd.AddCommand(batchCmd)

    // Add flags for batch command
    batchCmd.Flags().StringP("input", "i", "", "Input directory containing JSON files")
    batchCmd.Flags().StringP("output", "o", "", "Output directory for processed JSON files")
    batchCmd.Flags().StringP("operation", "p", "", "Batch operation to perform (merge, split, filter, api-process)")
    batchCmd.Flags().StringP("api-key", "k", "", "API key for external API")
    batchCmd.Flags().StringP("prompt", "f", "", "File containing the prompt for API processing")
    batchCmd.Flags().StringP("api-service", "s", "chatgpt", "API service to use (chatgpt or gemini)")
    batchCmd.Flags().StringP("model", "m", "", "AI model to use")
    batchCmd.MarkFlagRequired("input")
    batchCmd.MarkFlagRequired("output")
    batchCmd.MarkFlagRequired("operation")
    batchCmd.MarkFlagRequired("api-key")
    batchCmd.MarkFlagRequired("prompt")
}

func processWithAPI(inputDir, outputDir, apiKey, promptFile, apiService, model string) {
    promptContent, err := ioutil.ReadFile(promptFile)
    if err != nil {
        log.Fatalf("Failed to read prompt file: %v", err)
    }

    files, err := ioutil.ReadDir(inputDir)
    if err != nil {
        log.Fatalf("Failed to read input directory: %v", err)
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

        response, err := api.ProcessJSONWithAPI(string(content), string(promptContent), apiKey, apiService, model)
        if err != nil {
            log.Printf("Failed to process file %s: %v", inputFilePath, err)
            continue
        }

        err = ioutil.WriteFile(outputFilePath, []byte(response), 0644)
        if err != nil {
            log.Printf("Failed to write output file %s: %v", outputFilePath, err)
            continue
        }

        log.Printf("Processed file %s and saved to %s", inputFilePath, outputFilePath)
        fmt.Printf("Response from API for file %s: %s\n", inputFilePath, response)
    }
}
