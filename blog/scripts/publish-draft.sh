#!/bin/bash
# Script Name: publish-draft.sh
# Purpose: Publish a draft (move to content/learning-log and remove draft flag)
# Usage: ./scripts/publish-draft.sh <slug>
#
# This script will:
# 1. Find the matching draft in content/drafts/
# 2. Remove draft: true from front matter
# 3. Move it to content/learning-log/

# ============================================
# GET THE SLUG
# ============================================
slug="$1"

# ============================================
# CHECK IF SLUG WAS PROVIDED
# ============================================
if [ -z "$slug" ]; then
  echo "Error: Please provide a draft slug"
  echo "Usage: $0 <slug>"
  echo "Example: $0 my-draft-title"
  exit 1
fi

# ============================================
# DEFINE DIRECTORIES
# ============================================
drafts_dir="content/drafts"
posts_dir="content/learning-log"

if [ ! -d "$drafts_dir" ]; then
  echo "Error: Drafts directory not found: $drafts_dir"
  exit 1
fi

if ! mkdir -p "$posts_dir"; then
  echo "Error: Could not create posts directory: $posts_dir"
  exit 1
fi

# ============================================
# FIND THE DRAFT FILE
# ============================================
if [ -f "$drafts_dir/$slug" ]; then
  draft_file="$drafts_dir/$slug"
elif [ -f "$drafts_dir/$slug.md" ]; then
  draft_file="$drafts_dir/$slug.md"
else
  matches=$(find "$drafts_dir" -name "*${slug}*.md" -type f)
  match_count=$(echo "$matches" | grep -c .)

  if [ "$match_count" -eq 0 ]; then
    echo "Error: No draft found matching '$slug'"
    exit 1
  elif [ "$match_count" -gt 1 ]; then
    echo "Multiple drafts matched '$slug':"
    echo "$matches"
    echo "Please be more specific."
    exit 1
  else
    draft_file="$matches"
  fi
fi

# ============================================
# PREPARE DESTINATION
# ============================================
filename=$(basename "$draft_file")
post_file="$posts_dir/$filename"

if [ -e "$post_file" ]; then
  echo "Error: A post already exists with this name: $post_file"
  exit 1
fi

# ============================================
# REMOVE draft: true FROM FRONT MATTER
# ============================================
temp_file=$(mktemp)

# Remove any draft: true line in front matter
# We keep all other content unchanged
sed '/^draft: *true$/d' "$draft_file" > "$temp_file"

# ============================================
# MOVE TO POSTS
# ============================================
if ! mv "$temp_file" "$post_file"; then
  echo "Error: Failed to publish draft"
  rm -f "$temp_file"
  exit 1
fi

# Remove the original draft file
rm -f "$draft_file"

# ============================================
# CONFIRM
# ============================================
echo "✓ Published: $post_file"
