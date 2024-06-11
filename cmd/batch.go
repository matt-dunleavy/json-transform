// cmd/batch.go
package cmd

import (
    "log"
    "github.com/spf13/cobra"
    "github.com/matt-dunleavy/json-transform/internal/batch"
)

var (
    inputDir     string
    outputDir    string
    promptFile   string
    operation    string
    outputFormat string
)

// batchCmd represents the batch command
var batchCmd = &cobra.Command{
    Use:   "batch",
    Short: "Batch process JSON files",
    Run: func(cmd *cobra.Command, args []string) {
        err := batch.ProcessBatch(inputDir, outputDir, promptFile, operation, outputFormat)
        if err != nil {
            log.Fatalf("Error processing batch: %v", err)
        }
    },
}

func init() {
    rootCmd.AddCommand(batchCmd)

    batchCmd.PersistentFlags().StringVarP(&inputDir, "input", "i", "", "Input directory")
    batchCmd.PersistentFlags().StringVarP(&outputDir, "output", "o", "", "Output directory")
    batchCmd.PersistentFlags().StringVarP(&promptFile, "prompt", "p", "", "Prompt file")
    batchCmd.PersistentFlags().StringVarP(&operation, "operation", "O", "", "Operation to perform")
    batchCmd.PersistentFlags().StringVarP(&outputFormat, "output-format", "f", "md", "Output file format (e.g., json, txt, md)")
}
