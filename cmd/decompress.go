// cmd/decompress.go
package cmd

import (
    "github.com/fatih/color"
    "github.com/matt-dunleavy/json-transform/internal/decompress"
    "github.com/spf13/cobra"
    "log"
)

var decompressCmd = &cobra.Command{
    Use:   "decompress",
    Short: "Decompress JSON files",
    Long:  `The decompress command decompresses JSON files to make them readable again.`,
    Run: func(cmd *cobra.Command, args []string) {
        inputFile, _ := cmd.Flags().GetString("input")
        outputFile, _ := cmd.Flags().GetString("output")
        err := decompress.DecompressJSONFile(inputFile, outputFile)
        if err != nil {
            log.Fatalf(color.RedString("Error decompressing JSON file: %v"), err)
        }
        color.Green("Successfully decompressed JSON file and saved to %s\n", outputFile)
    },
}

func init() {
    rootCmd.AddCommand(decompressCmd)
    decompressCmd.Flags().StringP("input", "i", "", "Input JSON file to decompress")
    decompressCmd.Flags().StringP("output", "o", "", "Output file for decompressed data")
    decompressCmd.MarkFlagRequired("input")
    decompressCmd.MarkFlagRequired("output")
}
