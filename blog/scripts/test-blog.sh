#!/bin/bash
# Script Name: test-blog.sh
# Purpose: Run smoke tests for blog scripts and content paths
# Usage: ./scripts/test-blog.sh [--build]

set -euo pipefail

# ============================================
# SETUP PATHS
# ============================================
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
BLOG_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
ROOT_DIR="$(cd "$BLOG_DIR/.." && pwd)"

ROOT_CONTENT="$ROOT_DIR/content"
BLOG_CONTENT="$BLOG_DIR/content"

run_build=false

for arg in "$@"; do
  case "$arg" in
    --build)
      run_build=true
      ;;
    --help|-h)
      echo "Usage: $0 [--build]"
      exit 0
      ;;
  esac
done

# ============================================
# HELPERS
# ============================================
info() { echo -e "\033[0;36mℹ $1\033[0m"; }
pass() { echo -e "\033[0;32m✓ $1\033[0m"; }
warn() { echo -e "\033[1;33m⚠ $1\033[0m"; }
fail() { echo -e "\033[0;31m✗ $1\033[0m"; exit 1; }

# ============================================
# BASIC CHECKS
# ============================================
info "Checking required scripts..."
required_scripts=(
  "new-post.sh"
  "new-link-post.sh"
  "list-drafts.sh"
  "edit-draft.sh"
  "publish-draft.sh"
  "find-post.sh"
  "check-links.sh"
  "post-stats.sh"
  "generate-summary.sh"
)

for script in "${required_scripts[@]}"; do
  if [ ! -f "$SCRIPT_DIR/$script" ]; then
    fail "Missing script: $SCRIPT_DIR/$script"
  fi
  if [ ! -x "$SCRIPT_DIR/$script" ]; then
    warn "Script not executable: $SCRIPT_DIR/$script"
  fi
done
pass "All required scripts present."

info "Checking content directories..."
if [ -d "$ROOT_CONTENT" ] && [ -d "$BLOG_CONTENT" ]; then
  root_inode=$(stat -c %i "$ROOT_CONTENT")
  blog_inode=$(stat -c %i "$BLOG_CONTENT")
  if [ "$root_inode" != "$blog_inode" ]; then
    warn "Both root and blog content directories exist and are different."
    warn "Root: $ROOT_CONTENT"
    warn "Blog: $BLOG_CONTENT"
    warn "Make sure your scripts and 11ty are using the same content folder."
  else
    pass "Content directory is shared between root and blog."
  fi
elif [ -d "$BLOG_CONTENT" ]; then
  pass "Found blog content directory: $BLOG_CONTENT"
elif [ -d "$ROOT_CONTENT" ]; then
  pass "Found root content directory: $ROOT_CONTENT"
else
  fail "No content directory found."
fi

# ============================================
# DRAFT CREATION & PUBLISH TEST
# ============================================
info "Running draft creation test..."
export EDITOR=true

cd "$ROOT_DIR"

stamp=$(date +%Y%m%d%H%M%S)
title="Smoke Test Draft $stamp"

"$SCRIPT_DIR/new-post.sh" "$title" --draft

slug=$(echo "$title" | tr '[:upper:]' '[:lower:]' | tr ' ' '-' | sed 's/[^a-z0-9-]//g')
draft_file="$ROOT_CONTENT/drafts/$(date +%Y-%m-%d)-${slug}.md"

if [ ! -f "$draft_file" ]; then
  fail "Draft file not created: $draft_file"
fi
pass "Draft created."

info "Testing find-post.sh..."
"$SCRIPT_DIR/find-post.sh" "Smoke Test Draft" >/dev/null
pass "find-post.sh works."

info "Publishing draft..."
"$SCRIPT_DIR/publish-draft.sh" "$slug"

post_file="$ROOT_CONTENT/learning-log/$(date +%Y-%m-%d)-${slug}.md"
if [ ! -f "$post_file" ]; then
  fail "Draft was not published to: $post_file"
fi
pass "Draft published."

info "Cleaning up test post..."
rm -f "$post_file"
pass "Cleanup complete."

# ============================================
# OPTIONAL BUILD TEST
# ============================================
if [ "$run_build" = true ]; then
  info "Running Eleventy build..."
  (cd "$BLOG_DIR" && npm run build)
  pass "Build completed."
fi

pass "Smoke tests complete."
