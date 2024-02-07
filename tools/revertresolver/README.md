# Revert Resolver

This is a tool for parsing a revert message from a raw trace error corresponding to the hash of a custom error type. Sometimes this hash is not useful because we don't know the actual revert reason. For example, if when running `cast run <hash>`, we encounter this error:

`0xeb9266`

Then we can run:

`go run main.go -p ../../packages/contracts-core -f 0xeb9266`

which will resolve to the reason:

```
UnformattedAttestation(): 0xeb92662c687ecb0d91cd0350e030a511efe609e0ecfff5618dd8fbce75158ce1 (File: contracts/libs/memory/Attestation.sol, Line: 103)
```

Note that the `-f` flag is optional; if unspecified, the script will yield all revert reasons as output (with their corresponding hashes).

This tool is currently limited to error emits that don't take any parameters. A future version will use the abigen'd info.
