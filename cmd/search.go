// cmd/search.go
package cmd

import (
    "github.com/fatih/color"
    "github.com/matt-dunleavy/json-transform/internal/search"
    "github.com/spf13/cobra"
    "log"
)

var searchCmd = &cobra.Command{
    Use:   "search",
    Short: "Search for specific values or patterns within JSON files",
    Long:  `The search command searches for specific values or patterns within JSON files and highlights or extracts matching data.`,
    Run: func(cmd *cobra.Command, args []string) {
        inputFile, _ := cmd.Flags().GetString("input")
        outputFile, _ := cmd.Flags().GetString("output")
        pattern, _ := cmd.Flags().GetString("pattern")
        err := search.SearchJSONFile(inputFile, outputFile, pattern)
        if err != nil {
            log.Fatalf(color.RedString("Error searching JSON file: %v"), err)
        }
        color.Green("Successfully searched JSON data and saved to %s\n", outputFile)
    },
}

func init() {
    rootCmd.AddCommand(searchCmd)
    searchCmd.Flags().StringP("input", "i", "", "Input JSON file to search")
    searchCmd.Flags().StringP("output", "o", "", "Output file for search results")
    searchCmd.Flags().StringP("pattern", "p", "", "Pattern to search for")
    searchCmd.MarkFlagRequired("input")
    searchCmd.MarkFlagRequired("output")
    searchCmd.MarkFlagRequired("pattern")
}
