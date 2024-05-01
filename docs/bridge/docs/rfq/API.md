---
sidebar_position: 0
sidebar_label: API
---


The canonical implementation for the rfq api can be found [here](https://github.com/synapsecns/sanguine/tree/master/services/rfq/api). Please note that end-users and solvers will not need to run their own version of the API.

**Configuration:**

The RFQ API takes in a yaml config that allows the user to specify which contracts, chains and interfaces it should run on. The config is structured like this:

```yaml
database:
  type: mysql # can be other mysql or sqlite
  dsn: root:password@hostname:3306)/database?parseTime=true # should be the dsn of your database. If using sqlite, this can be a path
omnirpc_url: https://route-to-my-omnirpc # omnipc route
bridges:
  1: '0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E' # FastBridge address on ethereum
  10: '0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E' # FastBridge address on op
port: '8081'  # port to run your http server on
```

**Building From Source:**

1. `git clone [https://github.com/synapsecns/sanguin](https://github.com/synapsecns/sanguine-)e --recursive`
2. `cd sanguine/services/rfq`
3. `go run main.go --config /path/to/config.yaml`

**Running with Docker**

1. `docker run  --config /path/to/config`

