#!/usr/bin/env bash

dirs=()

OG_DIR=$(pwd)

# list all chart dirs
for filename in *; do
    # go back to the beginning after cding below
    cd "$OG_DIR" || exit 1
    # skip files
    if [[ -f "$filename" ]]; then
      continue
    fi

    # skip symbolic links
    if [[ -L "$filename" ]]; then
      continue
    fi

    cd "$filename" || exit 1

    # auto add https://github.com/helm/helm/issues/8036#issuecomment-1126959239
    if [ ! -f "./Chart.yaml" ]; then
      continue
    fi

    dirs+=("$filename")
done

for I in "${dirs[@]}"
do
    OUT=$I,${OUT:+$OUT }
done

cutVal=$((${#dirs[@]} + 1))
echo "$OUT"  | rev | cut -c${cutVal}- | rev
