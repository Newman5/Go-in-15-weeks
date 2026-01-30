#!/bin/bash
# Script Name: edit-draft.sh
# Purpose: Open a draft for editing
# Usage: ./scripts/edit-draft.sh <slug>
#
# This script will:
# 1. Find the matching draft in content/drafts/
# 2. Open it in your editor

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
# DEFINE THE DRAFTS DIRECTORY
# ============================================
drafts_dir="content/drafts"

if [ ! -d "$drafts_dir" ]; then
  echo "Error: Drafts directory not found: $drafts_dir"
  exit 1
fi

# ============================================
# FIND THE DRAFT FILE
# ============================================
# Try exact matches first
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
# OPEN IN EDITOR
# ============================================
if [ -n "$EDITOR" ]; then
  echo "Opening in $EDITOR..."
  $EDITOR "$draft_file"
else
  echo "Opening in code (set \$EDITOR to use a different editor)..."
  code "$draft_file"
fi
