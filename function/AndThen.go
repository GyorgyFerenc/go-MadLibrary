package Function

import in "github.com/GyorgyFerenc/go-MadLibrary/interface"

/*
	Chains the function's together.
	Calls the first function with the initial_value parameter,
	after that it uses the returned value to call the next one and so on,
	returning the last result

	Essentialy AndThen(0,a,b,c) is c(b(a(0)))

*/
func AndThen[T any](initial_value T, funcs ...func(T) T) T {
	value := initial_value
	for _, function := range funcs {
		value = function(value)
	}

	return value
}

/*
	Constructs a function from the add_value paramater.
	The new function has 1 paramater, value.
	The new function will add value and add_value together.

	Essentially Adder(a)(b) = a + b
*/
func Adder[T in.Number](add_value T) func(T) T {
	return func(value T) T {
		return value + add_value
	}
}

/*
	Constructs a function from the mul_value paramater.
	The new function has 1 paramater, value.
	The new function will multiply value and mul_value together.

	Essentially Multiplier(a)(b) = a * b
*/
func Multiplier[T in.Number](mul_value T) func(T) T {
	return func(value T) T {
		return value * mul_value
	}
}

/*
	Constructs a function from the sub_value paramater.
	The new function has 1 paramater, value.
	The new function will substract value from sub_value.

	Essentially SubstracterLHS(a)(b) = a - b

	This function has a pair SubstracterRHS which swictches a and b
*/
func SubstracterLHS[T in.Number](sub_value T) func(T) T {
	return func(value T) T {
		return sub_value - value
	}
}

/*
	Constructs a function from the sub_value paramater.
	The new function has 1 paramater, value.
	The new function will substract sub_value from value.

	Essentially SubstracterRHS(a)(b) = b - a

	This function has a pair SubstracterLHS which swictches a and b
*/
func SubstracterRHS[T in.Number](sub_value T) func(T) T {
	return func(value T) T {
		return value - sub_value
	}
}

/*
	Constructs a function from the div_value paramater.
	The new function has 1 paramater, value.
	The new function will substract value from div_value.

	Essentially DividerLHS(a)(b) = a / b

	The user should be carefull to not divide by 0

	This function has a pair DividerRHS which swictches a and b
*/
func DividerLHS[T in.Number](div_value T) func(T) T {
	return func(value T) T {
		return div_value / value
	}
}

/*
	Constructs a function from the div_value paramater.
	The new function has 1 paramater, value.
	The new function will substract div_value from value.

	Essentially DividerRHS(a)(b) = b / a

	The user should be carefull to not divide by 0

	This function has a pair DividerLHS which swictches a and b
*/
func DividerRHS[T in.Number](div_value T) func(T) T {
	return func(value T) T {
		return value / div_value
	}
}
