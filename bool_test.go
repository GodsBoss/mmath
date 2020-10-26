package mmath_test

import (
	"fmt"
	"testing"

	"github.com/GodsBoss/mmath"
)

func TestConstantBool(t *testing.T) {
	t.Parallel()

	testcases := map[string]testcaseBool{
		"true": {
			calculation:   mmath.NewTrue(),
			expectedValue: true,
		},
		"false": {
			calculation:   mmath.NewFalse(),
			expectedValue: false,
		},
	}

	runTestcasesBool(t, testcases)
}

func TestNot(t *testing.T) {
	t.Parallel()

	testcases := map[string]testcaseBool{
		"error": {
			calculation:       mmath.NewNot(mmath.NewFailingCalculation(fmt.Errorf("whoopsie"))),
			expectedErrorFunc: errorContainsString("whoopsie"),
		},
	}

	runTestcasesBool(t, testcases)
}

func TestInt64Comparison(t *testing.T) {
	t.Parallel()

	testcases := map[string]testcaseBool{
		"errors": {
			calculation: mmath.NewInt64Equals(
				mmath.NewFailingCalculation(fmt.Errorf("broken")),
				mmath.NewFailingCalculation(fmt.Errorf("meh")),
			),
			expectedErrorFunc: errorAnd(
				errorContainsString("broken"),
				errorContainsString("meh"),
			),
		},
	}

	runTestcasesBool(t, testcases)
}

func runTestcasesBool(t *testing.T, testcases map[string]testcaseBool) {
	for name := range testcases {
		testcaseBool := testcases[name]

		t.Run(
			name,
			func(t *testing.T) {
				t.Parallel()

				actualValue, actualErr := testcaseBool.calculation.CalculateBool()

				if actualValue != testcaseBool.expectedValue {
					t.Errorf(
						"expected calculation result value to be %t, but got %t",
						testcaseBool.expectedValue,
						actualValue,
					)
				}

				if testcaseBool.expectedErrorFunc == nil && actualErr != nil {
					t.Errorf("expected no error, but got %+v", actualErr)
				}

				if testcaseBool.expectedErrorFunc != nil {
					if actualErr != nil {
						testcaseBool.expectedErrorFunc(t, actualErr)
					}
					if actualErr == nil {
						t.Errorf("expected non-nil error")
					}
				}
			},
		)
	}
}

type testcaseBool struct {
	// calculation is the calculation executed by the test. It's result is
	// compared against the expected values.
	calculation mmath.CalculationBool

	// expectedValue is the value the calculation should return.
	expectedValue bool

	// expectedErrorFunc checks the error. This being nil is equivalent to
	// checking wether the error should be nil.
	expectedErrorFunc errorTest
}
