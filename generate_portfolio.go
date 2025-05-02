// filepath: /Users/williamosborne/dev/portfolio-site/generate_portfolio.go
package main

import (
    "encoding/json"
    "html/template"
    "log"
    "os"
)

type Project struct {
    Title       string   `json:"Title"`
    Description string   `json:"Description"`
    ImageURL    string   `json:"ImageURL,omitempty"` // Optional: Can be used as poster/thumbnail for videos
    VideoURL    string   `json:"VideoURL,omitempty"` // For direct video file links
    YouTubeID   string   `json:"YouTubeID,omitempty"` // For YouTube video IDs
    GithubURL   string   `json:"GithubURL,omitempty"`
    SiteURL     string   `json:"SiteURL,omitempty"`
    Tags        []string `json:"Tags,omitempty"` // Added Tags field
}

type PortfolioData struct {
    Projects []Project
}

func main() {
    jsonFile, err := os.ReadFile("projects.json")
    if err != nil {
        log.Fatalf("Error reading projects.json: %v", err)
    }

    // Unmarshal JSON data into Project structs
    var projects []Project
    err = json.Unmarshal(jsonFile, &projects)
    if err != nil {
        log.Fatalf("Error unmarshalling projects.json: %v", err)
    }

    data := PortfolioData{
        Projects: projects,
    }

    tmpl, err := template.ParseFiles("template.html")
    if err != nil {
        log.Fatalf("Error parsing template: %v", err)
    }

    outputFile, err := os.Create("index.html")
    if err != nil {
        log.Fatalf("Error creating output file: %v", err)
    }
    defer outputFile.Close()

    // fill the template with the project data
    err = tmpl.Execute(outputFile, data)
    if err != nil {
        log.Fatalf("Error executing template: %v", err)
    }

    log.Println("Successfully generated index.html ✅")
}
