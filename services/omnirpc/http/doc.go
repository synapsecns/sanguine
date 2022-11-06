// Package http contains different http clients that can be used by the request forwarder
// we use this for benchmarking and because fasthttp is still in beta
// support for redirects/behavior we may encounter in the wild needs to be tested
// better before we abandan an alternative
//
// Additionally: fasthttp doesn't support context cancellation
package http
