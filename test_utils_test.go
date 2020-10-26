package mmath_test

import (
	"strings"
	"testing"
)

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
