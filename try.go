package try

import (
	"errors"
	"fmt"
)

// ErrCallbackPanic is the error type returned when the callback function panics
var ErrCallbackPanic = errors.New("callback panic")

func tryWithPanic(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%w: %v", ErrCallbackPanic, r)
		}
	}()
	f()

	return
}

// Try1WithError callback takes 1 return value, if the callback function panics, returns the error from the callback,
// if the callback returns an error, returns the error from the callback, otherwise returns the value and nil
func Try1WithError[A any](callback func() (A, error)) (ret A, err error) {
	panicErr := tryWithPanic(func() {
		ret, err = callback()
	})
	if panicErr != nil {
		err = panicErr
	}

	return
}

// Try2WithError callback takes 2 return values, if the callback function panics, returns the error from the callback,
// if the callback returns an error, returns the error from the callback, otherwise returns the values and nil
func Try2WithError[A, B any](callback func() (A, B, error)) (ret A, ret2 B, err error) {
	panicErr := tryWithPanic(func() {
		ret, ret2, err = callback()
	})
	if panicErr != nil {
		err = panicErr
	}

	return
}

// Try3WithError callback takes 3 return values, if the callback function panics, returns the error from the callback,
// if the callback returns an error, returns the error from the callback, otherwise returns the values and nil
func Try3WithError[A, B, C any](callback func() (A, B, C, error)) (ret A, ret2 B, ret3 C, err error) {
	panicErr := tryWithPanic(func() {
		ret, ret2, ret3, err = callback()
	})
	if panicErr != nil {
		err = panicErr
	}

	return
}
