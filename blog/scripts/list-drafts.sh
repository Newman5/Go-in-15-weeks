#!/bin/bash
# Script Name: list-drafts.sh
# Purpose: Lists all draft posts with titles and word counts
# Usage: ./scripts/list-drafts.sh
#
# This script will:
# 1. Find all .md files in content/drafts/
# 2. Extract the title from each file's front matter
# 3. Count words in each draft
# 4. Display them in a readable format

# ============================================
# DEFINE THE DRAFTS DIRECTORY
# ============================================
drafts_dir="content/drafts"

# ============================================
# CHECK IF DRAFTS DIRECTORY EXISTS
# ============================================
# The -d test checks if a directory exists
if [ ! -d "$drafts_dir" ]; then
  echo "Error: Drafts directory not found: $drafts_dir"
  exit 1
fi

# ============================================
# FIND ALL MARKDOWN FILES IN DRAFTS
# ============================================
# find command searches for files
# "$drafts_dir" - where to look
# -name "*.md" - only files ending in .md
# -type f - only files (not directories)
drafts=$(find "$drafts_dir" -name "*.md" -type f 2>/dev/null)

# ============================================
# CHECK IF ANY DRAFTS EXIST
# ============================================
# The -z test checks if a variable is empty
if [ -z "$drafts" ]; then
  echo ""
  echo "Draft Posts"
  echo "==========="
  echo ""
  echo "No drafts found in $drafts_dir"
  echo ""
  echo "Create a new draft by saving a post in content/drafts/"
  echo "Or create one with: ./scripts/new-post.sh \"My Draft Title\" --draft"
  exit 0
fi

# ============================================
# DISPLAY HEADER
# ============================================
echo ""
echo "Draft Posts"
echo "==========="
echo ""

# ============================================
# COUNTER FOR TOTAL DRAFTS
# ============================================
count=0

# ============================================
# LOOP THROUGH EACH DRAFT FILE
# ============================================
# for loops iterate over items
# $drafts contains the list of files
# while read -r file reads one file path at a time
echo "$drafts" | while read -r file; do
  # Skip if file doesn't exist (safety check)
  if [ ! -f "$file" ]; then
    continue
  fi
  
  # Increment counter
  count=$((count + 1))
  
  # ============================================
  # EXTRACT THE TITLE
  # ============================================
  # grep finds lines matching a pattern
  # "^title:" means lines starting with "title:"
  # head -n 1 takes only the first match
  # sed removes the "title: " part and quotes
  title=$(grep "^title:" "$file" | head -n 1 | sed 's/^title: *"\?\(.*\)"\?$/\1/')
  
  # If no title found, use filename
  if [ -z "$title" ]; then
    title=$(basename "$file" .md)
  fi
  
  # ============================================
  # COUNT WORDS IN THE FILE
  # ============================================
  # wc -w counts words
  # We exclude the front matter (between --- markers)
  # awk is used to skip lines between the first two ---
  word_count=$(awk '
    BEGIN { in_frontmatter=0; front_matter_count=0 }
    /^---$/ { 
      in_frontmatter = !in_frontmatter
      front_matter_count++
      next 
    }
    front_matter_count >= 2 && !in_frontmatter { print }
  ' "$file" | wc -w)
  
  # ============================================
  # EXTRACT THE FILENAME (WITHOUT PATH)
  # ============================================
  # basename removes the directory path
  filename=$(basename "$file")
  
  # ============================================
  # DISPLAY THE DRAFT INFO
  # ============================================
  echo "• $title"
  echo "  File: $filename"
  echo "  Words: $word_count"
  echo ""
done

# ============================================
# DISPLAY TOTAL COUNT
# ============================================
# Count the number of lines (files) in drafts
total=$(echo "$drafts" | wc -l)
echo "---"
echo "Total drafts: $total"
echo ""
