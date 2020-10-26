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

// NewNot returns the negated value of the wrapped calculation, except for when
// that calculation returns an error. In that case, the error is returned.
func NewNot(wrappedCalc CalculationBool) CalculationBoolFunc {
	return func() (bool, error) {
		b, err := wrappedCalc.CalculateBool()
		if err != nil {
			return false, err
		}
		return !b, nil
	}
}

// NewInt64Equals returns wether the result of the first and second calculations
// are equal. If one or both return an error, return an error combining those
// errors instead.
func NewInt64Equals(first, second CalculationInt64) CalculationBoolFunc {
	return func() (bool, error) {
		var errs errors

		firstValue, err := first.CalculateInt64()
		if err != nil {
			errs = append(errs, err)
		}

		secondValue, err := second.CalculateInt64()
		if err != nil {
			errs = append(errs, err)
		}

		if len(errs) > 0 {
			return false, errs
		}

		return firstValue == secondValue, nil
	}
}

// FailingCalculationBool creates a calculation which always fails. This is
// useful for testing when adding custom calculations.
func FailingCalculationBool(err error) CalculationBoolFunc {
	return func() (bool, error) {
		return false, err
	}
}
