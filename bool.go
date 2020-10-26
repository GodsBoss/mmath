package mmath

// CalculationBool represents a calculation that returns an bool.
type CalculationBool interface {
	// CalculateBool returns the bool value calculated by this calculator.
	CalculateBool() (bool, error)
}
