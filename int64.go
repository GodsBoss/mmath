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
	return NewReduceLeft(
		func(left, right int64) (int64, error) {
			return left + right, nil
		},
		NewConstantInt64(0),
		calculations,
	)
}

// NewProductInt64 returns a calculation which returns the product of all
// calculations passed to it. If one or more calculations fail, an error
// wrapping all those individual errors is returned.
func NewProductInt64(calculations ...CalculationInt64) CalculationInt64 {
	return NewReduceLeft(
		func(left, right int64) (int64, error) {
			return left * right, nil
		},
		NewConstantInt64(1),
		calculations,
	)
}

func reduceLeftInt64(
	reduce func(leftOp, rightOp int64) (int64, error),
	initialValue int64,
	calculations []CalculationInt64,
) (int64, error) {

	values, err := runCalculationsInt64(calculations...)
	if err != nil {
		return 0, err
	}

	result := initialValue
	for i := range values {
		var err error
		result, err = reduce(result, values[i])
		if err != nil {
			return 0, err
		}
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

// NewCreateBinaryInt64 wraps a simple binary arithmetic function (int64, int64) -> int64 and
// returns a calculation constructor representing the same calculation.
func NewCreateBinaryInt64(
	f func(left, right int64) int64,
) func(left, right CalculationInt64) CalculationInt64Func {
	return func(left, right CalculationInt64) CalculationInt64Func {
		return func() (int64, error) {
			var errs errors

			leftValue, err := left.CalculateInt64()
			if err != nil {
				errs = append(errs, err)
			}

			rightValue, err := right.CalculateInt64()
			if err != nil {
				errs = append(errs, err)
			}

			if len(errs) > 0 {
				return 0, errs
			}

			return f(leftValue, rightValue), nil
		}
	}
}

// NewSignumInt64 returns a calculation which returns the signum of another
// calculation. If that other calculation fails, that error is returned instead.
func NewSignumInt64(calculation CalculationInt64) CalculationInt64Func {
	return func() (int64, error) {
		v, err := calculation.CalculateInt64()
		if err != nil {
			return 0, err
		}

		if v > 0 {
			return 1, nil
		}
		if v < 0 {
			return -1, nil
		}
		return 0, nil
	}
}

// NewReduceLeft takes a function which reduces two values to one (or an error),
// an initial value and a number of calculations. The beginning value is the
// initial value, then for every calculated value is combined via reduce to
// create the final value which is returned.
//
// If creating the initial value fails, that error is returned. If one or more
// of the calculations fail, their combined errors are returned (as an error).
// Should reduce fail at any point, that error is returned.
func NewReduceLeft(
	reduce func(current int64, next int64) (int64, error),
	initialValue CalculationInt64,
	calculations []CalculationInt64,
) CalculationInt64 {
	return reduceLeft{
		reduce:       reduce,
		initialValue: initialValue,
		calculations: calculations,
	}
}

type reduceLeft struct {
	reduce       func(current int64, next int64) (int64, error)
	initialValue CalculationInt64
	calculations []CalculationInt64
}

func (rl reduceLeft) CalculateInt64() (int64, error) {
	initialValue, err := rl.initialValue.CalculateInt64()
	if err != nil {
		return 0, err
	}
	return reduceLeftInt64(rl.reduce, initialValue, rl.calculations)
}
