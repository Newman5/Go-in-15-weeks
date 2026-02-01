# Go Learning Blog

A simple 11ty-based blog for documenting learning progress through the Go in 15 weeks curriculum.

## Getting Started

### Install Dependencies

```bash
cd blog
npm install
```

### Develop

```bash
npm start
```

The site will be available at `http://localhost:8080`

### Build

```bash
npm run build
```

Output is in `_site/`
## Using the Blog Menu 

```bash
./blog-menu.sh 
```

## Using the Blog CLI (Phase 1)

Create and manage learning log posts with bash scripts:

### Create a New Post

```bash
./blog-cli.sh new "Day 3: Functions and Errors"
```

This creates a dated post in `content/learning-log/` with front matter and opens it in your editor.

### List Drafts

```bash
./blog-cli.sh drafts
```

Shows all posts in `content/drafts/` with word counts.

### Set Your Editor

By default, scripts open files in VS Code. To use a different editor:

```bash
export EDITOR=nano  # or vim, code, etc.
```

Add to `~/.bashrc` or `~/.zshrc` to make it permanent.

## Structure

- `content/` - Blog posts and pages
  - `learning-log/` - Learning log entries
  - `about/` - About page
- `_includes/` - Layouts and components
- `_data/` - Site data (metadata, navigation)
- `_config/` - Configuration modules (filters, etc)
- `css/` - Stylesheets
- `public/` - Static assets

## Writing Posts

Create a new markdown file in `content/learning-log/`:

```markdown
---
title: "Day 1: Fundamentals"
date: 2026-01-30
tags:
  - learning-log
  - day-1
---

Your content here...
```

## Future Enhancements

- [ ] Bash scripts for automated post creation
- [ ] Integration with Day directories
- [ ] Link to exercises and code examples
- [ ] Tag-based filtering
- [ ] Search functionality
