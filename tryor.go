package try

// Or callback takes 1 return value, if the callback function panics or error is not nil,
// returns the fallback value, otherwise returns the value of callback.
func Or[A any](callback func() (A, error), fallbackA A) (ret A) {
	_ = tryWithPanic(func() {
		ret, err := callback()
		if err == nil {
			fallbackA = ret
		}
	})

	return fallbackA
}

// Or2 callback takes 2 return values, if the callback function panics or error is not nil,
// returns the fallback values, otherwise returns the values of callback
func Or2[A, B any](callback func() (A, B, error), fallbackA A, fallbackB B) (ret A, ret2 B) {
	_ = tryWithPanic(func() {
		ret, ret2, err := callback()
		if err == nil {
			fallbackA = ret
			fallbackB = ret2
		}
	})

	return fallbackA, fallbackB
}

// Or3 callback takes 3 return values, if the callback function panics or error is not nil,
// returns the fallback values, otherwise returns the values of callback
func Or3[A, B, C any](callback func() (A, B, C, error),
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
