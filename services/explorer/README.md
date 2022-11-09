
# Explorer

See [#167](https://github.com/synapsecns/sanguine/issues/167) to learn more.

To access the clickhouse database, you can use the following command from the docker image:
# Refactor TODO
- basics, spacing, comments, function headers
- readme
- split up the consumer folder
- clean up the parser so it only parses and the types dont get switched up in the helper files for each topic
- consolidate contracts folder (solidity in one folder, either in /contracts or in /external
- see if resolvers + entirity of graphql can do without scribe (speedup maybe?)

```bash
clickhouse-client --database=clickhouse_test --user=clickhouse_test --password=clickhouse_test
```

