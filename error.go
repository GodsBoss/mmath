package mmath

import (
	"strings"
)

type errors []error

func (errs errors) Error() string {
	errStrings := make([]string, len(errs))
	for i := range errs {
		errStrings[i] = errs[i].Error()
	}
	return "multiple errors: " + strings.Join(errStrings, "; ")
}
