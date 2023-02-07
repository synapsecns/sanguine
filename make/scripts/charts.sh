#!/usr/bin/env bash

OG_DIR=$(pwd)

if [ $1 == "" ]; then
  echo "No chart name provided"
  exit 1
fi

# assumes that we're in the charts dir
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
    if [ -f "./Chart.lock" ]; then
      if [ $filename == "agents" ]; then \
      for i in "embedded" "remote-fresh" "remote-existing"; do \
          cd $1; \
          ct install --debug --helm-extra-set-args "--set=executor.type=$i" --chart-dirs agents --charts agents; \
      done; \
      fi;
      	if [ $filename != "" ] && [ $filename != "agents" ]; then cd $1; ct install --debug --chart-dirs $filename --charts $filename; fi;
    fi

#    helm dependency update
done
