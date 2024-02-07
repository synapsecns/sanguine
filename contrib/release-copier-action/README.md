# Release Copier

[![Go Reference](https://pkg.go.dev/badge/github.com/synapsecns/sanguine/contrib/release-copier-action.svg)](https://pkg.go.dev/github.com/synapsecns/sanguine/contrib/release-copier-action)
[![Go Report Card](https://goreportcard.com/badge/github.com/synapsecns/sanguine/contrib/release-copier-action)](https://goreportcard.com/report/github.com/synapsecns/sanguine/contrib/release-copier-action)

This is a tool to help with the release process. It copies the release from a tag to a new repository. This is used for terraform releases since terraform requires a separate repository for each provider in a specific format. It should be able to be used for any other release you want to copy.

## Usage


```yaml
      - name: Bump version and push tag
        id: tag_version
        if: steps.branch-name.outputs.is_default == 'true'
        uses: mathieudutour/github-tag-action@v6.0
        with:
          github_token: ${{ secrets.GITHUB_TOKEN }}
          tag_prefix: my-package-prefix/v
          release_branches: master
          fetch_all_tags: true
      - name: Copy Releases
        uses: docker://ghcr.io/synapsecns/sanguine/release-copier-action:latest
        with:
            # this token must have access to both the original repository and the new repository so GITHUB_TOKEN will not work
            github_token: ${{ secrets.PUBLISH_TOKEN }}
            # destination repo
            destination_repo: 'my-destination-repo/destination-package'
            # you can take this from anywhere
            tag_name: ${{ steps.tag_version.outputs.new_tag }}
            # we strip away anything package relative here
            strip_prefix: 'my-package-prefix/'
```

## A note on actions

This action is currently not published to the marketplace, partially because the [requirements](https://docs.github.com/en/actions/creating-actions/publishing-actions-in-github-marketplace#about-publishing-actions) require that each repository contain a single action and the action.yml must be in the root directory. We can get around this with a subdirectory copier in a future version.
