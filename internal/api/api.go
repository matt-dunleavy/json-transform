// internal/api/api.go
package api

import (
    "bytes"
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
    "net/http"
)

type APIRequest struct {
    Contents []Content `json:"contents"`
}

type Content struct {
    Parts []Part `json:"parts"`
}

type Part struct {
    Text string `json:"text"`
}

type APIResponse struct {
    Candidates []Candidate `json:"candidates"`
}

type Candidate struct {
    Content Content `json:"content"`
}

func ProcessJSONWithAPI(input, prompt, apiKey, apiService, model string) (string, error) {
    switch apiService {
    case "gemini":
        return processWithGemini(input, prompt, apiKey, model)
    case "chatgpt":
        return processWithChatGPT(input, prompt, apiKey, model)
    default:
        return "", errors.New("unsupported API service")
    }
}

func processWithGemini(input, prompt, apiKey, model string) (string, error) {
    requestBody, err := json.Marshal(APIRequest{
        Contents: []Content{
            {
                Parts: []Part{
                    {Text: input},
                },
            },
        },
    })
    if err != nil {
        return "", err
    }

    apiURL := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s", model, apiKey)
    req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
    if err != nil {
        return "", err
    }

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        body, _ := ioutil.ReadAll(resp.Body)
        return "", fmt.Errorf("API request failed with status: %s, body: %s", resp.Status, string(body))
    }

    var apiResponse APIResponse
    err = json.NewDecoder(resp.Body).Decode(&apiResponse)
    if err != nil {
        return "", err
    }

    if len(apiResponse.Candidates) == 0 {
        return "", errors.New("no candidates in response")
    }

    return apiResponse.Candidates[0].Content.Parts[0].Text, nil
}

func processWithChatGPT(input, prompt, apiKey, model string) (string, error) {
    requestBody, err := json.Marshal(map[string]interface{}{
        "prompt": prompt,
        "input":  input,
        "model":  model,
    })
    if err != nil {
        return "", err
    }

    apiURL := fmt.Sprintf("https://api.openai.com/v1/engines/%s/completions", model)

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
        body, _ := ioutil.ReadAll(resp.Body)
        return "", fmt.Errorf("API request failed with status: %s, body: %s", resp.Status, string(body))
    }

    var apiResponse APIResponse
    err = json.NewDecoder(resp.Body).Decode(&apiResponse)
    if err != nil {
        return "", err
    }

    return apiResponse.Candidates[0].Content.Parts[0].Text, nil
}
