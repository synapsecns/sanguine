## Remove Label

This action removes a specified label to an issue or pull request.

## Inputs
* `issue_number`: The issue number derived from the event payload. Default is `${{ github.event.number }}` then `${{ github.event.issue.number }}`. Action will fail if neither of these is present. If you need to use this action on a push event (or any event not associated directly with an `issue` or `pull_request`, please see [gh-find-current-pr](https://github.com/jwalton/gh-find-current-pr))
* `label`: The label to add to the issue


```yaml
  - name: Remove Label
    uses: ./.github/actions/remove-label
    with:
      label: 'needs-go-generate-${{matrix.package}}'
      issue-number: ${{github.event.number}}
```
