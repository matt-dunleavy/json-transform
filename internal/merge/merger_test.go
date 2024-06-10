// internal/merge/merger_test.go
package merge

import (
    "io/ioutil"
    "os"
    "path/filepath"
    "testing"
)

func TestMergeJSONFiles(t *testing.T) {
    tempDir := t.TempDir()
    inputDir := filepath.Join(tempDir, "input")
    outputFilePath := filepath.Join(tempDir, "output.json")

    os.Mkdir(inputDir, 0755)

    inputFiles := map[string]string{
        "file1.json": `[{"name": "Alice", "age": 30}]`,
        "file2.json": `[{"name": "Bob", "age": 25}]`,
        "file3.json": `[{"name": "Charlie", "age": 35}]`,
    }

    for fileName, content := range inputFiles {
        filePath := filepath.Join(inputDir, fileName)
        if err := ioutil.WriteFile(filePath, []byte(content), 0644); err != nil {
            t.Fatalf("failed to write input file %s: %v", filePath, err)
        }
    }

    if err := MergeJSONFiles(inputDir, outputFilePath); err != nil {
        t.Fatalf("failed to merge JSON files: %v", err)
    }

    result, err := ioutil.ReadFile(outputFilePath)
    if err != nil {
        t.Fatalf("failed to read output file: %v", err)
    }

    expected := `[
  {
    "age": 30,
    "name": "Alice"
  },
  {
    "age": 25,
    "name": "Bob"
  },
  {
    "age": 35,
    "name": "Charlie"
  }
]`

    if string(result) != expected {
        t.Errorf("unexpected result: got %s, want %s", string(result), expected)
    }
}
