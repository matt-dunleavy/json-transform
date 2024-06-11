// cmd/root.go
package cmd

import (
    "fmt"
    "github.com/fatih/color"
    "github.com/matt-dunleavy/json-transform/config"
    "github.com/spf13/cobra"
    "os"
)

var cfgFile string

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
    rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yaml", "config file (default is config.yaml)")
}

func initConfig() {
    // Initialize and load configuration, if needed
    err := config.LoadConfig(cfgFile)
    if err != nil {
        color.Red("Failed to load config: %v\n", err)
        os.Exit(1)
    }

    color.Green("Successfully loaded config from %s\n", cfgFile)
    printConfig()
}

func printConfig() {
    configColor := color.New(color.FgCyan).SprintFunc()
    fmt.Printf("Loaded config:\n")
    for _, service := range config.Cfg.AIServices {
        fmt.Printf("  AI Service:       %s\n", configColor(service.Name))
        fmt.Printf("    API Key:        %s\n", configColor(service.APIKey))
        fmt.Printf("    Model:          %s\n", configColor(service.Model))
        fmt.Printf("    API URL:        %s\n", configColor(service.APIURL))
    }
    fmt.Printf("  Google Project ID: %s\n", configColor(config.Cfg.GoogleProjectID))
}

func printHelp() {
    fmt.Println()
    color.New(color.FgGreen).Println("JSON Transform v1.0.0")
    color.New(color.FgCyan).Println("json-transform is a CLI tool for performing various batch operations and transformations on JSON files.")
    fmt.Println()

    color.New(color.FgYellow).Println("Available Commands:")
    fmt.Println()

    commands := []*cobra.Command{mergeCmd, splitCmd, filterCmd, aggregateCmd, sortCmd, searchCmd, exportCmd, compressCmd, decompressCmd, configCmd}
    for _, cmd := range commands {
        color.New(color.FgGreen).Printf("  %s\t", cmd.Use)
        color.New(color.FgWhite).Printf("%s\n", cmd.Short)
    }

    fmt.Println()
    color.New(color.FgYellow).Println("Use 'json-transform [command] --help' for more information about a command.")
}
