// Package stomp provides assertion helpers for testing.
package stomp

import (
	"fmt"
	"reflect"
)

// overview of how comparison works in go:
// https://medium.com/golangspec/comparison-operators-in-go-910d9d788ec0

// testChain contains the chainable assertion funcs
//   returned by the expect func created by `MakeExpect`
type testChain struct {
	ToEqual     func(interface{})
	ToDeepEqual func(interface{})
	ToNotEqual  func(interface{})
}

// Tester interface for easier testing
type Tester interface {
	Errorf(string, ...interface{})
}

// MakeExpect binds *testing.T to the returned expect function
//  expect is a `testChain` instance and implements those methods:
//    - ToEqual
//    - ToDeepEqual
//    - ToNotEqual
func MakeExpect(t Tester) func(interface{}) *testChain {
	return func(expected interface{}) *testChain {
		return &testChain{
			ToEqual: func(actual interface{}) {
				if expected != actual {
					t.Errorf("[!] error: expected %v, actual: %v\n", expected, actual)
				}
			},
			ToDeepEqual: func(actual interface{}) {
				if !reflect.DeepEqual(expected, actual) {
					t.Errorf("[!] error: expected %v, actual: %v\n", expected, actual)
				}
			},
			ToNotEqual: func(actual interface{}) {
				if expected == actual {
					t.Errorf("[!] error: expected %v, actual: %v  NOT to equal\n", expected, actual)
				}
			},
		}
	}
}

// Describe allows you to add output to `go test -v`
func Describe(msg string, f func()) {
	fmt.Printf(" * %s\n", msg)
	f()
}

// SameStringSlice compares two string slices regardless of order
func SameStringSlice(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	diff := make(map[string]int, len(x))

	for _, _x := range x {
		diff[_x]++
	}

	for _, _y := range y {
		if _, ok := diff[_y]; !ok {
			return false
		}

		diff[_y]--
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	return len(diff) == 0
}
