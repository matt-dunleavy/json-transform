// internal/api/api.go
package api

import (
    "bytes"
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
    "log"
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

func ProcessJSONWithAPI(input, prompt, apiKey, apiService, model string, logPayload bool) (string, error) {
    switch apiService {
    case "gemini":
        return processWithGemini(input, prompt, apiKey, model, logPayload)
    case "chatgpt":
        return processWithChatGPT(input, prompt, apiKey, model, logPayload)
    default:
        return "", errors.New("unsupported API service")
    }
}

func processWithGemini(input, prompt, apiKey, model string, logPayload bool) (string, error) {
    combinedInput := fmt.Sprintf("%s\n%s", prompt, input)
    requestBody, err := json.Marshal(APIRequest{
        Contents: []Content{
            {
                Parts: []Part{
                    {Text: combinedInput},
                },
            },
        },
    })
    if err != nil {
        return "", err
    }

    if logPayload {
        log.Printf("Sending request to Gemini with payload: %s", requestBody)
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
        log.Printf("Gemini API request failed with status: %s, body: %s", resp.Status, string(body))
        return "", fmt.Errorf("API request failed with status: %s, body: %s", resp.Status, string(body))
    }

    var apiResponse APIResponse
    err = json.NewDecoder(resp.Body).Decode(&apiResponse)
    if err != nil {
        return "", err
    }

    if logPayload {
        log.Printf("Received response from Gemini: %v", apiResponse)
    }

    if len(apiResponse.Candidates) == 0 {
        return "", errors.New("no candidates in response")
    }

    if len(apiResponse.Candidates[0].Content.Parts) == 0 {
        return "", errors.New("no parts in content")
    }

    return apiResponse.Candidates[0].Content.Parts[0].Text, nil
}

func processWithChatGPT(input, prompt, apiKey, model string, logPayload bool) (string, error) {
    combinedInput := fmt.Sprintf("%s\n%s", prompt, input)
    requestBody, err := json.Marshal(map[string]interface{}{
        "prompt": combinedInput,
        "model":  model,
    })
    if err != nil {
        return "", err
    }

    if logPayload {
        log.Printf("Sending request to ChatGPT with payload: %s", requestBody)
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
        log.Printf("ChatGPT API request failed with status: %s, body: %s", resp.Status, string(body))
        return "", fmt.Errorf("API request failed with status: %s, body: %s", resp.Status, string(body))
    }

    var apiResponse APIResponse
    err = json.NewDecoder(resp.Body).Decode(&apiResponse)
    if err != nil {
        return "", err
    }

    if logPayload {
        log.Printf("Received response from ChatGPT: %v", apiResponse)
    }

    if len(apiResponse.Candidates) == 0 {
        return "", errors.New("no candidates in response")
    }

    if len(apiResponse.Candidates[0].Content.Parts) == 0 {
        return "", errors.New("no parts in content")
    }

    return apiResponse.Candidates[0].Content.Parts[0].Text, nil
}
