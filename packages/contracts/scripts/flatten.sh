#!/bin/bash

mkdir flattened/
for i in `find ./contracts -name "*.sol" -type f`; do
    export OUTNAME="./flattened/$(basename $i)"
    forge flatten "$i" > $OUTNAME
done