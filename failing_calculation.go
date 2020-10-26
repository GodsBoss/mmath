package mmath

// FailingCalculation is a helper type useful for tests. It implements all
// calculations, but returns errors for all of them.
type FailingCalculation struct {
	// Err is the error returned by all calculations made by FailingCalculation.
	Err error
}

// NewFailingCalculation creates a FailingCalculation.
func NewFailingCalculation(err error) FailingCalculation {
	return FailingCalculation{
		Err: err,
	}
}

// CalculateBool returns calc.Err.
func (calc FailingCalculation) CalculateBool() (bool, error) {
	return false, calc.Err
}

// CalculateInt64 returns calc.Err.
func (calc FailingCalculation) CalculateInt64() (int64, error) {
	return 0, calc.Err
}
