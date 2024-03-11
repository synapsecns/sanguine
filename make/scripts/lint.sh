#!/usr/bin/env bash

set -e
find . -name go.mod -print0 | while IFS= read -r -d '' f; do
  echo "linting $(dirname "$f")"
  (cd "$(dirname "$f")" || exit; go mod tidy)
done
