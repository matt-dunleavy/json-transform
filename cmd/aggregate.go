// cmd/aggregate.go
package cmd

import (
    "github.com/fatih/color"
    "github.com/matt-dunleavy/json-transform/internal/aggregate"
    "github.com/spf13/cobra"
    "log"
)

var aggregateCmd = &cobra.Command{
    Use:   "aggregate",
    Short: "Perform aggregation operations on JSON data",
    Long:  `The aggregate command performs aggregation operations such as sum, average, count, etc., on specific fields in JSON files.`,
    Run: func(cmd *cobra.Command, args []string) {
        inputFile, _ := cmd.Flags().GetString("input")
        outputFile, _ := cmd.Flags().GetString("output")
        field, _ := cmd.Flags().GetString("field")
        operation, _ := cmd.Flags().GetString("operation")
        err := aggregate.AggregateJSONFile(inputFile, outputFile, field, operation)
        if err != nil {
            log.Fatalf(color.RedString("Error aggregating JSON file: %v"), err)
        }
        color.Green("Successfully aggregated JSON data and saved to %s\n", outputFile)
    },
}

func init() {
    rootCmd.AddCommand(aggregateCmd)
    aggregateCmd.Flags().StringP("input", "i", "", "Input JSON file to aggregate")
    aggregateCmd.Flags().StringP("output", "o", "", "Output file for aggregated JSON data")
    aggregateCmd.Flags().StringP("field", "f", "", "Field to aggregate")
    aggregateCmd.Flags().StringP("operation", "p", "", "Aggregation operation (sum, avg, count)")
    aggregateCmd.MarkFlagRequired("input")
    aggregateCmd.MarkFlagRequired("output")
    aggregateCmd.MarkFlagRequired("field")
    aggregateCmd.MarkFlagRequired("operation")
}
