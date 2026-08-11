// Package addition offers a complicated way of adding two numbers
package addition

import "golang.org/x/exp/constraints"

// Number is the set of types that support addition via the + operator:
// any integer or floating-point type.
type Number interface {
	constraints.Integer | constraints.Float
}

// Add adds two numbers and returns the result.
// More information about [Addition] can be found on the mathisfun website.
// [Addition]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
