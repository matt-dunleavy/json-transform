// internal/split/splitter_test.go
package split

import (
    "io/ioutil"
    "os"
    "path/filepath"
    "testing"
)

func TestSplitJSONFile(t *testing.T) {
    tempDir := t.TempDir()
    inputFilePath := filepath.Join(tempDir, "input.json")
    outputDir := filepath.Join(tempDir, "output")

    os.Mkdir(outputDir, 0755)

    inputJSON := `[
        {"name": "Alice", "age": 30},
        {"name": "Bob", "age": 25},
        {"name": "Charlie", "age": 35},
        {"name": "Dave", "age": 40},
        {"name": "Eve", "age": 28}
    ]`

    if err := ioutil.WriteFile(inputFilePath, []byte(inputJSON), 0644); err != nil {
        t.Fatalf("failed to write input file: %v", err)
    }

    chunkSize := 2

    if err := SplitJSONFile(inputFilePath, outputDir, chunkSize); err != nil {
        t.Fatalf("failed to split JSON file: %v", err)
    }

    expectedChunks := []string{
        `[
  {
    "age": 30,
    "name": "Alice"
  },
  {
    "age": 25,
    "name": "Bob"
  }
]`,
        `[
  {
    "age": 35,
    "name": "Charlie"
  },
  {
    "age": 40,
    "name": "Dave"
  }
]`,
        `[
  {
    "age": 28,
    "name": "Eve"
  }
]`,
    }

    for i, expected := range expectedChunks {
        chunkFilePath := filepath.Join(outputDir, fmt.Sprintf("chunk_%d.json", i))
        result, err := ioutil.ReadFile(chunkFilePath)
        if err != nil {
            t.Fatalf("failed to read chunk file %s: %v", chunkFilePath, err)
        }

        if string(result) != expected {
            t.Errorf("unexpected result in chunk %d: got %s, want %s", i, string(result), expected)
        }
    }
}
