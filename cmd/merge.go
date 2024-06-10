// cmd/merge.go
package cmd

import (
    "github.com/fatih/color"
    "github.com/matt-dunleavy/json-transform/internal/merge"
    "github.com/spf13/cobra"
    "log"
)

var mergeCmd = &cobra.Command{
    Use:   "merge",
    Short: "Merge multiple JSON files into one",
    Long:  `The merge command combines multiple JSON files into a single JSON file.`,
    Run: func(cmd *cobra.Command, args []string) {
        inputDir, _ := cmd.Flags().GetString("input")
        outputFile, _ := cmd.Flags().GetString("output")
        err := merge.MergeJSONFiles(inputDir, outputFile)
        if err != nil {
            log.Fatalf(color.RedString("Error merging JSON files: %v"), err)
        }
        color.Green("Successfully merged JSON files into %s\n", outputFile)
    },
}

func init() {
    rootCmd.AddCommand(mergeCmd)
    mergeCmd.Flags().StringP("input", "i", "", "Input directory containing JSON files")
    mergeCmd.Flags().StringP("output", "o", "", "Output file for merged JSON data")
    mergeCmd.MarkFlagRequired("input")
    mergeCmd.MarkFlagRequired("output")
}
