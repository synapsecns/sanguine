package chainlist

// TODO: I'm not a huge fan of this approach. It can create inconsistent diffs. We do provide an online alternative, but it's a bit annoying to use
// because of network flakes in ci. One other option is to setup a github action cron job to auto update this. We'll likely do this at some point to get
// around the non-deterministic go-generation behavior here
//go:generate go run github.com/u-root/u-root/cmds/core/wget -O chains.json https://chainid.network/chains.json
// ignore this line: go:generate cannot be the last line of a file
