package types

// Snapshot is the snapshot interface.
type Snapshot interface {
	// States are the states of the snapshot.
	States() []State
}

type snapshot struct {
	states []State
}

// NewSnapshot creates a new snapshot.
func NewSnapshot(states []State) Snapshot {
	return &snapshot{
		states: states,
	}
}

func (s snapshot) States() []State {
	return s.states
}
