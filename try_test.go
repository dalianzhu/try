package try

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func Assert(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
	}
}

func Test_Try(t *testing.T) {
	// Check if panic is being handled
	_, err := Try1WithError(func() (int, error) {
		panic("hello")
		return 0, nil
	})
	// The returned error is a wrapped type of ErrCallbackPanic
	Assert(t, errors.Is(err, ErrCallbackPanic), true)

	// Check if the error is being handled
	_, err = Try1WithError(func() (string, error) {
		return "123", fmt.Errorf("test error")
	})
	// Returns the error returned by the callback
	Assert(t, err != nil, true)

	tryVal1, tryVal2, tryVal3, err := Try3WithError(func() (string, int, bool, error) {
		return "123", 0, false, nil
	})
	Assert(t, tryVal1, "123")
	Assert(t, tryVal2, 0)
	Assert(t, tryVal3, false)
	Assert(t, err, nil)
}
