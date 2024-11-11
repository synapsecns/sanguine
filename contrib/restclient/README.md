# REST Client

Auto-generated REST client for the Synapse Protocol REST API.

## Usage

```go
package main

import "github.com/synapsecns/sanguine/contrib/restclient"
func main() {
  client := restclient.NewClient("http://localhost:3000")
  // Use the client methods...
}

```

> **Note:** This stub is experimental and may be removed without notice until this README indicates otherwise.

## TODOs

- Implement built-in/native tracing for the REST client to monitor and trace API calls.
- Add a health endpoint to check the status and health of the REST client.
- Improve error handling and logging mechanisms.
- Enhance the client with retry logic for transient errors.
- Write comprehensive unit tests and integration tests for the client methods.
