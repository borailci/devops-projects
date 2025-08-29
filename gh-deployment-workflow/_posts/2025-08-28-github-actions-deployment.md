---
layout: post
title: "Building Automated Deployment Workflows with GitHub Actions"
date: 2025-08-28 12:00:00 +0000
categories: devops automation
---

In this post, I'll walk through creating an automated deployment workflow using GitHub Actions that deploys changes to GitHub Pages whenever specific files are modified.

## The Challenge

Manual deployments are error-prone and time-consuming. We need an automated system that:

- Only deploys when specific files change
- Ensures consistent deployment process
- Provides visibility into deployment status

## The Solution

Using GitHub Actions, we can create a workflow that monitors file changes and automatically triggers deployments:

```yaml
name: Deploy to GitHub Pages

on:
  push:
    branches:
      - main
    paths:
      - "index.md"
      - "_posts/**"
      - "_config.yml"

jobs:
  deploy:
    runs-on: ubuntu-latest
    permissions:
      pages: write
      id-token: write
    steps:
      - name: Checkout repository
        uses: actions/checkout@v4

      - name: Setup Ruby
        uses: ruby/setup-ruby@v1
        with:
          ruby-version: "3.1"
          bundler-cache: true

      - name: Build with Jekyll
        run: bundle exec jekyll build

      - name: Upload artifact
        uses: actions/upload-pages-artifact@v3
        with:
          path: ./_site

      - name: Deploy to GitHub Pages
        id: deployment
        uses: actions/deploy-pages@v4
```

## Benefits

- **Automated**: No manual intervention required
- **Selective**: Only triggers on relevant file changes
- **Fast**: Quick feedback and deployment
- **Reliable**: Consistent process every time

This approach demonstrates Infrastructure as Code principles applied to deployment workflows.
