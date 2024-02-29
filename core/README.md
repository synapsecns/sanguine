# Core

[![Go Reference](https://pkg.go.dev/badge/github.com/synapsecns/sanguine/core.svg)](https://pkg.go.dev/github.com/synapsecns/sanguine/core)
[![Go Report Card](https://goreportcard.com/badge/github.com/synapsecns/sanguine/core)](https://goreportcard.com/report/github.com/synapsecns/sanguine/core)

Core contains common libraries used across the synapse Go repositories.


## Directory Structure

<pre>
root
├── <a href="./bytemap">bytemap</a>: Implements a map using `[]rune` or `[]byte` instead of `string`
├── <a href="./commandline">commandline</a>: Provides utilities for creating an interactive shell command for a [`cli`](github.com/urfave/cli) using the [`go-prompt`](github.com/c-bata/go-prompt) package.
├── <a href="./config">config</a>: Contains the configuration for the core package.
├── <a href="./dbcommon">dbcommon</a>: Contains common database utilities used with gorm.
├── <a href="./dockerutil">dockerutil</a>: Provides tools for working with Docker.
├── <a href="./ginhelper">ginhelper</a>: Contains a set of utilities for working with the Gin framework and a set of common middleware.
├── <a href="./mapmutex">mapmutex</a>: Implements a map that uses a mutex to protect concurrent access.
├── <a href="./merkle">merkle</a>: Provides a go based merkle tree implementation.
├── <a href="./metrics">metrics</a>: Provides a set of utilities for working with metrics/otel tracing.
├── <a href="./mocktesting">mocktesting</a>: Provides a mocked tester for use with `testing.TB`
├── <a href="./observer">observer</a>: Provides an interface for adding/removing listeners.
├── <a href="./processlog">processlog</a>: Provides a way to interact with detatched processes as streams.
├── <a href="./retry">retry</a>: Retries a function until it succeeds or the timeout is reached. This comes with a set of backoff strategies/options.
├── <a href="./server">server</a>: Provides a context-safe server that can be used to start/stop a server.
├── <a href="./testsuite">testsuite</a>: Provides a wrapper around testify/suite.
├── <a href="./threaditer">threaditer</a>: Provides a thread-safe generic iterator for a slice.
</pre>


