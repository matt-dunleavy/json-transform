// cmd/root.go
package cmd

import (
    "fmt"
    "github.com/fatih/color"
    "github.com/spf13/cobra"
    "os"
)

var rootCmd = &cobra.Command{
    Use:   "json-transform",
    Short: "A tool for batch processing and transforming JSON files",
    Long:  `json-transform is a CLI tool for performing various batch operations and transformations on JSON files.`,
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) == 0 {
            printHelp()
        } else {
            fmt.Println("Use json-transform --help for more information on available commands.")
        }
    },
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
    cobra.OnInitialize(initConfig)
}

func initConfig() {
    // Initialize configuration, if needed
}

func printHelp() {
    fmt.Println()
    color.New(color.FgGreen).Println("Welcome to json-transform!")
    color.New(color.FgCyan).Println("json-transform is a CLI tool for performing various batch operations and transformations on JSON files.")
    fmt.Println()

    color.New(color.FgYellow).Println("Available Commands:")
    fmt.Println()

    commands := []*cobra.Command{mergeCmd, splitCmd, filterCmd, aggregateCmd, sortCmd, searchCmd, exportCmd, compressCmd, decompressCmd}
    for _, cmd := range commands {
        color.New(color.FgGreen).Printf("  %s\t", cmd.Use)
        color.New(color.FgWhite).Printf("%s\n", cmd.Short)
    }

    fmt.Println()
    color.New(color.FgYellow).Println("Use 'json-transform [command] --help' for more information about a command.")
}
