// cmd/compress.go
package cmd

import (
    "github.com/fatih/color"
    "github.com/matt-dunleavy/json-transform/internal/compress"
    "github.com/spf13/cobra"
    "log"
)

var compressCmd = &cobra.Command{
    Use:   "compress",
    Short: "Compress JSON files",
    Long:  `The compress command compresses JSON files to reduce storage space.`,
    Run: func(cmd *cobra.Command, args []string) {
        inputFile, _ := cmd.Flags().GetString("input")
        outputFile, _ := cmd.Flags().GetString("output")
        err := compress.CompressJSONFile(inputFile, outputFile)
        if err != nil {
            log.Fatalf(color.RedString("Error compressing JSON file: %v"), err)
        }
        color.Green("Successfully compressed JSON file and saved to %s\n", outputFile)
    },
}

func init() {
    rootCmd.AddCommand(compressCmd)
    compressCmd.Flags().StringP("input", "i", "", "Input JSON file to compress")
    compressCmd.Flags().StringP("output", "o", "", "Output file for compressed data")
    compressCmd.MarkFlagRequired("input")
    compressCmd.MarkFlagRequired("output")
}
