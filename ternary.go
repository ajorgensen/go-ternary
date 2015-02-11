package ternary

func If(conditional bool, trueValue interface{}, falseValue interface{}) interface{} {
	if conditional {
		return trueValue
	} else {
		return falseValue
	}
}
