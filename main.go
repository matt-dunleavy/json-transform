// main.go
package main

import (
    "log"
    "github.com/matt-dunleavy/json-transform/cmd"
    "github.com/matt-dunleavy/json-transform/config"
)

func main() {
    // Load configuration
    config.LoadConfig("config.yaml")

    // Execute the root command
    if err := cmd.Execute(); err != nil {
        log.Fatalf("Error executing command: %v", err)
    }
}
