#!/bin/bash
# Script Name: find-post.sh
# Purpose: Search posts and drafts by keyword
# Usage: ./scripts/find-post.sh <keyword>
#
# This script will:
# 1. Search content/learning-log and content/drafts for a keyword
# 2. Print matching files and lines

# ============================================
# GET THE KEYWORD
# ============================================
keyword="$1"

# ============================================
# CHECK IF KEYWORD WAS PROVIDED
# ============================================
if [ -z "$keyword" ]; then
  echo "Error: Please provide a keyword to search"
  echo "Usage: $0 <keyword>"
  echo "Example: $0 goroutines"
  exit 1
fi

# ============================================
# DEFINE SEARCH TARGETS
# ============================================
posts_dir="content/learning-log"
drafts_dir="content/drafts"

# ============================================
# RUN THE SEARCH
# ============================================
# -R = recursive
# -n = show line numbers
# -i = case-insensitive
# --color=always = highlight matches

found=false

echo ""
echo "Search Results"
echo "=============="
echo ""

if [ -d "$posts_dir" ]; then
  posts_matches=$(grep -Rni --color=always "$keyword" "$posts_dir" 2>/dev/null)
  if [ -n "$posts_matches" ]; then
    echo "Posts:"
    echo "$posts_matches"
    echo ""
    found=true
  fi
fi

if [ -d "$drafts_dir" ]; then
  drafts_matches=$(grep -Rni --color=always "$keyword" "$drafts_dir" 2>/dev/null)
  if [ -n "$drafts_matches" ]; then
    echo "Drafts:"
    echo "$drafts_matches"
    echo ""
    found=true
  fi
fi

if [ "$found" = false ]; then
  echo "No matches found for '$keyword'"
fi
