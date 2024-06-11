// Package try provides a way to safely execute callback functions and handle panics and errors that may occur.
//
// As a quick start:
//
//	// Example for callback with value and error
//	result, err := try.WithError1(func() (int, error) {
//		// Your callback logic here
//		return 42, nil
//	})
//	if err != nil {
//		// if panic occurs, err will be ErrCallbackPanic
//		if errors.Is(err, try.ErrCallbackPanic) {
//			// callback panic: PANIC INFO, file:/FILEPATH/file_xxx.go line: 17
//			fmt.Println("Panic:", err)
//			return
//		}
//		fmt.Println("Error:", err)
//		return
//	}
//	fmt.Println("Result:", result)
//
//	// Example for panic handling
//	resultA := try.Or(func() (int, error) {
//		panic("panic")
//		return 42, nil
//	}, -1)
//	fmt.Println("Result A:", resultA) // Output: Result A: -1
package try

import (
	"errors"
	"fmt"
	"runtime"
)

// ErrCallbackPanic is the error type returned when the callback function panics
var ErrCallbackPanic = errors.New("callback panic")

// RuntimeCallerSkip when a panic occurs,
// skip the appropriate number of stack frames. so that file and line numbers print correctly.
// The default value is 2.
var RuntimeCallerSkip = 2

func tryWithPanic(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			// print the filepath and line number of the file where the panic occurred
			_, file, line, _ := runtime.Caller(RuntimeCallerSkip)
			err = fmt.Errorf("%w: %v, file:%s line: %d", ErrCallbackPanic, r, file, line)
		}
	}()
	f()

	return
}

// WithError callback returns an error.
// If the callback panics, the panic is returned as an error of type `ErrCallbackPanic`.
// If the callback returns an error, that error is returned. Otherwise, it returns `nil`.
func WithError(callback func() error) (err error) {
	panicErr := tryWithPanic(func() {
		err = callback()
	})
	if panicErr != nil {
		err = panicErr
	}

	return
}

// WithError1 callback returns a value and an error.
// If the callback panics, the panic is returned as an error of type `ErrCallbackPanic`.
// If the callback returns an error, that error is returned. Otherwise, it returns the value and `nil`.
func WithError1[A any](callback func() (A, error)) (ret A, err error) {
	panicErr := tryWithPanic(func() {
		ret, err = callback()
	})
	if panicErr != nil {
		err = panicErr
	}

	return
}

// WithError2 callback takes 2 return values.
// If the callback panics, the panic is returned as an error of type `ErrCallbackPanic`.
// If the callback returns an error, that error is returned. Otherwise, it returns the values and `nil`.
func WithError2[A, B any](callback func() (A, B, error)) (ret A, ret2 B, err error) {
	panicErr := tryWithPanic(func() {
		ret, ret2, err = callback()
	})
	if panicErr != nil {
		err = panicErr
	}

	return
}

// WithError3 callback takes 3 return values,
// If the callback panics, the panic is returned as an error of type `ErrCallbackPanic`.
// If the callback returns an error, that error is returned. Otherwise, it returns the values and `nil`.
func WithError3[A, B, C any](callback func() (A, B, C, error)) (ret A, ret2 B, ret3 C, err error) {
	panicErr := tryWithPanic(func() {
		ret, ret2, ret3, err = callback()
	})
	if panicErr != nil {
		err = panicErr
	}

	return
}
