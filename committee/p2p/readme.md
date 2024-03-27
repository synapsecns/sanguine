# p2p spec

1. Listen to transactions. These come from two souces:
    - on-chain listeners. These can be added through an interface from another module, likely asynchronously or through the db
    - peers: these can be got from peers via the announce transaction module. Signing peers are expected to ack.
