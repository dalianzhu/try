package try

import (
	"fmt"
	"testing"
)

func TestTry1Or(t *testing.T) {
	// Check if panic is being handled, returns the fallback value
	ret1 := Or(func() (int, error) {
		panic("hello")
		return 0, nil
	}, 1)

	Assert(t, ret1, 1)

	// Check if the error is nil, returns the value of the callback
	ret2 := Or(func() (string, error) {
		return "123", nil
	}, "fallback")
	Assert(t, ret2, "123")

	// Check if the error is not nil, returns the fallback value
	ret3 := Or(func() (string, error) {
		return "123", fmt.Errorf("test error")
	}, "fallback")
	Assert(t, ret3, "fallback")
}

func TestTry2Or(t *testing.T) {
	// Check if panic is being handled, returns the fallback values
	gotRet, gotRet2 := Or2(func() (int, int, error) {
		panic("hello")
		return 0, 0, nil
	}, 1, 2)
	Assert(t, gotRet, 1)
	Assert(t, gotRet2, 2)

	// Check if error is being handled, returns the fallback values
	gotRet, gotRet2 = Or2(func() (int, int, error) {
		return 0, 0, fmt.Errorf("test error")
	}, 3, 4)
	Assert(t, gotRet, 3)
	Assert(t, gotRet2, 4)

	gotRet, gotRet2 = Or2(func() (int, int, error) {
		return 0, 0, nil
	}, 3, 4)
	Assert(t, gotRet, 0)
	Assert(t, gotRet2, 0)
}
