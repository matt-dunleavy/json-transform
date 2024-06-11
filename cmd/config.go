// cmd/config.go
package cmd

import (
    "fmt"
    "github.com/fatih/color"
    "github.com/matt-dunleavy/json-transform/config"
    "github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
    Use:   "config",
    Short: "Display the current configuration",
    Run: func(cmd *cobra.Command, args []string) {
        displayConfig()
    },
}

func displayConfig() {
    configColor := color.New(color.FgCyan).SprintFunc()
    fmt.Printf("Current Configuration:\n")
    for _, service := range config.Cfg.AIServices {
        fmt.Printf("  AI Service:       %s\n", configColor(service.Name))
        fmt.Printf("    API Key:        %s\n", configColor(service.APIKey))
        fmt.Printf("    Model:          %s\n", configColor(service.Model))
        fmt.Printf("    API URL:        %s\n", configColor(service.APIURL))
    }
    fmt.Printf("  Google Project ID: %s\n", configColor(config.Cfg.GoogleProjectID))
}

func init() {
    rootCmd.AddCommand(configCmd)
}
