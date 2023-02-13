# Git Changes Action

This action will export a variable containing the list of go modules changed in the current pull request, and any dependent modules.


Usage:

```yaml
    steps:
      - uses: actions/checkout@v3
        with:
          fetch-depth: 0
          submodules: 'recursive'

      - uses: docker://ghcr.io/synapsecns/sanguine/git-changes-action:80b3bd017f46acd960a72cfc0095587d33928c34
        id: filter_go
        with:
          include_deps: true
          github_token: ${{ secrets.github_token }}
```
