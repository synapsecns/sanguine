#!/usr/bin/env bash

dirs=()

# list all chart dirs
for filename in *; do
    # skip files
    [ -e "$filename" ] || continue

    # skip symbolic links
    if [[ -L "$filename" ]]; then
      continue
    fi

    cd "$filename" || exit 1

    # auto add https://github.com/helm/helm/issues/8036#issuecomment-1126959239
    if [ ! -f "./Chart.yaml" ]; then
      continue
    fi

    dirs+=("charts/$filename")
done

for I in "${dirs[@]}"
do
    OUT=$I,${OUT:+$OUT }
done

echo "$OUT"  | rev | cut -c2- | rev
