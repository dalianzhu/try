package try

func Try1Or[A any](callback func() (A, error), fallbackA A) (ret A) {
	_ = tryWithPanic(func() {
		ret, err := callback()
		if err == nil {
			fallbackA = ret
		}
	})

	return fallbackA
}

// Try2Or callback takes 2 return values, if the callback function panics or error is not nil,
// returns the fallback values, otherwise returns the values of callback
func Try2Or[A, B any](callback func() (A, B, error), fallbackA A, fallbackB B) (ret A, ret2 B) {
	_ = tryWithPanic(func() {
		ret, ret2, err := callback()
		if err == nil {
			fallbackA = ret
			fallbackB = ret2
		}
	})

	return fallbackA, fallbackB
}

// Try3Or callback takes 3 return values, if the callback function panics or error is not nil,
// returns the fallback values, otherwise returns the values of callback
func Try3Or[A, B, C any](callback func() (A, B, C, error),
	fallbackA A, fallbackB B, fallbackC C,
) (ret A, ret2 B, ret3 C) {
	_ = tryWithPanic(func() {
		ret, ret2, ret3, err := callback()
		if err == nil {
			fallbackA = ret
			fallbackB = ret2
			fallbackC = ret3
		}
	})

	return fallbackA, fallbackB, fallbackC
}
