#!/bin/bash
# Script Name: import-day-files.sh
# Purpose: Import .txt files from Day directories as blog posts
# Usage: ./scripts/import-day-files.sh
#
# This script will:
# 1. Scan all Day* directories in the parent folder
# 2. Find all .txt files in each Day directory
# 3. Convert them to markdown blog posts with front matter
# 4. Preserve original files (non-destructive)

# ============================================
# SETUP PATHS
# ============================================
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
BLOG_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
ROOT_DIR="$(cd "$BLOG_DIR/.." && pwd)"
OUTPUT_DIR="$BLOG_DIR/content/learning-log"

# ============================================
# HELPERS
# ============================================
info() { echo -e "\033[0;36mℹ $1\033[0m"; }
success() { echo -e "\033[0;32m✓ $1\033[0m"; }
warn() { echo -e "\033[1;33m⚠ $1\033[0m"; }

# ============================================
# ENSURE OUTPUT DIRECTORY EXISTS
# ============================================
if ! mkdir -p "$OUTPUT_DIR"; then
  echo "Error: Could not create output directory: $OUTPUT_DIR"
  exit 1
fi

# ============================================
# GET TODAY'S DATE
# ============================================
today=$(date +%Y-%m-%d)

# ============================================
# SCAN DAY DIRECTORIES
# ============================================
info "Scanning for Day directories in: $ROOT_DIR"
echo ""

day_count=0
file_count=0

for day_dir in "$ROOT_DIR"/Day*; do
  # Skip if no Day directories exist
  [ -d "$day_dir" ] || continue
  
  day_count=$((day_count + 1))
  
  # Extract day number from directory name (Day 1 -> 1, Day1 -> 1)
  day_name=$(basename "$day_dir")
  day_num=$(echo "$day_name" | sed 's/Day *//')
  
  info "Processing $day_name..."
  
  # Find all .txt files in this Day directory and subdirectories
  found_files=false
  while IFS= read -r -d '' txt_file; do
    found_files=true
    file_count=$((file_count + 1))
    
    # Get filename without extension and path
    filename=$(basename "$txt_file" .txt)
    
    # Create a URL-friendly slug
    slug=$(echo "$filename" | tr '[:upper:]' '[:lower:]' | tr ' ' '-' | sed 's/[^a-z0-9-]//g')
    
    # Build the output filename
    output_file="$OUTPUT_DIR/${today}-day-${day_num}-${slug}.md"
    
    # Check if file already exists
    if [ -e "$output_file" ]; then
      warn "Skipping (already exists): $output_file"
      continue
    fi
    
    # Create title
    title="Day $day_num: $filename"
    
    # Get relative path for display
    rel_path=$(echo "$txt_file" | sed "s|$ROOT_DIR/||")
    
    # Create the markdown post with front matter
    cat > "$output_file" << EOF
---
title: "$title"
date: $today
tags:
  - day-$day_num
  - learning-log
---

<!-- Imported from: $rel_path -->

$(cat "$txt_file")
EOF
    
    success "Created: $(basename "$output_file")"
    
  done < <(find "$day_dir" -type f -name "*.txt" -print0)
  
  if [ "$found_files" = false ]; then
    echo "  No .txt files found in $day_name"
  fi
  
  echo ""
done

# ============================================
# SUMMARY
# ============================================
if [ $day_count -eq 0 ]; then
  warn "No Day directories found in: $ROOT_DIR"
  exit 0
fi

if [ $file_count -eq 0 ]; then
  warn "No .txt files found in Day directories"
  exit 0
fi

echo "───────────────────────────────────────"
success "Import complete!"
echo "  Day directories scanned: $day_count"
echo "  Files imported: $file_count"
echo "  Output directory: $OUTPUT_DIR"
echo ""
info "Run 'npm run build' to regenerate your site"
