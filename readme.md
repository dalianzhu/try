# Try Package

The `try` package provides utility functions to handle Go callback functions that may panic or return an error.

It ensures that panics are caught and returned as errors, allowing for more graceful error handling in your applications.

## Installation

To install the package, run:

```sh
go get github.com/dalianzhu/try
```

## Usage

### Handling Panics in Callbacks

The package defines a set of functions to safely execute callbacks that may return multiple values and an error. If a callback panics, the panic is caught and returned as an error.

### Functions
#### Try1WithError

```go
func Try1WithError[A any](callback func() (A, error)) (ret A, err error)
```

- **Description**: Executes a callback that returns one value and an error. 
If the callback panics, the panic is returned as an error of type `ErrCallbackPanic`. 

If the callback returns an error, that error is returned. Otherwise, it returns the value and `nil`.
- **Parameters**:
  - **callback** (`func() (A, error)`) - The callback function to be executed, which returns one value of type `A` and an error.
- **Returns**: The value returned by the callback or the fallback value.
  - **ret** (`A`) - The value returned by the callback function if no error or panic occurs.
  - **err** (`error`) - Returns an error if a panic occurs or if the callback function returns an error. Returns `nil` if no error or panic occurs.

#### Try2WithError

```go
func Try2WithError[A, B any](callback func() (A, B, error)) (ret A, ret2 B, err error)
```

- **Description**: Executes a callback that returns two values and an error. If the callback panics, the panic is returned as an error of type `ErrCallbackPanic`. 

If the callback returns an error, that error is returned. Otherwise, it returns the values and `nil`.
- **Parameters**:
    - **callback** (`func() (A, B, error)`) - The callback function to be executed, which returns two values of type `A` and `B`, and an error.
- **Returns**:
  - **ret** (`A`) - The first value returned by the callback function if no error or panic occurs.
  - **ret2** (`B`) - The second value returned by the callback function if no error or panic occurs.
  - **err** (`error`) - Returns an error if a panic occurs or if the callback function returns an error. Returns `nil` if no error or panic occurs.

#### Try3WithError

```go
func Try3WithError[A, B, C any](callback func() (A, B, C, error)) (ret A, ret2 B, ret3 C, err error)
```

- **Description**: Executes a callback that returns two values and an error. If the callback panics, the panic is returned as an error of type `ErrCallbackPanic`. 


Executes a callback that returns three values and an error. If the callback panics, the panic is returned as an error of type `ErrCallbackPanic`. 

If the callback returns an error, that error is returned. Otherwise, it returns the values and `nil`.

- **Parameters**:
  - **callback** (`func() (A, B, C, error)`) - The callback function to be executed, which returns three values of type `A`, `B`, and `C`, and an error.
- **Returns**:
  - **ret** (`A`) - The first value returned by the callback function if no error or panic occurs.
  - **ret2** (`B`) - The second value returned by the callback function if no error or panic occurs.
  - **ret3** (`C`) - The third value returned by the callback function if no error or panic occurs.
  - **err** (`error`) - Returns an error if a panic occurs or if the callback function returns an error. Returns `nil` if no error or panic occurs.
#### Try1Or

```go
func Try1Or[A any](callback func() (A, error), fallbackA A) (ret A)
```

- **Description**: Executes a callback function that returns one value and an error. If the callback panics or returns an error, the specified fallback value is returned. Otherwise, the callback's return value is returned.
- **Parameters**:
  - `callback`: A function that returns a value of type `A` and an error.
  - `fallbackA`: The fallback value to return if the callback panics or returns an error.
- **Returns**: The value returned by the callback or the fallback value.

#### Try2Or

```go
func Try2Or[A, B any](callback func() (A, B, error), fallbackA A, fallbackB B) (ret A, ret2 B)
```

- **Description**: Executes a callback function that returns two values and an error. If the callback panics or returns an error, the specified fallback values are returned. Otherwise, the callback's return values are returned.
- **Parameters**:
  - `callback`: A function that returns two values of types `A` and `B`, and an error.
  - `fallbackA`: The fallback value to return if the callback panics or returns an error.
  - `fallbackB`: The fallback value to return if the callback panics or returns an error.
- **Returns**: The values returned by the callback or the fallback values.

#### Try3Or

```go
func Try3Or[A, B, C any](callback func() (A, B, C, error), fallbackA A, fallbackB B, fallbackC C) (ret A, ret2 B, ret3 C)
```

- **Description**: Executes a callback function that returns three values and an error. If the callback panics or returns an error, the specified fallback values are returned. Otherwise, the callback's return values are returned.
- **Parameters**:
  - `callback`: A function that returns three values of types `A`, `B`, and `C`, and an error.
  - `fallbackA`: The fallback value to return if the callback panics or returns an error.
  - `fallbackB`: The fallback value to return if the callback panics or returns an error.
  - `fallbackC`: The fallback value to return if the callback panics or returns an error.
- **Returns**: The values returned by the callback or the fallback values.

## Example

### Try1WithError

```go
package main

import (
	"fmt"
	"github.com/dalianzhu/try"
)

func main() {
	result, err := try.Try1WithError(func() (int, error) {
		// Your callback logic here
		return 42, nil
	})
	if err != nil {
        // if panic occurs, err will be ErrCallbackPanic
        if errors.Is(err, try.ErrCallbackPanic) {
            fmt.Println("Panic:", err)
            return
        }
		fmt.Println("Error:", err)
        return
	} 
	fmt.Println("Result:", result)
}
```

### Try Or

```go
package main

import (
	"fmt"
	"github.com/dalianzhu/try"
)

func main() {
	// Example using Try1Or
	result := try.Try1Or(func() (int, error) {
		// Your callback logic here
		return 42, nil
	}, -1)
	fmt.Println("Result:", result) // Output: Result: 42

	// Example using Try2Or
	resultA, resultB := try.Try2Or(func() (int, string, error) {
		// Your callback logic here
		return 42, "hello", nil
	}, -1, "fallback")
	fmt.Println("Result A:", resultA, "Result B:", resultB) // Output: Result A: 42 Result B: hello

    // Example for panic handling
	resultA, resultB := try.Try2Or(func() (int, string, error) {
        panic("panic")
		return 42, "hello", nil
	}, -1, "fallback")
    fmt.Println("Result A:", resultA, "Result B:", resultB) // Output: Result A: -1 Result B: fallback
}
```

## Error Handling

The package defines a custom error type:

```go
var ErrCallbackPanic = errors.New("callback panic")
```

This error is returned when a panic occurs in the callback function, allowing you to distinguish between regular errors and panics.

## License
MIT license 

## Contribution

Feel free to submit issues, fork the repository and send pull requests. For major changes, please open an issue first to discuss what you would like to change.
