// internal/api/api_test.go
package api

import (
    "context"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    dialogflow "cloud.google.com/go/dialogflow/apiv2"
    dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
    "google.golang.org/api/option"
    "github.com/matt-dunleavy/json-transform/config"
)

func TestProcessJSONWithGoogleAPI(t *testing.T) {
    ctx := context.Background()
    mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        response := dialogflowpb.DetectIntentResponse{
            QueryResult: &dialogflowpb.QueryResult{
                FulfillmentText: "google processed result",
            },
        }
        json.NewEncoder(w).Encode(response)
    }))
    defer mockServer.Close()

    client, err := dialogflow.NewSessionsClient(ctx, option.WithEndpoint(mockServer.URL), option.WithAPIKey("test-key"), option.WithoutAuthentication())
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }

    sessionPath := fmt.Sprintf("projects/%s/agent/sessions/%s", "test-project-id", "some-unique-session-id")

    request := &dialogflowpb.DetectIntentRequest{
        Session: sessionPath,
        QueryInput: &dialogflowpb.QueryInput{
            Input: &dialogflowpb.QueryInput_Text{
                Text: &dialogflowpb.TextInput{
                    Text:         "Test input",
                    LanguageCode: "en-US",
                },
            },
        },
    }

    response, err := client.DetectIntent(ctx, request)
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }

    if response.GetQueryResult().GetFulfillmentText() != "google processed result" {
        t.Fatalf("expected 'google processed result', got %v", response.GetQueryResult().GetFulfillmentText())
    }
}
