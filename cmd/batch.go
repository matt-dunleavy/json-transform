// cmd/batch.go
package cmd

import (
    "github.com/fatih/color"
    "github.com/matt-dunleavy/json-transform/config"
    "github.com/matt-dunleavy/json-transform/internal/batch"
    "github.com/spf13/cobra"
    "log"
)

var batchCmd = &cobra.Command{
    Use:   "batch",
    Short: "Perform batch operations on JSON files",
    Long:  `The batch command performs specified operations on multiple JSON files in a directory.`,
    Run: func(cmd *cobra.Command, args []string) {
        inputDir, _ := cmd.Flags().GetString("input")
        outputDir, _ := cmd.Flags().GetString("output")
        apiKey, _ := cmd.Flags().GetString("apikey")
        promptFile, _ := cmd.Flags().GetString("prompt")
        apiService, _ := cmd.Flags().GetString("service")
        model, _ := cmd.Flags().GetString("model")
        operation, _ := cmd.Flags().GetString("operation")

        // Use config values if CLI arguments are not provided
        if apiKey == "" {
            apiKey = config.Cfg.APIKey
        }
        if apiService == "" {
            apiService = config.Cfg.APIService
        }
        if model == "" {
            model = config.Cfg.Model
        }

        err := batch.ProcessBatchFiles(inputDir, outputDir, apiKey, promptFile, apiService, model, operation)
        if err != nil {
            log.Fatalf(color.RedString("Error processing batch files: %v"), err)
        }
        color.Green("Successfully processed batch files and saved to %s\n", outputDir)
    },
}

func init() {
    rootCmd.AddCommand(batchCmd)
    batchCmd.Flags().StringP("input", "i", "", "Input directory containing JSON files")
    batchCmd.Flags().StringP("output", "o", "", "Output directory for processed JSON data")
    batchCmd.Flags().StringP("apikey", "k", "", "API key for the AI service")
    batchCmd.Flags().StringP("prompt", "p", "", "File containing the prompt for AI processing")
    batchCmd.Flags().StringP("service", "s", "", "AI service to use (chatgpt, gemini)")
    batchCmd.Flags().StringP("model", "m", "", "AI model to use")
    batchCmd.Flags().String("operation", "", "Operation to perform (api-process)")
    batchCmd.MarkFlagRequired("input")
    batchCmd.MarkFlagRequired("output")
    batchCmd.MarkFlagRequired("prompt")
    batchCmd.MarkFlagRequired("operation")
}
