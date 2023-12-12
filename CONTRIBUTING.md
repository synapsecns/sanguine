# Contributing

Contributions should roughly follow the [uber style guide](https://github.com/uber-go/guide/blob/master/style.md)

<!-- todo: more-->

# Adding a new Go Module

If you need to make a new go module, here are the steps to follow:

1. Create a new directory in the right sub-directory. If the file tree already has a `go.mod` or a `go.sum` file, you don't need a new module, you're just creating a package. *Note: the `packages` directory is for javascript and should not be used.
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
   #### Note: The codecov.yaml used in ci will not be updated until your branch is merged to master. This is expected.
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
        files:
          - README.md

    checksum:
      name_template: checksums.txt

    # Add a changelog
    changelog:
      sort: asc
    ```
   If this is not a library release, please see any other `.goreleaser.yml` file for an example. It is important the production docker files are put in the [`docker/`](docker/) directory and named `[module_name].Dockerfile` so that others can find them.
7. Create a `README.md` file in the directory. This should include a description of the module and a link to the godoc page and go report card <!-- (for vanity) -->. If you're not sure what to put in the readme, look at the other modules for inspiration. If the module is on the more complex side, consider including a directory tree like we have [here](#directory-structure). If the application is runnable, instructions on running should be present.
8. Add the new directory and description to the [Directory Structure](#directory-structure) section of this [README](./README.md) in alphabetical order. <!-- TODO: we should really lint that this is in alphabetical order. Also, markdown files should be linted for broken links. -->

