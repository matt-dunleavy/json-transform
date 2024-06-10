// cmd/split.go
package cmd

import (
    "github.com/fatih/color"
    "github.com/matt-dunleavy/json-transform/internal/split"
    "github.com/spf13/cobra"
    "log"
)

var splitCmd = &cobra.Command{
    Use:   "split",
    Short: "Split JSON files into smaller chunks",
    Long:  `The split command splits a large JSON file into smaller files with a specified number of objects per file.`,
    Run: func(cmd *cobra.Command, args []string) {
        inputFile, _ := cmd.Flags().GetString("input")
        outputDir, _ := cmd.Flags().GetString("output")
        chunkSize, _ := cmd.Flags().GetInt("chunk")
        err := split.SplitJSONFile(inputFile, outputDir, chunkSize)
        if err != nil {
            log.Fatalf(color.RedString("Error splitting JSON file: %v"), err)
        }
        color.Green("Successfully split JSON data into chunks and saved to %s\n", outputDir)
    },
}

func init() {
    rootCmd.AddCommand(splitCmd)
    splitCmd.Flags().StringP("input", "i", "", "Input JSON file to split")
    splitCmd.Flags().StringP("output", "o", "", "Output directory for split JSON data")
    splitCmd.Flags().IntP("chunk", "c", 0, "Number of objects per chunk")
    splitCmd.MarkFlagRequired("input")
    splitCmd.MarkFlagRequired("output")
    splitCmd.MarkFlagRequired("chunk")
}
