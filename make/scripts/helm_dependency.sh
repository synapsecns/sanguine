#!/usr/bin/env bash

# assumes that we're in the charts dir
for filename in *; do
    # skip files
    [ -e "$filename" ] || continue

    # skip symbolic links
    if [[ -L "$filename" ]]; then
      continue
    fi

    cd $filename
    helm dependency update
done
