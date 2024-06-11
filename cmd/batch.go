// cmd/batch.go
package cmd

import (
    "log"
    "github.com/spf13/cobra"
    "github.com/matt-dunleavy/json-transform/internal/batch"
)

var (
    inputDir      string
    outputDir     string
    promptFile    string
    operation     string
    outputFormat  string
    removeStyles  bool
    removeScripts bool
    removeMarkup  bool
    removeMarkdown bool
    removeCode    bool
    removeAll     bool
    logPayload    bool // New flag for logging payload
)

// batchCmd represents the batch command
var batchCmd = &cobra.Command{
    Use:   "batch",
    Short: "Batch process JSON files",
    Run: func(cmd *cobra.Command, args []string) {
        err := batch.ProcessBatch(inputDir, outputDir, promptFile, operation, outputFormat, logPayload)
        if err != nil {
            log.Fatalf("Error processing batch: %v", err)
        }
    },
}

func init() {
    rootCmd.AddCommand(batchCmd)

    batchCmd.PersistentFlags().StringVarP(&inputDir, "input", "i", "", "Input directory")
    batchCmd.PersistentFlags().StringVarP(&outputDir, "output", "o", "", "Output directory")
    batchCmd.PersistentFlags().StringVarP(&promptFile, "prompt", "p", "", "Prompt file (required for AI operations)")
    batchCmd.PersistentFlags().StringVarP(&operation, "operation", "O", "", "Operation to perform")
    batchCmd.PersistentFlags().StringVarP(&outputFormat, "output-format", "f", "json", "Output file format (e.g., json, txt, md)")
    batchCmd.PersistentFlags().BoolVar(&removeStyles, "remove-styles", false, "Remove styles (CSS/SCSS/SASS) from the JSON body")
    batchCmd.PersistentFlags().BoolVar(&removeScripts, "remove-scripts", false, "Remove scripts (JavaScript) from the JSON body")
    batchCmd.PersistentFlags().BoolVar(&removeMarkup, "remove-markup", false, "Remove HTML/XML markup from the JSON body")
    batchCmd.PersistentFlags().BoolVar(&removeMarkdown, "remove-markdown", false, "Remove Markdown syntax from the JSON body")
    batchCmd.PersistentFlags().BoolVar(&removeCode, "remove-code", false, "Remove all code (inline and blocks) from the JSON body")
    batchCmd.PersistentFlags().BoolVar(&removeAll, "remove-all", false, "Remove all artifacts, leaving only the plain text")
    batchCmd.PersistentFlags().BoolVar(&logPayload, "log-payload", false, "Log the API request payload")
}
