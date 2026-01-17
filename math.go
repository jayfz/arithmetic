// Package arithmetic provides a set of simple mathematical functio
package arithmetic

import "golang.org/x/exp/constraints"

// Add takes two int values and returns the sum of them
// See more at [MathIsFun]
// [MathIsFun]: https://www.mathsisfun.com/numbers/addi

type Number interface {
	constraints.Float | constraints.Integer
}

func Add[T Number](num1, num2 T) T {
	return num1 + num2
}
