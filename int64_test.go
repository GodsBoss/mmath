package mmath_test

import (
	"fmt"
	"testing"

	"github.com/GodsBoss/mmath"
)

func TestSumInt64(t *testing.T) {
	t.Parallel()

	testcases := map[string]testcaseInt64{
		"no_summands": {
			calculation:   mmath.NewSumInt64(),
			expectedValue: 0,
		},
		"some_numbers": {
			calculation: mmath.NewSumInt64(
				mmath.NewConstantInt64(10),
				mmath.NewConstantInt64(-5),
				mmath.NewConstantInt64(18),
			),
			expectedValue: 23,
		},
		"errors": {
			calculation: mmath.NewSumInt64(
				newErrCalcInt64(fmt.Errorf("foobar")),
				newErrCalcInt64(fmt.Errorf("xyz")),
			),
			expectedErrorFunc: errorAnd(
				errorContainsString("foobar"),
				errorContainsString("xyz"),
			),
		},
	}

	runTestcasesInt64(t, testcases)
}

func TestProductInt64(t *testing.T) {
	t.Parallel()

	testcases := map[string]testcaseInt64{
		"no_factors": {
			calculation:   mmath.NewProductInt64(),
			expectedValue: 1,
		},
		"some_numbers": {
			calculation: mmath.NewProductInt64(
				mmath.NewConstantInt64(3),
				mmath.NewConstantInt64(-2),
				mmath.NewConstantInt64(-7),
			),
			expectedValue: 42,
		},
		"errors": {
			calculation: mmath.NewProductInt64(
				newErrCalcInt64(fmt.Errorf("abc")),
				newErrCalcInt64(fmt.Errorf("123")),
			),
			expectedErrorFunc: errorAnd(
				errorContainsString("abc"),
				errorContainsString("123"),
			),
		},
	}

	runTestcasesInt64(t, testcases)
}

func TestConditionalInt64(t *testing.T) {
	t.Parallel()

	testcases := map[string]testcaseInt64{
		"error": {
			calculation: mmath.NewConditionalInt64(
				mmath.NewFailingCalculation(fmt.Errorf("hello")),
				mmath.NewConstantInt64(1),
				mmath.NewConstantInt64(2),
			),
			expectedErrorFunc: errorContainsString("hello"),
		},
	}

	runTestcasesInt64(t, testcases)
}

func runTestcasesInt64(t *testing.T, testcases map[string]testcaseInt64) {
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

func newErrCalcInt64(err error) mmath.CalculationInt64 {
	return mmath.NewFailingCalculation(err)
}
