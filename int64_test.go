package mmath_test

import (
	"fmt"
	"testing"

	"github.com/GodsBoss/mmath"
)

func TestInt64(t *testing.T) {
	t.Parallel()

	testcases := map[string]testcaseInt64{
		"sum/no_summands": {
			calculation:   mmath.NewSumInt64(),
			expectedValue: 0,
		},
		"sum/some_numbers": {
			calculation: mmath.NewSumInt64(
				mmath.NewConstantInt64(10),
				mmath.NewConstantInt64(-5),
				mmath.NewConstantInt64(18),
			),
			expectedValue: 23,
		},
		"sum/errors": {
			calculation: mmath.NewSumInt64(
				mmath.NewFailingCalculation(fmt.Errorf("foobar")),
				mmath.NewFailingCalculation(fmt.Errorf("xyz")),
			),
			expectedErrorFunc: errorAnd(
				errorContainsString("foobar"),
				errorContainsString("xyz"),
			),
		},
		"product/no_factors": {
			calculation:   mmath.NewProductInt64(),
			expectedValue: 1,
		},
		"product/some_numbers": {
			calculation: mmath.NewProductInt64(
				mmath.NewConstantInt64(3),
				mmath.NewConstantInt64(-2),
				mmath.NewConstantInt64(-7),
			),
			expectedValue: 42,
		},
		"product/errors": {
			calculation: mmath.NewProductInt64(
				mmath.NewFailingCalculation(fmt.Errorf("abc")),
				mmath.NewFailingCalculation(fmt.Errorf("123")),
			),
			expectedErrorFunc: errorAnd(
				errorContainsString("abc"),
				errorContainsString("123"),
			),
		},
		"conditional/error": {
			calculation: mmath.NewConditionalInt64(
				mmath.NewFailingCalculation(fmt.Errorf("hello")),
				mmath.NewConstantInt64(1),
				mmath.NewConstantInt64(2),
			),
			expectedErrorFunc: errorContainsString("hello"),
		},
		"binary/error": {
			calculation: mmath.NewCreateBinaryInt64(
				func(i, j int64) int64 {
					return i + j
				},
			)(
				mmath.NewFailingCalculation(fmt.Errorf("hello")),
				mmath.NewFailingCalculation(fmt.Errorf("world")),
			),
			expectedErrorFunc: errorAnd(
				errorContainsString("hello"),
				errorContainsString("world"),
			),
		},
		"signum/error": {
			calculation: mmath.NewSignumInt64(
				mmath.NewFailingCalculation(fmt.Errorf("xyz")),
			),
			expectedErrorFunc: errorContainsString("xyz"),
		},
		"reduceLeft/initialValueError": {
			calculation: mmath.NewReduceLeft(
				func(_, _ int64) (int64, error) {
					return 0, nil
				},
				mmath.NewFailingCalculation(fmt.Errorf("initial calculation failed")),
				[]mmath.CalculationInt64{},
			),
			expectedErrorFunc: errorContainsString("initial calculation failed"),
		},
		"reduceLeft/calculationsError": {
			calculation: mmath.NewReduceLeft(
				func(_, _ int64) (int64, error) {
					return 0, nil
				},
				mmath.NewConstantInt64(0),
				[]mmath.CalculationInt64{
					mmath.NewFailingCalculation(fmt.Errorf("nobody")),
					mmath.NewFailingCalculation(fmt.Errorf("knows")),
				},
			),
			expectedErrorFunc: errorAnd(
				errorContainsString("nobody"),
				errorContainsString("knows"),
			),
		},
		"reduceLeft/reduceError": {
			calculation: mmath.NewReduceLeft(
				func(_, _ int64) (int64, error) {
					return 0, fmt.Errorf("reduce failure")
				},
				mmath.NewConstantInt64(0),
				[]mmath.CalculationInt64{
					mmath.NewConstantInt64(0),
					mmath.NewConstantInt64(0),
				},
			),
			expectedErrorFunc: errorAnd(
				errorContainsString("reduce failure"),
			),
		},
	}

	for name := range testcases {
		testcaseInt64 := testcases[name]

		t.Run(
			name,
			func(t *testing.T) {
				t.Parallel()

				actualValue, actualErr := testcaseInt64.calculation.CalculateInt64()

				if actualValue != testcaseInt64.expectedValue {
					t.Errorf(
						"expected calculation result value to be %d, but got %d",
						testcaseInt64.expectedValue,
						actualValue,
					)
				}

				if testcaseInt64.expectedErrorFunc == nil && actualErr != nil {
					t.Errorf("expected no error, but got %+v", actualErr)
				}

				if testcaseInt64.expectedErrorFunc != nil {
					if actualErr != nil {
						testcaseInt64.expectedErrorFunc(t, actualErr)
					}
					if actualErr == nil {
						t.Errorf("expected non-nil error")
					}
				}
			},
		)
	}
}

type testcaseInt64 struct {
	// calculation is the calculation executed by the test. It's result is
	// compared against the expected values.
	calculation mmath.CalculationInt64

	// expectedValue is the value the calculation should return.
	expectedValue int64

	// expectedErrorFunc checks the error. This being nil is equivalent to
	// checking wether the error should be nil.
	expectedErrorFunc errorTest
}
