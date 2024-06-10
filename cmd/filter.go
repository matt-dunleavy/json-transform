// cmd/filter.go
package cmd

import (
    "github.com/fatih/color"
    "github.com/matt-dunleavy/json-transform/internal/filter"
    "github.com/spf13/cobra"
    "log"
)

var filterCmd = &cobra.Command{
    Use:   "filter",
    Short: "Filter JSON data based on criteria",
    Long:  `The filter command extracts specific data from JSON files based on defined filter criteria.`,
    Run: func(cmd *cobra.Command, args []string) {
        inputFile, _ := cmd.Flags().GetString("input")
        outputFile, _ := cmd.Flags().GetString("output")
        criteria, _ := cmd.Flags().GetString("criteria")
        err := filter.FilterJSONFile(inputFile, outputFile, criteria)
        if err != nil {
            log.Fatalf(color.RedString("Error filtering JSON file: %v"), err)
        }
        color.Green("Successfully filtered JSON data and saved to %s\n", outputFile)
    },
}

func init() {
    rootCmd.AddCommand(filterCmd)
    filterCmd.Flags().StringP("input", "i", "", "Input JSON file to filter")
    filterCmd.Flags().StringP("output", "o", "", "Output file for filtered JSON data")
    filterCmd.Flags().StringP("criteria", "c", "", "Filter criteria")
    filterCmd.MarkFlagRequired("input")
    filterCmd.MarkFlagRequired("output")
    filterCmd.MarkFlagRequired("criteria")
}
