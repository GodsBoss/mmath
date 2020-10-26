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

// NewVariableInt64 creates a variable. In calculations, it returns the value
// it was set to. Calculating the result of a variable never fails.
func NewVariableInt64() VariableInt64 {
	return &variableInt64{}
}

// VariableInt64 represents a variable value, which can be set from the outside.
type VariableInt64 interface {
	CalculationInt64

	// Set sets the variable. Afterwards, calling CalculateInt64() will return i.
	Set(i int64)
}

type variableInt64 struct {
	value int64
}

func (v *variableInt64) CalculateInt64() (int64, error) {
	return v.value, nil
}

func (v *variableInt64) Set(i int64) {
	v.value = i
}
