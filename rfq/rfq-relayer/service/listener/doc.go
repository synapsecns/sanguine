/*
Package listener provides the event listener service for the RFQ relayer.
This package implements a service that listens for events on the FastBridge contract (for one chain).
Functionality includes:
- Listening for BridgeRequested and BridgeRelayed events
- Dynamically ranging over both blocks in and out of the confirmation range
- Caching already seen events in a LRU cache
- Sending new unconfirmed events to the RFQ relayer quote freezing service
- Storing new confirmed events to the RFQ relayer event processing service
*/

package listener
