# stomp
status [![CircleCI](https://circleci.com/gh/madwire-media/stomp.svg?style=svg&circle-token=9ec978c27ff93132c7003c9235121f02d7839999)](https://circleci.com/gh/madwire-media/stomp)
> basic chained golang assertion library

put your foot down hard when it comes to `==`



### goals
 1. reduce code reuse
 2. make asserting more pleasant
 3. feel _almost_ like javascript assertions

## usage
### download
`go get -u github.com/madwire-media/stomp`

### example test
```
package stomp

import (
  "testing"

  "github.com/madwire-media/stomp"
)

func MyMap(inputArray []int, f func(int) int) []int {
  var ret []int

  for _, v := range inputArray {
    ret = append(ret, f(v))
  }
  return ret
}

func TestStomp(t *testing.T) {
  expect := stomp.MakeExpect(t)

  seed := []int{1, 2, 3}
  expected := []int{3, 6, 9}

  actual := MyMap(seed, func(i int) int {
    return i * 3
  })

  expect(expected[0]).ToEqual(actual[0])

  // use DeepEqual since we are using slices
  expect(expected).ToDeepEqual(actual)
}
```

## provided
### expect
`.ToEqual` standard `==` comparison

`.ToNotEqual` standard `!=` comparison

`.ToDeepEqual` uses `reflect.DeepEqual` to do a deep comparison 

**note**: Deep equality checking can be slow, and you should 
read up on golang equality if you get to a point where you use this!

### Describe
Add output to `go test -v`

i.e
```
...

Describe("test answer of life-ness of 42", func() {
  ...
  // expect some stuff
  ...
})
...
```

### helpers
#### SameStringSlice
Compares 2 string slices regaurdless of order

```
s1 := []string{"camp", "band", "once", "a", "time", "once", "upon"}
s2 := []string{"once", "upon", "a", "time", "at", "band", "camp", "once"}
expect(stomp.SameStringSlice(s1, s2)).toEqual(true)
```
