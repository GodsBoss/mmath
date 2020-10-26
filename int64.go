package mmath

// CalculationInt64 represents a calculation that returns an int64.
type CalculationInt64 interface {
	// CalculateInt64 returns the int64 value calculated by this calculator.
	CalculateInt64() (int64, error)
}

// CalculationInt64Func implements CalculationInt64 by wrapping a function.
type CalculationInt64Func func() (int64, error)

// CalculateInt64 calls f and returns its result.
func (f CalculationInt64Func) CalculateInt64() (int64, error) {
	return f()
}

// ConstantInt64 returns a calculation which always returns the same value and
// no error.
func ConstantInt64(c int64) CalculationInt64 {
	return CalculationInt64Func(
		func() (int64, error) {
			return c, nil
		},
	)
}
