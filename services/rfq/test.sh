#!/usr/bin/env bash

# Load the number of docker configurations
docker_configs=$(yq e '.dockers | length' .goreleaser.yml)

## Check if there are no docker configurations
#if [ "$docker_configs" -eq "0" ]; then
#  echo "No docker images to push"
#  exit 0
#fi

# Iterate through each docker configuration
for (( i=0; i<docker_configs; i++ )); do
  # Extract the first image template
  image_template=$(yq e ".dockers[$i].image_templates[0]" .goreleaser.yml)

  # Extract the base name from the image template
  image_name=$(echo "$image_template" | sed -E 's|^(.*):[^:]+$|\1|')

  # Tag and push the docker image
  docker tag "$image_name:latest" "$image_name:${GITHUB_SHA}"
  docker push "$image_name:${GITHUB_SHA}"
done
