# GitHub Actions Deployment Workflow - Setup Instructions

## Quick Setup Guide

1. **Create GitHub Repository**

   ```bash
   # Navigate to GitHub and create a new repository named 'gh-deployment-workflow'
   # Or use GitHub CLI if you have it installed:
   gh repo create gh-deployment-workflow --public
   ```

2. **Initialize Git and Push Code**

   ```bash
   cd gh-deployment-workflow
   git init
   git add .
   git commit -m "Initial commit: Add GitHub Actions deployment workflow"
   git branch -M main
   git remote add origin https://github.com/<your-username>/gh-deployment-workflow.git
   git push -u origin main
   ```

3. **Enable GitHub Pages**

   - Go to your repository on GitHub
   - Click on **Settings** tab
   - Scroll down to **Pages** section
   - Under **Source**, select **GitHub Actions**
   - Save the settings

4. **Test the Workflow**

   - Edit `index.html` (change the text, colors, etc.)
   - Commit and push the changes:

   ```bash
   git add index.html
   git commit -m "Update website content"
   git push
   ```

5. **Monitor Deployment**
   - Go to the **Actions** tab in your GitHub repository
   - Watch the workflow run
   - Once complete, visit `https://<your-username>.github.io/gh-deployment-workflow/`

## Key Features Explained

### Workflow Triggers

```yaml
on:
  push:
    branches: [main]
    paths:
      - "index.html"
```

- Only triggers when `index.html` is changed
- Saves resources and prevents unnecessary deployments

### Security Permissions

```yaml
permissions:
  contents: read
  pages: write
  id-token: write
```

- Minimal required permissions for security
- Uses GitHub's built-in GITHUB_TOKEN

### Concurrency Control

```yaml
concurrency:
  group: "pages"
  cancel-in-progress: false
```

- Prevents multiple deployments running simultaneously
- Ensures deployment integrity

## Testing Changes

Try making these changes to see the workflow in action:

1. **Change the greeting**:

   ```html
   <h1>Hello, DevOps World!</h1>
   ```

2. **Update the description**:

   ```html
   <p>Deployed automatically with CI/CD best practices!</p>
   ```

3. **Modify the styling**:
   ```css
   background: linear-gradient(135deg, #ff6b6b 0%, #4ecdc4 100%);
   ```

Each change will trigger a new deployment automatically!
