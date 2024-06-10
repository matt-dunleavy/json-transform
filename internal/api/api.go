// internal/api/api.go
package api

import (
    "bytes"
    "encoding/json"
    "errors"
    "net/http"
)

type APIRequest struct {
    Prompt string `json:"prompt"`
    Input  string `json:"input"`
    Model  string `json:"model,omitempty"`
}

type APIResponse struct {
    Result string `json:"result"`
}

func ProcessJSONWithAPI(input, prompt, apiKey, apiService, model string) (string, error) {
    requestBody, err := json.Marshal(APIRequest{
        Prompt: prompt,
        Input:  input,
        Model:  model,
    })
    if err != nil {
        return "", err
    }

    var apiURL string
    if apiService == "gemini" {
        apiURL = "https://api.gemini.com/v1/process"
    } else {
        apiURL = "https://api.openai.com/v1/engines/" + model + "/completions"
    }

    req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
    if err != nil {
        return "", err
    }

    req.Header.Set("Authorization", "Bearer "+apiKey)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", errors.New("API request failed with status: " + resp.Status)
    }

    var apiResponse APIResponse
    err = json.NewDecoder(resp.Body).Decode(&apiResponse)
    if err != nil {
        return "", err
    }

    return apiResponse.Result, nil
}
