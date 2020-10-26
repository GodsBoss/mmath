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
func NewConstantBool(b bool) CalculationBool {
	return CalculationBoolFunc(
		func() (bool, error) {
			return b, nil
		},
	)
}
