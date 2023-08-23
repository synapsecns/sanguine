# Sanguine

[![Go Workflows](https://github.com/synapsecns/sanguine/actions/workflows/go.yml/badge.svg)](https://github.com/synapsecns/sanguine/actions/workflows/go.yml)
[![Foundry Tests](https://github.com/synapsecns/sanguine/actions/workflows/foundry-tests.yml/badge.svg)](https://github.com/synapsecns/sanguine/actions/workflows/foundry-tests.yml)

## Contributing

Read through [CONTRIBUTING.md](./CONTRIBUTING.md) for a general overview of our contribution process.
Then check out our list of [good first issues](https://github.com/synapsecns/sanguine/contribute) to find something fun to work on!

## Directory Structure

<pre>
root
├── <a href="./agents">agents</a>: agents contain all the agents used in optimistic messaging
├── <a href="./charts">charts</a>: The helm charts used for deploying sanguine related services
├── <a href="./contrib">contrib</a>: Devops related tools
│   ├── <a href="./contrib/devnet">devnet</a>: CLI for running a local devnet
│   ├── <a href="./contrib/git-changes-action">git-changes-action</a>: Github action for identifying changes in dependent modules in a go workspace
│   ├── <a href="./contrib/promexporter">promexporter</a>: Multi-service prometheus exporter
│   ├── <a href="./contrib/release-copier-action">release-copier-action</a>: Github action for copying releases from one repo to another
│   ├── <a href="./contrib/terraform-provider-iap">terraform-provider-iap</a>: Terraform provider used for bastion proxy tunneling
│   ├── <a href="./contrib/terraform-provider-helmproxy">terraform-provider-helmproxy</a>: Terraform provider that allows helm to be proxied through an iap bastion proxy
│   ├── <a href="./contrib/terraform-provider-kubeproxy">terraform-provider-kubeproxy</a>: Terraform provider that allows kube to be proxied through an iap bastion proxy
│   ├── <a href="./contrib/tfcore">tfcore</a>: Terraform core utilities + iap utilities
├── <a href="./core">core</a>: The Go core library with common utilities for use across the monorepo
├── <a href="./ethergo">ethergo</a>: Go-based ethereum testing + common library
├── <a href="./packages">packages</a>
│   ├── <a href="./packages/contracts-core">contracts-core</a>: Core contracts used for synapse, powered by <a href="https://github.com/foundry-rs/foundry">Foundry</a>
│   ├── <a href="./packages/coverage-aggregator">coverage-aggregator</a>: Javascript coverage aggregator based on <a href="https://www.npmjs.com/package/nyc">nyc</a>
│   ├── <a href="./packages/docs">docs</a>: Docasaurus documentation
│   ├── <a href="./packages/explorer-ui">explorer-ui</a>: Explorer UI
│   ├── <a href="./packages/sdk-router">sdk-router</a>: SDK router
│   ├── <a href="./packages/sdk-router">synapse-interface</a>: Synapse frontend code
├── <a href="./tools">services</a>
│   ├── <a href="./services/explorer">explorer</a>: Bridge/messaging explorer backend
│   ├── <a href="./services/omnirpc">omnirpc</a>: Latency aware RPC Client used across multiple-chains at once
│   ├── <a href="./services/scribe">scribe</a>: Generalized ethereum event logger
├── <a href="./tools">tools</a>
│   ├── <a href="./tools/abigen">abigen</a>: Used to generate abigen bindings for go
│   ├── <a href="./tools/bundle">bundle</a>: Modified version of <a href="https://pkg.go.dev/golang.org/x/tools@v0.5.0/cmd/bundle"> go bundler </a> with improved shadowing support
│   ├── <a href="./tools/modulecopier">module copier</a>: Used to copy internal modules and export methods for testing
</pre>

## Setup

Clone the repository, open it, and install nodejs packages with `yarn`:

```bash
git clone https://github.com/synapsecns/sanguine --recurse-submodules -j10
cd sanguine
yarn install
```


### Install the Correct Version of NodeJS

Using `nvm`, install the correct version of NodeJS.

```
nvm use
```

### Building the TypeScript packages

To build all of the [TypeScript packages](./packages), run:

```bash
yarn clean
yarn build
```

Packages compiled when on one branch may not be compatible with packages on a different branch.
**You should recompile all packages whenever you move from one branch to another.**
Use the above commands to recompile the packages.

## Dealing with submodules

This repo make use of [multiple](.gitattributes) [submodules](https://git-scm.com/book/en/v2/Git-Tools-Submodules). To avoid issues when checking out different branches, you can use `git submodule update --init --recursive` after switching to a branch or `git checkout feat/branch-name --recurse-submodules` when switching branches.

# Building Agents Locally

<!-- TODO: we need to move this thing into an ops docs package. Given that the docs are still a work in progress, I'm leaving this here for now. -->
<!-- Actually, it's unclear if this belongs in a contributing.md file, the docs or both. Maybe a symlink? -->

In order to minimize risks coming from extraneous dependencies or supply chain attacks in a production like enviornment, all distributed images are built as [scratch](https://hub.docker.com/_/scratch) or [distroless](https://github.com/GoogleContainerTools/distroless#distroless-container-images) images. Builder containers are also not used to restrict the build enviornment to the [goreleaser container](https://github.com/synapsecns/sanguine/pkgs/container/sanguine-goreleaser). All production images are kept in the `docker/` file as `[dir].Dockerfile`. Local

# Adding a new Go Module

If you need to make a new go module, here are the steps to follow:

1. Create a new directory in the right sub-directory. If the fiel tree already has a `go.mod` or a `go.sum` file, you don't need a new module, you're just creating a package. *Note: the `packages` directory is for javascript and should not be used.
2. Create a `go.mod` file in the new directory. You'll want the module name to match the directory path and the package name to be part of go.mod file. The go version should match the version in the root [go.work](go.work) file unless there's a good reaon for it no to.
    ```go
    module github.com/synapsecns/sanguine/path/to/your/module

    go 1.19 // or whatever the version is in go.work
   ```

   Any local dependencies should use replaces like this:
    ```go
    module github.com/synapsecns/sanguine/path/to/your/module

    go 1.19 // or whatever the version is in go.work

   require (
        github.com/synapsecns/sanguine/core v0.0.1
   )

   replace (
	      github.com/synapsecns/sanguine/core => ../path/to/core
    )
    ```

    In so far as you have issues running `go mod tidy`, you may need to add additional replace directives. If you can't figure out what these are, please look at other requires of the module you're trying to link to

3. Add the module to the [go.work](go.work) file. The path should be in alphabetical order.
    ```go
    go 1.19
   use (
    ./a_module
    ./another_module
    ./path/to/your/module
   )
    ```
4. Add the module to the [.codecov.yml](.codecov.yml) file under flags in alphabetical order <!--TODO: enforce alphabetical order w/ linter-->. This allows codecov to re-use old coverage information if no changes have been made, which speeds up testing. For an explanation of when changes are ran please see [this post](https://threadreaderapp.com/thread/1693572913662775510.html) <!-- todo: this needs to be moved into ci docs-->, the [go workflow](.github/workflows/go.yml) and the [git-changes-action](contrib/git-changes-action/README.md). For an explanation of the carryforward flag, please see the [codecov docs](https://docs.codecov.com/docs/carryforward-flags):
    ```yaml
      # Lots of other stuff....

      # from the go.work file
      flags:
        # other flags...

        your-module-name: # in the case of github.com/synapsecns/sanguine/path/to/your/module, this would be module
          path: path/to/your/module/
          carryforward: true
    ```

    #### Note: The codecov.yaml used in ci will not be updated until your branch is merged to master. This is expected

5. Create a Makefile.

    If your makefile has no custom commands (it shouldn't if you're just starting), simply create a symlink to the go.Makefile by running `ln -sv ../path/to/repo/root/make/go.Makefile Makefile`.

    Otherwise, create the makefile in `$REPO_ROOT/make/[module_name].Makefile` with the following text:
    ```makefile
    include ../path/to/go.Makefile # this is the path to the go.Makefile from the module directory


    custom command:
      @eval $(echo "do something new!")
    ```

    then symlink it like above.

    Note: please do your best to make `Makefile` commands as portable as possible. For example, this installs tfenv for testing the terraform modules on osx and linux:

    ```makefile
    tfenv-install:
      @#Brew - MacOS
      @if [ "$(shell which tflint)" = "" ] && [ "$(shell which brew)" != "" ]; then brew install rflint; fi;
      # default
      @if [ "$(shell which tflint)" = "" ]; then curl -s https://raw.githubusercontent.com/terraform-linters/tflint/master/install_linux.sh | bash; fi;
    ```

6. Add a `.goreleaser.yml` file. If you're just starting the directory, it's likely you don't have a binary/Dockerfile yet so it can look something like this:
    ```yaml
    project_name: your-module-name

    monorepo:
      tag_prefix: path/from/repo/root/to/module
      dir: path/from/repo/root/to/module

    # for now, this is a library release
    builds:
      - skip: true

    # add a source archive at release time
    source:
      enabled: true

    # Archives
    archives:
      - format: tar.gz
        wrap_in_directory: true
        format_overrides:
          - goos: windows
            format: zip
        name_template: '{{.ProjectName}}-{{.Version}}_{{.Os}}_{{.Arch}}'
        replacements:
          amd64: x86_64
          arm64: ARM64
          darwin: macOS
          linux: Linux
          windows: Windows
        files:
          - README.md

    checksum:
      name_template: checksums.txt

    # Add a changelog
    changelog:
      sort: asc
    ```

    If this is not a library release, please see any other `.goreleaser.yml` file for an example. It is important the production docker files are put in the [`docker/`](docker/) directory and named `[module_name].Dockerfile` so that others can find them.

7. Create a `README.md` file in the directory. This should include a description of the module and a link to the godoc page and go report card <!-- (for vanity). If you're not sure what to put in the readme, look at the other modules for inspiration. If the module is on the more complex side, consider including a directory tree like we have [here](#directory-structure). If the application is runnable, instructions on running should be present.
8. Add the new directory and description to the [Directory Structure](#directory-structure) section of this README in alphabetical order. <!-- TODO: we should really lint that this is in alphabetical order. Also, markdown files should be linted for broken links. -->
