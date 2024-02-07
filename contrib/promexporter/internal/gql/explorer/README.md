# Explorer

This package contains a monitoring interface for the explorer api. The gqlgenc works by spinning up a local version of the explorer and introspecting the graphql schema. The mechanism for this happening exists in `contrib`. We do not reuse the model definitions from the explorer package isnce these may not be useful.
