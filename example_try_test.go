package try_test

import (
	"errors"
	"fmt"

	"github.com/dalianzhu/try"
)

func ExampleWithError() {
	err := try.WithError(func() error {
		// Your callback logic here
		return nil
	})
	if err != nil {
		// if panic occurs, err type will be ErrCallbackPanic
		if errors.Is(err, try.ErrCallbackPanic) {
			// callback panic: PANIC INFO, file:/FILEPATH/file_xxx.go line: 17
			fmt.Println("Panic:", err)
			return
		}
		fmt.Println("Error:", err)
		return
	}
}

func ExampleWithError1() {
	result, err := try.WithError1(func() (int, error) {
		// Your callback logic here
		return 42, nil
	})
	if err != nil {
		// if panic occurs, err will be ErrCallbackPanic
		if errors.Is(err, try.ErrCallbackPanic) {
			// callback panic: PANIC INFO, file:/FILEPATH/file_xxx.go line: 17
			fmt.Println("Panic:", err)
			return
		}
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
	// Output: Result: 42
}

func ExampleOr() {
	result := try.Or(func() (int, error) {
		// Your callback logic here
		return 42, nil
	}, -1)

	fmt.Println("Result:", result)
	// Output: Result: 42
}

func ExampleOr_second() {
	// Example for panic handling
	resultA := try.Or(func() (int, error) {
		panic("panic")
		return 42, nil
	}, -1)

	fmt.Println("Result A:", resultA)
	// Output: Result A: -1
}
