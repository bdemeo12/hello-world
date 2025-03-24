// To execute Go code, please declare a func main() in a package "main"

package main

import "fmt"

func main() {

	array := []int{1, 2, 3, 4, 5}
	position := 2
	val := 3

	// #   for_list([1,2,3,4,5]).elem_at(2).should_be(equal_to(3))
	t := newTestingPackage()

	fmt.Println(t.For_list(array).Elem_at(position).Should_be(equal_to(val)))

}

type TestingPackage interface {
	For_list(arr []int) *testingPackage
	Elem_at(n int) *testingPackage
	Should_be(f func(n int) bool) bool
}

type testingPackage struct {
	Array  []int
	Actual int
}

func newTestingPackage() TestingPackage {
	return &testingPackage{}
}

func (t *testingPackage) For_list(arr []int) *testingPackage {
	t.Array = arr
	return t
}

func (t *testingPackage) Elem_at(n int) *testingPackage {
	t.Actual = t.Array[n]
	return t
}

func (t *testingPackage) Should_be(f func(n int) bool) bool {
	return f(t.Actual)
}

// Testing Functions:

func equal_to(expected int) func(actual int) bool {
	return func(actual int) bool {
		return actual == expected
	}
}

func greater_than(expected int) func(actual int) bool {
	return func(actual int) bool {
		return actual > expected
	}
}

func less_than(expected int) func(actual int) bool {
	return func(actual int) bool {
		return actual < expected
	}
}
