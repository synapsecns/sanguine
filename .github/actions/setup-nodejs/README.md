## Setup NodeJS

This action sets up NodeJS using a specified `.nvmrc` file. It supports using either yarn or npm as package manager.

### Inputs

* `nvmrc_path`: set to a relative path to use a different `.nvmrc` file. Default is `.nvmrc`.
* `cache`: set to &quot;npm&quot; to use npm instead of yarn. Default is &quot;yarn&quot;.
* `install_dependencies`: set to &quot;false&quot; to skip installing dependencies. Default is &quot;true&quot;.
* `cache-path`: set to a relative path to use a different cache path. Default is `yarn.lock`.

### Outputs

* `node-version`: the installed Node.js version, for tools that install their own runtime (such as Slither).

### Example usage

```yaml
with:
- name: Setup NodeJS
  uses: ./.github/actions/setup-nodejs
  with:
    nvmrc_path: '.nvmrc'
    cache: 'npm'
    install_dependencies: 'true'
    cache-path: 'my-custom-path'
```
