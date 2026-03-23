#!/usr/bin/env bash
set -euo pipefail

FEATURE="$1"
PLAN_DIR=".plans/$FEATURE"
PENDING="$PLAN_DIR/pending"
COMPLETED="$PLAN_DIR/completed"
SERVER="${OPENCODE_URL:-http://localhost:4096}"
MODEL="${OPENCODE_MODEL:-openai/gpt-5.4}"

shopt -s nullglob
slices=("$PENDING"/*.md)
shopt -u nullglob

if [ ${#slices[@]} -eq 0 ]; then
  echo "No pending slices in $PENDING"
  exit 0
fi

mkdir -p "$COMPLETED"

for slice in "${slices[@]}"; do
  name=$(basename "$slice")
  echo "=== $name ==="

  opencode run --attach "$SERVER" -m "$MODEL" --variant high \
    --file "$PLAN_DIR/plan.md" \
    --file "$slice" \
    "Implement the slice in $slice. Read plan.md for context. Read completed slices in $COMPLETED/ if there are dependencies. Update the slice file with a Completion section when done. Then use the commit skill to commit." \
    > /dev/null

  mv "$slice" "$COMPLETED/$name"
done

echo "Done: $FEATURE"
