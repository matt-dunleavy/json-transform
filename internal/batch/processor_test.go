// internal/batch/processor_test.go
package batch

import (
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "os"
    "path/filepath"
    "testing"

    "github.com/matt-dunleavy/json-transform/internal/api"
)

func TestProcessBatchFiles(t *testing.T) {
    tempDir := t.TempDir()
    inputDir := filepath.Join(tempDir, "input")
    outputDir := filepath.Join(tempDir, "output")
    promptFile := filepath.Join(tempDir, "prompt.txt")
    apiKey := "test-key"
    apiService := "chatgpt"
    model := "davinci"
    operation := "api-process"

    os.Mkdir(inputDir, 0755)
    os.Mkdir(outputDir, 0755)

    inputFilePath := filepath.Join(inputDir, "input.json")
    ioutil.WriteFile(inputFilePath, []byte(`{"data":"test"}`), 0644)
    ioutil.WriteFile(promptFile, []byte("Test prompt"), 0644)

    testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        response := api.APIResponse{
            Result: "processed result",
        }
        json.NewEncoder(w).Encode(response)
    }))
    defer testServer.Close()

    originalURL := "https://api.openai.com/v1/engines/davinci/completions"
    defer func() {
        http.DefaultClient = &http.Client{}
    }()

    apiURL := testServer.URL

    err := ProcessBatchFiles(inputDir, outputDir, apiKey, promptFile, apiService, model, operation)
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }

    outputFilePath := filepath.Join(outputDir, "input.json")
    result, err := ioutil.ReadFile(outputFilePath)
    if err != nil {
        t.Fatalf("failed to read output file: %v", err)
    }

    expectedResult := "processed result"
    if string(result) != expectedResult {
        t.Fatalf("unexpected result: got %s, want %s", result, expectedResult)
    }
}
