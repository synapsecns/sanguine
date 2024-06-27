# Screening API

[![Go Reference](https://pkg.go.dev/badge/github.com/synapsecns/sanguine/contrib/screener-api.svg)](https://pkg.go.dev/github.com/synapsecns/sanguine/contrib/screener-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/synapsecns/sanguine/contrib/screener-api)](https://goreportcard.com/report/github.com/synapsecns/sanguine/contrib/screener-api)

The screening api provides a simple restful interface for checking whether an address is blocked or not against a multiple data sources. Right now, two data sources are supported:

- Blacklist URL: a json list of addresses that are blocked
- Chainalysis: the Entity API runs a screen against an address to quantify the risk associated with it, `Severe`, `High`, `Medium`, or `Low`.

Addresses themselves are checked against specific rulesets:

`https://screener-url/[address]`

<pre>
root
├── <a href="./chainalysis">chainalysis</a>: chainalysis client stub.
├── <a href="./client">client</a>: client library for using the screening api.
├── <a href="./cmd">cmd</a>: contains the command line interface to be used for the screener.
├── <a href="./config">config</a>: Yaml config struct/parsing.
├── <a href="./db">db</a>: db interface for the screener.
├── <a href="./screener">screener</a>: screening code.
</pre>
