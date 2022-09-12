#!/usr/bin/env bash

# assumes that we're in the charts dir
for filename in *; do
    # skip files
    [ -e "$filename" ] || continue

    # skip symbolic links
    if [[ -L "$filename" ]]; then
      continue
    fi

    cd "$filename" || exit 1

    # auto add https://github.com/helm/helm/issues/8036#issuecomment-1126959239
    if [ -f "./Chart.lock" ]; then
      yq --indent 0 '.dependencies | map(["helm", "repo", "add", .name, .repository] | join(" ")) | .[]' "./Chart.lock"  | sh --;
    fi

    helm dependency update
done
