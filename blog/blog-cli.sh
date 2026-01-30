#!/bin/bash
# Script Name: blog-cli.sh
# Purpose: Command-line interface for managing learning log posts
# Usage: ./blog-cli.sh [command] [arguments]
#
# This is the main entry point for blog commands.
# Phase 1: Supports new post creation and listing drafts

# ============================================
# HELPER FUNCTIONS
# ============================================

show_help() {
  echo ""
  echo "Go Learning Blog CLI - Phase 1"
  echo "=============================="
  echo ""
  echo "Usage: ./blog-cli.sh [command] [arguments]"
  echo ""
  echo "Available Commands:"
  echo "  new <title>      Create a new learning log post"
  echo "  drafts           List all draft posts"
  echo "  help             Show this help message"
  echo ""
  echo "Examples:"
  echo "  ./blog-cli.sh new \"Day 3: Functions and Errors\""
  echo "  ./blog-cli.sh drafts"
  echo ""
  echo "Tips:"
  echo "  - Posts are created in content/learning-log/"
  echo "  - Drafts are stored in content/drafts/"
  echo "  - Set your editor: export EDITOR=nano"
  echo ""
}

# ============================================
# MAIN EXECUTION
# ============================================

command="$1"

# If no command provided, show help
if [ -z "$command" ]; then
  show_help
  exit 0
fi

# Route commands to appropriate scripts
case "$command" in
  new)
    # Check if title was provided
    if [ -z "$2" ]; then
      echo "Error: Please provide a post title"
      echo "Usage: ./blog-cli.sh new \"Post Title\""
      exit 1
    fi
    ./scripts/new-post.sh "$2"
    ;;
    
  drafts)
    ./scripts/list-drafts.sh
    ;;
    
  --help|-h|help)
    show_help
    ;;
    
  *)
    echo "Error: Unknown command '$command'"
    echo ""
    show_help
    exit 1
    ;;
esac
