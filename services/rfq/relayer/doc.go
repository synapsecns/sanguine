// Package relayer contains the core of the relayer.
// The entire application is a db sourced event architecture
// so here we:
// 1. listen to the chains, mark any events as seen or our own events as complete
// 2. Check each status in the db and see what, if any, action needs to be taken. This makes the application safe against abrupt crashes.
package relayer
