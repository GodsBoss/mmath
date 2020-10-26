package mmath_test

import (
	"github.com/GodsBoss/mmath"

	"fmt"
)

func ExampleNewConstantBool() {
	t, err := mmath.NewConstantBool(true).CalculateBool()

	fmt.Printf("%t", t)
	if err != nil {
		fmt.Printf("Error is: %v\n", err)
	}

	// Output:
	// true
}
