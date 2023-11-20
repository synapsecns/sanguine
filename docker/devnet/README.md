Local contains dockerfiles for local development. These files should be able to be built from the root of the repo without any external dependencies.

These don't need to be made for every package, just those involved in devnet.

# Let's keep track of commands here we'll need in our devnet script

1. cd `docker/devnet`
2. `docker compose build --progress=plain`

## Default Port Map:

| Container              | Port |
|------------------------|------|
| omnirpc                | 9001 |
| chain_a (chain_id: 42) | 8042 |
| chain_b (chain_id: 43) | 8043 |
| chain_c (chain_id: 44) | 8044 |
| scribe                 | 9002 |



