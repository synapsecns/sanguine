#!/bin/bash

# This script assumes that you're running it from the root directory of your go.work workspace.

# Get a list of all the modules in the workspace
MODULES=$(go list -m all)
CURRENT_MODULE=$(go list -m)
CURRENT_MODULE_PATH=$(echo $CURRENT_MODULE | awk '{print $1}')
CURRENT_MODULE_VERSION=$(echo $CURRENT_MODULE | awk '{print $2}')


# Loop over each module
for MODULE in $MODULES
do
  # Check if the module path starts with "github.com/synapsecns/sanguine"
  if ! echo "$MODULE" | grep -q "^github.com/synapsecns/sanguine/"
  then
    # If it doesn't, skip to the next module
    continue
  fi

  # Check if the module is the current module
  if [ "$MODULE" = "$CURRENT_MODULE_PATH" ]
  then
    # If it is, skip to the next module
    continue
  fi

  # Check if the module is listed in the go.mod file
  if ! grep -qF "$MODULE " go.mod
  then
    # If it's not listed, skip to the next module
    continue
  fi

  # Get the latest tag for the module
  go get $MODULE
done
