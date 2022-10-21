package metrics

type Tracer interface {
	// Start starts the tracer.
	Start()
	// Stop stops the tracer.
	Stop()
}
