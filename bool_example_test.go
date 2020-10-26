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

func ExampleNewVariableBool() {
	v := mmath.NewVariableBool()
	v.Set(true)

	b, err := v.CalculateBool()

	fmt.Printf("%t\n", b)
	if err != nil {
		fmt.Printf("Error is: %v\n", err)
	}

	v.Set(false)

	b, err = v.CalculateBool()

	fmt.Printf("%t\n", b)
	if err != nil {
		fmt.Printf("Error is: %v\n", err)
	}

	// Output:
	// true
	// false
}
