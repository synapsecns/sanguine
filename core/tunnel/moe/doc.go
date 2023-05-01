// Package moe provides a tunnel that uses the Moe protocol.
// this module is currently broken due to an issue with moe.
// this is caused by two issues:
// 1. Moe returns localhost as the hostname: https://github.com/fasmide/remotemoe/blob/13a9ba0f5ddadffdf4fb395ebff7c366b88a0745/ssh/session.go#L363
// 2. Go assumes the response returned by forwarded-tcpip is a tcp address: https://github.com/golang/crypto/blob/master/ssh/tcpip.go#L211
// A PR has been made to fix this: https://github.com/fasmide/remotemoe/pull/18 and this module will be ready when that's merged
// should the PR be unmergable for an extended period of time, we can intercept the requests in the net.Conn through a reverse ssh proxy
package moe
