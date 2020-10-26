package mmath_test

import (
	"github.com/GodsBoss/mmath"

	"fmt"
)

func ExampleNewConstantBool() {
	b, err := mmath.NewConstantBool(true).CalculateBool()

	fmt.Printf("%t", b)
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

func ExampleNewNot() {
	b, err := mmath.NewNot(mmath.NewConstantBool(true)).CalculateBool()

	fmt.Printf("%t\n", b)
	if err != nil {
		fmt.Printf("Error is: %v\n", err)
	}

	// Output:
	// false
}
