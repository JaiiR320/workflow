#!/usr/bin/env bash
set -euo pipefail

FEATURE="$1"
PLAN_DIR=".plans/$FEATURE"
PENDING="$PLAN_DIR/pending"
COMPLETED="$PLAN_DIR/completed"
SERVER="${OPENCODE_URL:-http://localhost:4096}"
MODEL="${OPENCODE_MODEL:-openai/gpt-5.4}"

if [ $# -lt 1 ]; then
  echo "Usage: ralph <feature-name>"
  exit 1
fi

if [ ! -f "$PLAN_DIR/plan.md" ]; then
  echo "Error: No plan found at $PLAN_DIR/plan.md"
  exit 1
fi

if [ ! -d "$PENDING" ] || [ -z "$(ls -A "$PENDING"/*.md 2>/dev/null)" ]; then
  echo "No pending slices in $PENDING"
  exit 0
fi

mkdir -p "$COMPLETED"

for slice in $(ls "$PENDING"/*.md | sort); do
  name=$(basename "$slice")
  echo "=== Working on: $name ==="

  prompt="You are executing a vertical slice from a feature plan.\n\n"
  prompt+="## Plan\n$(cat "$PLAN_DIR/plan.md")\n\n"

  for dep in "$COMPLETED"/*.md; do
    [ -f "$dep" ] && prompt+="## Completed: $(basename "$dep")\n$(cat "$dep")\n\n"
  done

  prompt+="## Current Slice\n$(cat "$slice")\n\n"
  prompt+="Implement this slice end-to-end. Update the current slice file itself before finishing.\n\n"
  prompt+="In the slice file, fill in:\n"
  prompt+="- ## Completion: what you built, decisions made, deviations, and anything the next slice should know\n\n"
  prompt+="Do not move the file. The slice file on disk must be updated.\n\n"
  prompt+="When done, use the commit skill to commit your changes."

  opencode run --attach "$SERVER" -m "$MODEL" --variant high "$prompt" > /dev/null

  mv "$slice" "$COMPLETED/$name"
  echo "=== Completed: $name ==="
done

echo "All slices complete for: $FEATURE"
