package errors

// AsType finds the first error in err's tree that matches the type E, and
// if one is found, returns that error value and true. Otherwise, it
// returns the zero value of E and false.
//
// The tree consists of err itself, followed by the errors obtained by
// repeatedly calling its Unwrap() error or Unwrap() []error method. When
// err wraps multiple errors, AsType examines err followed by a
// depth-first traversal of its children.
//
// An error err matches the type E if the type assertion err.(E) holds,
// or if the error has a method As(any) bool such that err.As(target)
// returns true when target is a non-nil *E. In the latter case, the As
// method is responsible for setting target.
func AsType[E error](err error) (E, bool) {
	if err == nil {
		var zero E
		return zero, false
	}
	var pe *E // lazily initialized
	return asType(err, &pe)
}

func asType[E error](err error, ppe **E) (_ E, _ bool) {
	for {
		if e, ok := err.(E); ok {
			return e, true
		}
		if x, ok := err.(interface{ As(any) bool }); ok {
			if *ppe == nil {
				*ppe = new(E)
			}
			if x.As(*ppe) {
				return **ppe, true
			}
		}
		switch x := err.(type) {
		case interface{ Unwrap() error }:
			err = x.Unwrap()
			if err == nil {
				return
			}
		case interface{ Unwrap() []error }:
			for _, err := range x.Unwrap() {
				if err == nil {
					continue
				}
				if x, ok := asType(err, ppe); ok {
					return x, true
				}
			}
			return
		default:
			return
		}
	}
}
