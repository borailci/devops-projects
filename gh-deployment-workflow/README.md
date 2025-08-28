# GitHub Actions Deployment Workflow

This project demonstrates continuous integration and continuous deployment (CI/CD) using GitHub Actions to deploy a static website to GitHub Pages.

## Overview

This repository contains a simple static website that is automatically deployed to GitHub Pages whenever changes are made to the `index.html` file. The deployment is handled by a GitHub Actions workflow that triggers on pushes to the main branch.

## Project Structure

```
gh-deployment-workflow/
├── .github/
│   └── workflows/
│       └── deploy.yml          # GitHub Actions workflow file
├── index.html                  # Main website file
└── README.md                   # This file
```

## Features

- **Automated Deployment**: Automatically deploys changes when `index.html` is modified
- **Conditional Triggers**: Only deploys when the HTML file actually changes
- **GitHub Pages Integration**: Serves the website directly from GitHub Pages
- **Modern Workflow**: Uses the latest GitHub Actions features and best practices

## How It Works

1. **Trigger**: The workflow is triggered when there's a push to the `main` branch that includes changes to `index.html`
2. **Build**: The workflow checks out the code and prepares it for deployment
3. **Deploy**: The website is deployed to GitHub Pages using the official GitHub Pages action
4. **Access**: The website becomes available at `https://<username>.github.io/gh-deployment-workflow/`

## GitHub Actions Workflow

The `.github/workflows/deploy.yml` file contains the workflow configuration that:

- Monitors pushes to the main branch
- Uses path filtering to only trigger when `index.html` changes
- Checks out the repository code
- Configures GitHub Pages
- Uploads and deploys the website

## Setup Instructions

1. **Create Repository**: Create a new GitHub repository named `gh-deployment-workflow`
2. **Enable Pages**: Go to repository Settings > Pages > Source > GitHub Actions
3. **Push Code**: Push this code to your repository
4. **Make Changes**: Edit `index.html` and push to trigger deployment
5. **View Website**: Access your site at `https://<username>.github.io/gh-deployment-workflow/`

## Concepts Demonstrated

- **GitHub Actions**: Automated workflows for CI/CD
- **GitHub Pages**: Static website hosting
- **Conditional Workflows**: Path-based triggers
- **Modern Deployment**: Using official GitHub Actions
- **Version Control Integration**: Git-based deployment pipeline

## Making Changes

To see the workflow in action:

1. Edit the `index.html` file
2. Commit and push your changes to the `main` branch
3. Watch the Actions tab to see the workflow run
4. Check your GitHub Pages URL to see the updated website

## Workflow Features

- **Path Filtering**: Only runs when `index.html` changes
- **Efficient**: Skips unnecessary deployments
- **Reliable**: Uses official GitHub Actions
- **Secure**: Leverages GitHub's built-in GITHUB_TOKEN

## Next Steps (Stretch Goals)

Consider enhancing this project by:

- Adding a static site generator (Hugo, Jekyll, Astro)
- Including CSS/JS build steps
- Adding automated testing
- Implementing staging environments
- Adding custom domain support

## License

This project is open source and available under the MIT License.
