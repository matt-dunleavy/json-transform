#!/bin/bash

# Build the application
go build -o json-transform main.go

# Check if the build was successful
if [ $? -eq 0 ]; then
  echo "Build successful. The executable is named 'json-transform'."
else
  echo "Build failed."
  exit 1
fi
