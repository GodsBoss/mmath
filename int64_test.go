package mmath_test

import (
	"fmt"
	"strings"

	"github.com/GodsBoss/mmath"

	"testing"
)

func TestSumInt64(t *testing.T) {
	t.Parallel()

	testcases := map[string]testcase{
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
				newErrCalc(fmt.Errorf("foobar")),
				newErrCalc(fmt.Errorf("xyz")),
			),
			expectedErrorFunc: errorAnd(
				errorContainsString("foobar"),
				errorContainsString("xyz"),
			),
		},
	}

	runTestcases(t, testcases)
}

func TestProductInt64(t *testing.T) {
	t.Parallel()

	testcases := map[string]testcase{
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
				newErrCalc(fmt.Errorf("abc")),
				newErrCalc(fmt.Errorf("123")),
			),
			expectedErrorFunc: errorAnd(
				errorContainsString("abc"),
				errorContainsString("123"),
			),
		},
	}

	runTestcases(t, testcases)
}

func runTestcases(t *testing.T, testcases map[string]testcase) {
	for name := range testcases {
		testcase := testcases[name]

		t.Run(
			name,
			func(t *testing.T) {
				t.Parallel()

				actualValue, actualErr := testcase.calculation.CalculateInt64()

				if actualValue != testcase.expectedValue {
					t.Errorf(
						"expected calculation result value to be %d, but got %d",
						testcase.expectedValue,
						actualValue,
					)
				}

				if testcase.expectedErrorFunc == nil && actualErr != nil {
					t.Errorf("expected no error, but got %+v", actualErr)
				}

				if testcase.expectedErrorFunc != nil {
					if actualErr != nil {
						testcase.expectedErrorFunc(t, actualErr)
					}
					if actualErr == nil {
						t.Errorf("expected non-nil error")
					}
				}
			},
		)
	}
}

type testcase struct {
	// calculation is the calculation executed by the test. It's result is
	// compared against the expected values.
	calculation mmath.CalculationInt64

	// expectedValue is the value the calculation should return.
	expectedValue int64

	// expectedErrorFunc checks the error. This being nil is equivalent to
	// checking wether the error should be nil.
	expectedErrorFunc errorTest
}

func newErrCalc(err error) mmath.CalculationInt64 {
	return mmath.FailingCalculationInt64(err)
}

type errorTest func(t *testing.T, actualErr error)

func errorAnd(fs ...errorTest) errorTest {
	return func(t *testing.T, actualErr error) {
		for i := range fs {
			fs[i](t, actualErr)
		}
	}
}

func errorContainsString(s string) errorTest {
	return func(t *testing.T, actualErr error) {
		if !strings.Contains(actualErr.Error(), s) {
			t.Errorf("expected error %+v to contain string '%s'", actualErr, s)
		}
	}
}
