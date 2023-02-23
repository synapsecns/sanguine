# Git Changes Action

This GitHub Action exports a variable that contains the list of Go modules changed in the current pull request, along with any dependent modules. This can be useful for automating build, test, or deployment workflows that involve Go projects.


## Usage:

1. To use this action, add the following steps to your workflow:

    Check out the current pull request using the actions/checkout action. It's recommended to set fetch-depth to 0 and submodules to recursive to ensure that all necessary dependencies are fetched.


```yaml
    steps:
      - uses: actions/checkout@v3
        with:
          fetch-depth: 0
          submodules: 'recursive'
```

1. Use the synapsecns/sanguine/git-changes-action Docker image to run the git-changes script, which exports a variable that contains the list of changed Go modules.

```yaml
      - uses: docker://ghcr.io/synapsecns/sanguine/git-changes-action:latest
        id: filter_go
        with:
          include_deps: true
          github_token: ${{ secrets.github_token }}
          timeout: "1m" # optional, defaults to 1m
```

You can customize the behavior of the git-changes script by using the following inputs:

 - `include_deps`: A boolean that controls whether dependent modules are included in the list of changed modules. Set to true by default.
 - `github_token`: The token to use for authentication with the GitHub API. This is required to fetch information about the current pull request.
 - `timeout`: The maximum time to wait for the GitHub API to respond. Defaults to 1 minute.

The output of the git-changes script is a comma-separated list of Go module paths. You can access this list using the `filter_go` output variable, like so:

```yaml
      - run: echo "Changed modules: ${{ steps.filter_go.outputs.changed_modules }}"
```

## Example

Here's an example workflow that uses the `git-changes` action to run tests for changed Go modules:


```yaml
name: Test Go Modules

on:
  pull_request:
    types: [opened, synchronize]

jobs:
  test_go_modules:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
        with:
          fetch-depth: 0
          submodules: 'recursive'

      - uses: docker://ghcr.io/synapsecns/sanguine/git-changes-action:latest
        id: filter_go
        with:
          include_deps: true
          github_token: ${{ secrets.github_token }}
          timeout: "1m"

      - name: Run tests
        run: go test -v ${{ steps.filter_go.outputs.changed_modules }}
```

This workflow will run tests for all changed Go modules and their dependencies whenever a pull request is opened or synchronized.



