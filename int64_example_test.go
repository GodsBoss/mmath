package mmath_test

import (
	"github.com/GodsBoss/mmath"

	"fmt"
)

func ExampleConstantInt64() {
	i64 := mmath.ConstantInt64(100)

	v, err := i64.CalculateInt64()

	fmt.Printf("Value is %d.\n", v)
	if err != nil {
		fmt.Printf("Error is: %v\n", err)
	}

	// Output:
	// Value is 100.
}
