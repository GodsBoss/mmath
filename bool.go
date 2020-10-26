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

// NewVariableBool creates a variable. In calculations, it returns the value
// it was set to. Calculating the result of a variable never fails.
func NewVariableBool() VariableBool {
	return &variableBool{}
}

// VariableBool represents a variable value, which can be set from the outside.
type VariableBool interface {
	CalculationBool

	Set(b bool)
}

type variableBool struct {
	b bool
}

func (v *variableBool) CalculateBool() (bool, error) {
	return v.b, nil
}

func (v *variableBool) Set(b bool) {
	v.b = b
}

// FailingCalculationBool creates a calculation which always fails. This is
// useful for testing when adding custom calculations.
func FailingCalculationBool(err error) CalculationBoolFunc {
	return func() (bool, error) {
		return false, err
	}
}
