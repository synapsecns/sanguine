#!/usr/bin/env bash

find . -name go.mod -print0 | while IFS= read -r -d '' f; do
  dir=$(dirname "$f")
  echo "Changing to directory $dir"
  (cd "$dir" || exit; go mod tidy)
done
