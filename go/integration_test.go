package main

import (
  "io/ioutil"
  "net/http"
  "os"
  "testing"
)

func TestAPIIntegration(t *testing.T) {
  apiAddr := os.Getenv("API_ADDR")
  if apiAddr == "" {
    apiAddr = "http://127.0.0.1:8080"
  }
  url := apiAddr + "/api/v1/resource"  // adjust to a real path

  resp, err := http.Get(url)
  if err != nil {
    t.Fatalf("GET %s failed: %v", url, err)
  }
  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
    t.Fatalf("expected 200 OK, got %d", resp.StatusCode)
  }

  body, _ := ioutil.ReadAll(resp.Body)
  if len(body) == 0 {
    t.Error("empty response body")
  }
}
