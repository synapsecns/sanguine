package core

// PtrTo returns a pointer to the given value.
func PtrTo[T any](v T) *T {
	return &v
}
