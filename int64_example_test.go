package mmath_test

import (
	"github.com/GodsBoss/mmath"

	"fmt"
)

func ExampleNewConstantInt64() {
	i64 := mmath.NewConstantInt64(100)

	v, err := i64.CalculateInt64()

	fmt.Printf("Value is %d.\n", v)
	if err != nil {
		fmt.Printf("Error is: %v\n", err)
	}

	// Output:
	// Value is 100.
}

func ExampleVariableInt64() {
	v64 := mmath.NewVariableInt64()
	v64.Set(666)

	v, err := v64.CalculateInt64()

	fmt.Printf("Value is %d.\n", v)
	if err != nil {
		fmt.Printf("Error is: %v\n", err)
	}

	// Output:
	// Value is 666.
}

func ExampleNewSumInt64() {
	sum := mmath.NewSumInt64(
		mmath.NewConstantInt64(25),
		mmath.NewConstantInt64(-50),
		mmath.NewConstantInt64(42),
	)

	v, err := sum.CalculateInt64()

	fmt.Printf("Value is %d.\n", v)
	if err != nil {
		fmt.Printf("Error is: %v\n", err)
	}

	// Output:
	// Value is 17.
}
