package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// GitHubEvent represents the structure of a GitHub event from the API
type GitHubEvent struct {
	Type      string    `json:"type"`
	Actor     Actor     `json:"actor"`
	Repo      Repo      `json:"repo"`
	Payload   Payload   `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
	Public    bool      `json:"public"`
}

type Actor struct {
	Login string `json:"login"`
}

type Repo struct {
	Name string `json:"name"`
}

type Payload struct {
	Size        int         `json:"size"`
	Commits     []Commit    `json:"commits"`
	Action      string      `json:"action"`
	Issue       Issue       `json:"issue"`
	PullRequest PullRequest `json:"pull_request"`
	Forkee      Forkee      `json:"forkee"`
}

type Commit struct {
	Message string `json:"message"`
	SHA     string `json:"sha"`
}

type Issue struct {
	Title  string `json:"title"`
	Number int    `json:"number"`
}

type PullRequest struct {
	Title  string `json:"title"`
	Number int    `json:"number"`
}

type Forkee struct {
	FullName string `json:"full_name"`
}

func main() {
	// Check if username argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: github-activity <username>")
		fmt.Println("Example: github-activity kamranahmedse")
		os.Exit(1)
	}

	username := os.Args[1]

	// Validate username (basic validation)
	if strings.TrimSpace(username) == "" {
		fmt.Println("Error: Username cannot be empty")
		os.Exit(1)
	}

	// Fetch GitHub user activity
	events, err := fetchGitHubActivity(username)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if len(events) == 0 {
		fmt.Printf("No recent activity found for user: %s\n", username)
		return
	}

	// Display the activity
	fmt.Printf("Recent activity for %s:\n\n", username)
	displayActivity(events)
}

// fetchGitHubActivity fetches the recent activity for a GitHub user
func fetchGitHubActivity(username string) ([]GitHubEvent, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set User-Agent header (GitHub API requires this)
	req.Header.Set("User-Agent", "github-activity-cli")

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from GitHub API: %w", err)
	}
	defer resp.Body.Close()

	// Handle HTTP errors
	switch resp.StatusCode {
	case 200:
		// Success, continue
	case 404:
		return nil, fmt.Errorf("user '%s' not found", username)
	case 403:
		return nil, fmt.Errorf("API rate limit exceeded or access forbidden")
	default:
		return nil, fmt.Errorf("GitHub API returned status code: %d", resp.StatusCode)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse JSON response
	var events []GitHubEvent
	if err := json.Unmarshal(body, &events); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return events, nil
}

// displayActivity formats and displays the GitHub activity
func displayActivity(events []GitHubEvent) {
	// Limit to first 10 events to keep output manageable
	maxEvents := 10
	if len(events) > maxEvents {
		events = events[:maxEvents]
	}

	for _, event := range events {
		message := formatEvent(event)
		if message != "" {
			fmt.Printf("- %s\n", message)
		}
	}
}

// formatEvent converts a GitHub event into a human-readable message
func formatEvent(event GitHubEvent) string {
	switch event.Type {
	case "PushEvent":
		commitCount := len(event.Payload.Commits)
		if commitCount == 0 {
			commitCount = event.Payload.Size // Fallback to size if commits array is empty
		}
		if commitCount == 1 {
			return fmt.Sprintf("Pushed 1 commit to %s", event.Repo.Name)
		}
		return fmt.Sprintf("Pushed %d commits to %s", commitCount, event.Repo.Name)

	case "IssuesEvent":
		action := event.Payload.Action
		if action == "opened" {
			return fmt.Sprintf("Opened a new issue in %s", event.Repo.Name)
		} else if action == "closed" {
			return fmt.Sprintf("Closed an issue in %s", event.Repo.Name)
		}
		return fmt.Sprintf("Updated an issue in %s", event.Repo.Name)

	case "WatchEvent":
		return fmt.Sprintf("Starred %s", event.Repo.Name)

	case "CreateEvent":
		return fmt.Sprintf("Created repository %s", event.Repo.Name)

	case "ForkEvent":
		return fmt.Sprintf("Forked %s", event.Repo.Name)

	case "PullRequestEvent":
		action := event.Payload.Action
		if action == "opened" {
			return fmt.Sprintf("Opened a new pull request in %s", event.Repo.Name)
		} else if action == "closed" {
			return fmt.Sprintf("Closed a pull request in %s", event.Repo.Name)
		}
		return fmt.Sprintf("Updated a pull request in %s", event.Repo.Name)

	case "DeleteEvent":
		return fmt.Sprintf("Deleted a branch or tag in %s", event.Repo.Name)

	case "PublicEvent":
		return fmt.Sprintf("Made %s public", event.Repo.Name)

	case "MemberEvent":
		return fmt.Sprintf("Added a collaborator to %s", event.Repo.Name)

	case "IssueCommentEvent":
		return fmt.Sprintf("Commented on an issue in %s", event.Repo.Name)

	case "PullRequestReviewEvent":
		return fmt.Sprintf("Reviewed a pull request in %s", event.Repo.Name)

	case "ReleaseEvent":
		return fmt.Sprintf("Created a release in %s", event.Repo.Name)

	default:
		// For unknown event types, return a generic message
		return fmt.Sprintf("Performed %s in %s", strings.ToLower(strings.TrimSuffix(event.Type, "Event")), event.Repo.Name)
	}
}
