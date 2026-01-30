#!/bin/bash
# Script Name: new-post.sh
# Purpose: Creates a new learning log post with front matter
# Usage: ./scripts/new-post.sh "My Learning Post Title"
#
# This script will:
# 1. Take a title as input
# 2. Generate a filename with today's date
# 3. Create a markdown file with front matter
# 4. Open it in your editor

# ============================================
# GET THE TITLE
# ============================================
# $1 is the first argument passed to this script
title="$1"

# ============================================
# CHECK IF TITLE WAS PROVIDED
# ============================================
# The -z test checks if a variable is empty (zero length)
# If title is empty, show an error and exit
if [ -z "$title" ]; then
  echo "Error: Please provide a title"
  echo "Usage: $0 \"Post Title\""
  echo ""
  echo "Example: $0 \"Day 3: Functions and Errors\""
  exit 1
fi

# ============================================
# GENERATE THE DATE
# ============================================
# Get today's date in YYYY-MM-DD format
# The date command formats dates
# +%Y-%m-%d tells it to output: Year-Month-Day
# Example: 2026-01-30
date=$(date +%Y-%m-%d)

# ============================================
# CREATE A URL-FRIENDLY SLUG
# ============================================
# A slug is a URL-friendly version of the title
# "My First Post" becomes "my-first-post"

# Step 1: Convert to lowercase
# tr translates characters
# '[:upper:]' means all uppercase letters
# '[:lower:]' means all lowercase letters
slug=$(echo "$title" | tr '[:upper:]' '[:lower:]')

# Step 2: Replace spaces with hyphens
# tr ' ' '-' replaces spaces with hyphens
slug=$(echo "$slug" | tr ' ' '-')

# Step 3: Remove any characters that aren't letters, numbers, or hyphens
# sed is a stream editor - it modifies text
# 's/[^a-z0-9-]//g' means:
#   s/       - substitute
#   [^...]   - anything NOT in this list
#   a-z0-9-  - letters, numbers, hyphens
#   //       - replace with nothing (delete)
#   g        - globally (all occurrences)
slug=$(echo "$slug" | sed 's/[^a-z0-9-]//g')

# ============================================
# BUILD THE IMG DIRECTORY PATH
# check if directory exists and create it if not
# ============================================
img_dir="content/images/post-img/${date}-${slug}"
if ! mkdir -p "$img_dir"; then
  echo "Error: Could not create image directory $img_dir"
  exit 1
fi
echo "✓ Image directory created: $img_dir"

# ============================================
# BUILD THE FILENAME
# ============================================
# Combine the date and slug to make the filename
# Example: content/learning-log/2026-01-30-day-3-functions.md
filename="content/learning-log/${date}-${slug}.md"

# ============================================
# CHECK IF FILE ALREADY EXISTS
# ============================================
# The -e test checks if a file exists
# If it does, warn the user and exit
if [ -e "$filename" ]; then
  echo "Error: File already exists: $filename"
  echo "Try a different title or delete the existing file"
  exit 1
fi

# ============================================
# CREATE THE FILE WITH FRONT MATTER
# ============================================
# Use a "here document" to write multiple lines to a file
# Everything between 'cat > "$filename" << EOF' and 'EOF' 
# goes into the file
#
# The front matter is YAML between --- markers
# 11ty uses this metadata to build the site

cat > "$filename" << EOF
---
title: "$title"
date: $date
tags:
  - learning-log
layout: layouts/post.njk
---

## What I Learned

<!-- Main learning points from today -->

### Key Concepts

- Concept 1
- Concept 2
- Concept 3

### Code Examples

\`\`\`go
// Add relevant Go code examples here
\`\`\`

## Challenges & Insights

<!-- What was challenging? What did you discover? -->

## Practice & Exercises

<!-- Links to exercises completed today -->

## Next Steps

<!-- What comes next? What will you explore? -->
EOF

# ============================================
# CONFIRM CREATION
# ============================================
echo ""
echo "✓ Created: $filename"
echo ""

# ============================================
# OPEN IN EDITOR
# ============================================
# Open the file in the user's preferred editor
# ${EDITOR:-vim} means:
#   - Use the $EDITOR environment variable if it's set
#   - Otherwise, use vim as the default
#
# Users can set their editor with: export EDITOR=nano
# Or add it to their ~/.bashrc file

# Check if EDITOR is set, otherwise use code (VS Code)
if [ -n "$EDITOR" ]; then
  echo "Opening in $EDITOR..."
  $EDITOR "$filename"
else
  echo "Opening in code (set \$EDITOR to use a different editor)..."
  code "$filename"
fi
