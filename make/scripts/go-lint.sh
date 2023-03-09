#!/usr/bin/env bash

find . -name go.mod -print0 | while IFS= read -r -d '' f; do
  (cd "$(dirname "$f")" || exit; make lint)
done
