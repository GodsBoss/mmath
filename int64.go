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

// NewConstantInt64 returns a calculation which always returns the same value and
// no error.
func NewConstantInt64(c int64) CalculationInt64Func {
	return func() (int64, error) {
		return c, nil
	}
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

// NewSumInt64 returns a calculation which returns the sum of all calculations
// passed to it. If one or more calculations fail, an error wrapping all those
// individual errors is returned.
func NewSumInt64(calculations ...CalculationInt64) CalculationInt64 {
	return sum(calculations)
}

type sum []CalculationInt64

func (s sum) CalculateInt64() (int64, error) {
	results, err := runCalculationsInt64(s...)
	if err != nil {
		return 0, err
	}

	var result int64
	for i := range results {
		result += results[i]
	}
	return result, nil
}

// NewProductInt64 returns a calculation which returns the product of all
// calculations passed to it. If one or more calculations fail, an error
// wrapping all those individual errors is returned.
func NewProductInt64(calculations ...CalculationInt64) CalculationInt64 {
	return product(calculations)
}

type product []CalculationInt64

func (p product) CalculateInt64() (int64, error) {
	results, err := runCalculationsInt64(p...)
	if err != nil {
		return 0, err
	}

	var result int64 = 1
	for i := range results {
		result *= results[i]
	}
	return result, nil
}

func runCalculationsInt64(calculations ...CalculationInt64) ([]int64, error) {
	var errs errors
	results := make([]int64, len(calculations))

	for i := range calculations {
		result, err := calculations[i].CalculateInt64()
		results[i] = result
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return nil, errs
	}

	return results, nil
}

// NewConditionalInt64 returns a calculation which returns the result of ifTrue
// or ifFalse, depending on wether boolCalc returns true or false. If boolCalc
// returns an error, that error is returned instead.
func NewConditionalInt64(boolCalc CalculationBool, ifTrue, ifFalse CalculationInt64) CalculationInt64 {
	return conditionalInt64{
		boolCalc: boolCalc,
		ifTrue:   ifTrue,
		ifFalse:  ifFalse,
	}
}

type conditionalInt64 struct {
	boolCalc CalculationBool
	ifTrue   CalculationInt64
	ifFalse  CalculationInt64
}

func (cond conditionalInt64) CalculateInt64() (int64, error) {
	b, err := cond.boolCalc.CalculateBool()
	if err != nil {
		return 0, err
	}
	if b {
		return cond.ifTrue.CalculateInt64()
	}
	return cond.ifFalse.CalculateInt64()
}

// FailingCalculationInt64 creates a calculation which always fails. This is
// useful for testing when adding custom calculations.
func FailingCalculationInt64(err error) CalculationInt64Func {
	return func() (int64, error) {
		return 0, err
	}
}
