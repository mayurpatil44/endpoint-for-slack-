package main

import (
  "encoding/json"
  "net/http"
  "time"
)

func handler(w http.ResponseWriter, r *http.Request) {

  slackName := r.URL.Query().Get("slack_name")
  track := r.URL.Query().Get("track")

  currentDay := time.Now().Weekday().String()

  utcTime := time.Now().UTC().Format(time.RFC3339)

  githubFileURL := "https://github.com/username/repo/blob/main/main.go"

  githubRepoURL := "https://github.com/username/repo"

  response := map[string]interface{}{
    "slack_name": slackName,
    "current_day": currentDay,
    "utc_time": utcTime, 
    "track": track,
    "github_file_url": githubFileURL,
    "github_repo_url": githubRepoURL,
    "status_code": 200,
  }

  jsonResponse, _ := json.Marshal(response)

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(200)
  w.Write(jsonResponse)

}

func main() {
  http.HandleFunc("/api", handler)
  http.ListenAndServe(":5000", nil)
}