package iterfunc

// New0 returns an iterator function that invokes yield until next returns false.
func New0[T any](next func() bool) func(yield func() bool) bool {
	return func(yield func() bool) bool {
		for next() {
			if !yield() {
				return false
			}
		}
		return true
	}
}

// New returns an iterator function that yields one value until next returns false.
func New[T any](next func() bool, value func() T) func(yield func(T) bool) bool {
	return func(yield func(T) bool) bool {
		for next() {
			if !yield(value()) {
				return false
			}
		}
		return true
	}
}

// New2 returns an iterator function that yields two values until next returns false.
func New2[T, T2 any](next func() bool, value func() (T, T2)) func(yield func(T, T2) bool) bool {
	return func(yield func(T, T2) bool) bool {
		for next() {
			if !yield(value()) {
				return false
			}
		}
		return true
	}
}
