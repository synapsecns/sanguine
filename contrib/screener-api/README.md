# Screening API

[![Go Reference](https://pkg.go.dev/badge/github.com/synapsecns/sanguine/contrib/screener-api.svg)](https://pkg.go.dev/github.com/synapsecns/sanguine/contrib/screener-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/synapsecns/sanguine/contrib/screener-api)](https://goreportcard.com/report/github.com/synapsecns/sanguine/contrib/screener-api)

The screening api provides a simple restful interface for checking wether an address is blocked or not against a variety of data sources. Right now, two data sources are supported:

- Blacklist URL: a json list of addresses that are blocked
- TRM Labs: a list of rules that are used to determine if an address is blocked, can be different per "rule set"

Addresses themselves are checked against specific rulesets:

`https://screener-url/[ruleset]/address/[address]`

<pre>
root
├── <a href="./client">client</a>: client library for using the screening api.
├── <a href="./cmd">cmd</a>: contains the command line interface to be used for the screener.
├── <a href="./config">config</a>: Yaml config struct/parsing.
├── <a href="./db">db</a>: db interface for the screener.
├── <a href="./screener">screener</a>: screening code.
├── <a href="./trmlabs">trmlabs</a>: trm client stub.
</pre>
