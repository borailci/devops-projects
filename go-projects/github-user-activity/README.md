# GitHub User Activity CLI

A simple command-line interface to fetch and display recent GitHub user activity.

## Features

- Fetch recent activity for any GitHub user
- Display activity in a human-readable format
- Handle various GitHub event types (pushes, issues, stars, forks, etc.)
- Graceful error handling for invalid usernames and API failures
- No external dependencies - uses only Go standard library

## Installation

1. Make sure you have Go installed (version 1.21 or later)
2. Clone or download this project
3. Navigate to the project directory
4. Build the application:

```bash
go build -o github-activity
```

## Usage

Run the application with a GitHub username:

```bash
./github-activity <username>
```

### Examples

```bash
# Fetch activity for a specific user
./github-activity kamranahmedse

# Example output:
Recent activity for kamranahmedse:

- Pushed 3 commits to kamranahmedse/developer-roadmap
- Opened a new issue in kamranahmedse/developer-roadmap
- Starred kamranahmedse/developer-roadmap
- Forked microsoft/vscode
- Created repository kamranahmedse/new-project
```

## Supported Activity Types

The CLI can display the following types of GitHub activities:

- **Push Events**: Shows number of commits pushed to repositories
- **Issue Events**: Shows when issues are opened, closed, or updated
- **Watch Events**: Shows when repositories are starred
- **Fork Events**: Shows when repositories are forked
- **Create Events**: Shows when new repositories are created
- **Pull Request Events**: Shows when pull requests are opened, closed, or updated
- **Delete Events**: Shows when branches or tags are deleted
- **Public Events**: Shows when repositories are made public
- **Member Events**: Shows when collaborators are added
- **Issue Comment Events**: Shows when comments are made on issues
- **Pull Request Review Events**: Shows when pull requests are reviewed
- **Release Events**: Shows when releases are created

## Error Handling

The application handles various error scenarios:

- **Invalid username**: Returns an error if the username doesn't exist
- **Network issues**: Handles connection timeouts and network failures
- **API rate limits**: Provides informative error messages for rate limiting
- **Empty username**: Validates that a username is provided
- **No activity**: Handles cases where users have no recent activity

## API Information

This application uses the GitHub Events API:

- Endpoint: `https://api.github.com/users/<username>/events`
- No authentication required for public events
- Rate limited to 60 requests per hour for unauthenticated requests

## Building for Different Platforms

```bash
# For Linux
GOOS=linux GOARCH=amd64 go build -o github-activity-linux

# For Windows
GOOS=windows GOARCH=amd64 go build -o github-activity.exe

# For macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o github-activity-macos-intel

# For macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o github-activity-macos-arm64
```

## Development

To run the application without building:

```bash
go run main.go <username>
```

To run tests (if you add them):

```bash
go test
```
