// cmd/sort.go
package cmd

import (
    "github.com/fatih/color"
    "github.com/matt-dunleavy/json-transform/internal/sort"
    "github.com/spf13/cobra"
    "log"
)

var sortCmd = &cobra.Command{
    Use:   "sort",
    Short: "Sort JSON data by a specified field",
    Long:  `The sort command sorts JSON objects in a file by a specified field and outputs the sorted data.`,
    Run: func(cmd *cobra.Command, args []string) {
        inputFile, _ := cmd.Flags().GetString("input")
        outputFile, _ := cmd.Flags().GetString("output")
        field, _ := cmd.Flags().GetString("field")
        err := sort.SortJSONFile(inputFile, outputFile, field)
        if err != nil {
            log.Fatalf(color.RedString("Error sorting JSON file: %v"), err)
        }
        color.Green("Successfully sorted JSON data and saved to %s\n", outputFile)
    },
}

func init() {
    rootCmd.AddCommand(sortCmd)
    sortCmd.Flags().StringP("input", "i", "", "Input JSON file to sort")
    sortCmd.Flags().StringP("output", "o", "", "Output file for sorted JSON data")
    sortCmd.Flags().StringP("field", "f", "", "Field to sort by")
    sortCmd.MarkFlagRequired("input")
    sortCmd.MarkFlagRequired("output")
    sortCmd.MarkFlagRequired("field")
}
