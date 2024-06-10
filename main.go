package main

import (
    "github.com/matt-dunleavy/json-transform/cmd"
    "github.com/matt-dunleavy/json-transform/config"
)

func main() {
    config.LoadConfig("config.yaml")
    cmd.Execute()
}
