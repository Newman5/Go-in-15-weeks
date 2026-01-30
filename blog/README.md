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
