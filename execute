#!/usr/bin/env bash
set -euo pipefail

FEATURE="$1"

while [ -n "$(ls .plans/$FEATURE/pending/*.md 2>/dev/null)" ]; do
  echo "Executing next slice..."
  opencode run \
    --model "openai/gpt-5.4" \
    --variant high \
    "execute the next slice"
done

echo "All slices completed for $FEATURE"
