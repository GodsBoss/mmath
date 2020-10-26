package mmath_test

import (
	"fmt"
	"strings"

	"github.com/GodsBoss/mmath"

	"testing"
)

func TestSumInt64(t *testing.T) {
	t.Parallel()

	testcases := map[string]struct {
		// calculation is the calculation executed by the test. It's result is
		// compared against the expected values.
		calculation mmath.CalculationInt64

		// expectedValue is the value the calculation should return.
		expectedValue int64

		// expectedErrorFunc checks the error. This being nil is equivalent to
		// checking wether the error should be nil.
		expectedErrorFunc errorTest
	}{
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

type errCalc struct {
	err error
}

func (c errCalc) CalculateInt64() (int64, error) {
	return 0, c.err
}

func newErrCalc(err error) *errCalc {
	return &errCalc{
		err: err,
	}
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