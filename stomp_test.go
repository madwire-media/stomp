package stomp_test

import (
	"testing"

	s "github.com/madwire-media/stomp"
)

type testSpy struct {
	called bool
}

func (t *testSpy) Errorf(s string, args ...interface{}) {
	t.called = true
}

func testMakeExpectSpy(t s.Tester) s.Tester {
	expect := s.MakeExpect(t)

	expected := 2
	actual := 1

	expect(expected).ToEqual(actual)

	return t
}

func TestMakeExpect(t *testing.T) {
	expect := s.MakeExpect(t)

	actual := 1
	expected := 1

	// this is our happy place
	expect(expected).ToEqual(actual)

	// this is our sad place
	spy := &testSpy{}

	spyResults := testMakeExpectSpy(spy)
	res := spyResults.(*testSpy)

	expect(res.called).ToEqual(true)
	// the error was called
}

func TestAssertionMethods(t *testing.T) {
	expect := s.MakeExpect(t)

	s.Describe("test `==` comparison", func() {
		expect(1).ToEqual(1)
	})

	s.Describe("test deep comparison", func() {
		expect([]int{1, 2, 3}).ToDeepEqual([]int{1, 2, 3})
	})
	s.Describe("test ToNotEqual", func() {
		expect(2 + 2).ToNotEqual(5)
	})
}
