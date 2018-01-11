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
}

// teser : used for testing `MakeExpect`
// otherwise, *testing.T is used
type tester interface {
	Errorf(string, ...interface{})
}

// MakeExpect :: tester -> expected -> actual
// Takes *testing.T | tester
// example:
//  expect := MakeExpect(t)
//  expect(1).toEqual(1)
func MakeExpect(t tester) func(interface{}) *testChain {
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
		}
	}
}

// Describe allows you to add some output to `go test -v`
func Describe(msg string, f func()) {
	fmt.Printf(" * %s\n", msg)
	f()
}
