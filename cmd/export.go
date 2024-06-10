// cmd/export.go
package cmd

import (
    "github.com/fatih/color"
    "github.com/matt-dunleavy/json-transform/internal/export"
    "github.com/spf13/cobra"
    "log"
)

var exportCmd = &cobra.Command{
    Use:   "export",
    Short: "Export JSON data to other formats",
    Long:  `The export command converts JSON files to other formats (e.g., CSV, XML) for easier analysis or integration with other systems.`,
    Run: func(cmd *cobra.Command, args []string) {
        inputFile, _ := cmd.Flags().GetString("input")
        outputFile, _ := cmd.Flags().GetString("output")
        format, _ := cmd.Flags().GetString("format")
        err := export.ExportJSONFile(inputFile, outputFile, format)
        if err != nil {
            log.Fatalf(color.RedString("Error exporting JSON file: %v"), err)
        }
        color.Green("Successfully exported JSON data to %s\n", outputFile)
    },
}

func init() {
    rootCmd.AddCommand(exportCmd)
    exportCmd.Flags().StringP("input", "i", "", "Input JSON file to export")
    exportCmd.Flags().StringP("output", "o", "", "Output file for exported data")
    exportCmd.Flags().StringP("format", "f", "", "Export format (csv, xml)")
    exportCmd.MarkFlagRequired("input")
    exportCmd.MarkFlagRequired("output")
    exportCmd.MarkFlagRequired("format")
}
