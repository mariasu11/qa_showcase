package main

import (
  "encoding/json"
  "net/http"
  "os"
  "testing"
  "time"
)

type HealthResp struct {
  Status string `json:"status"`
}

func TestAPIHealth(t *testing.T) {
  apiAddr := os.Getenv("API_ADDR")
  if apiAddr == "" {
    apiAddr = "http://127.0.0.1:8080"
  }
  url := apiAddr + "/health"

  client := &http.Client{Timeout: 5 * time.Second}
  resp, err := client.Get(url)
  if err != nil {
    t.Fatalf("failed GET %s: %v", url, err)
  }
  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
    t.Fatalf("expected 200 OK, got %d", resp.StatusCode)
  }

  var hr HealthResp
  if err := json.NewDecoder(resp.Body).Decode(&hr); err != nil {
    t.Fatalf("invalid JSON: %v", err)
  }
  if hr.Status != "OK" {
    t.Errorf("unexpected status: %v", hr.Status)
  }
}
