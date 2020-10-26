package mmath

// CalculationBool represents a calculation that returns an bool.
type CalculationBool interface {
	// CalculateBool returns the bool value calculated by this calculator.
	CalculateBool() (bool, error)
}

// CalculationBoolFunc implements CalculationBool by wrapping a function.
type CalculationBoolFunc func() (bool, error)

// CalculateBool calls f and returns its result.
func (f CalculationBoolFunc) CalculateBool() (bool, error) {
	return f()
}

// NewConstantBool returns a calculation which always returns the value passed
// on creation.
func NewConstantBool(b bool) CalculationBoolFunc {
	return func() (bool, error) {
		return b, nil
	}
}

// NewTrue returns a calculation which is always true.
func NewTrue() CalculationBool {
	return NewConstantBool(true)
}

// NewFalse returns a calculation which is always false.
func NewFalse() CalculationBool {
	return NewConstantBool(false)
}
