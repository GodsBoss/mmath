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

	v64.Set(667)

	v, err = v64.CalculateInt64()

	fmt.Printf("Value is %d.\n", v)
	if err != nil {
		fmt.Printf("Error is: %v\n", err)
	}

	// Output:
	// Value is 666.
	// Value is 667.
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

func ExampleNewProductInt64() {
	product := mmath.NewProductInt64(
		mmath.NewConstantInt64(3),
		mmath.NewConstantInt64(15),
		mmath.NewConstantInt64(2),
	)

	v, err := product.CalculateInt64()

	fmt.Printf("Value is %d.\n", v)
	if err != nil {
		fmt.Printf("Error is: %v\n", err)
	}

	// Output:
	// Value is 90.
}

func ExampleNewConditionalInt64() {
	boolVar := mmath.NewVariableBool()
	decision := mmath.NewConditionalInt64(
		boolVar,
		mmath.NewConstantInt64(23),
		mmath.NewConstantInt64(9001),
	)

	boolVar.Set(true)

	v, err := decision.CalculateInt64()

	fmt.Printf("Value is %d.\n", v)
	if err != nil {
		fmt.Printf("Error is: %v\n", err)
	}

	boolVar.Set(false)

	v, err = decision.CalculateInt64()

	fmt.Printf("Value is %d.\n", v)
	if err != nil {
		fmt.Printf("Error is: %v\n", err)
	}

	// Output:
	// Value is 23.
	// Value is 9001.
}

func ExampleNewCreateBinaryInt64() {
	add := func(left, right int64) int64 {
		return left + right
	}

	createAddCalc := mmath.NewCreateBinaryInt64(add)

	addCalc := createAddCalc(
		mmath.NewConstantInt64(17),
		mmath.NewConstantInt64(71),
	)

	v, err := addCalc.CalculateInt64()

	fmt.Printf("Value is %d.\n", v)
	if err != nil {
		fmt.Printf("Error is: %v\n", err)
	}

	// Output:
	// Value is 88.
}

func ExampleNewSignumInt64() {
	var i int64
	for i = -1; i <= 1; i++ {
		v, err := mmath.NewSignumInt64(mmath.NewConstantInt64(i)).CalculateInt64()

		fmt.Printf("Value is %d.\n", v)
		if err != nil {
			fmt.Printf("Error is: %v\n", err)
		}
	}

	// Output:
	// Value is -1.
	// Value is 0.
	// Value is 1.
}

func ExampleNewReduceLeft() {
	v, err := mmath.NewReduceLeft(
		func(current, next int64) (int64, error) {
			return current*10 + next, nil
		},
		mmath.NewConstantInt64(2),
		[]mmath.CalculationInt64{
			mmath.NewConstantInt64(3),
			mmath.NewConstantInt64(7),
			mmath.NewConstantInt64(5),
		},
	).CalculateInt64()

	fmt.Printf("Value is %d.\n", v)
	if err != nil {
		fmt.Printf("Error is: %v\n", err)
	}

	// Output:
	// Value is 2375.
}
