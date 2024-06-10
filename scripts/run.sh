#!/bin/bash

# Check if a configuration file exists
if [ ! -f "./config.yaml" ]; then
  echo "Configuration file not found. Please create a config.yaml file."
  exit 1
fi

# Run the application
go run main.go "$@"
