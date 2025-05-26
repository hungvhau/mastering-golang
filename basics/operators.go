// Package basics - Arithmetic operators demonstration
package basics

// Calculate demonstrates basic arithmetic operations
// This function takes two integers and returns their sum
// Parameters:
//   - a: first integer operand
//   - b: second integer operand
// Returns: the sum of a and b
func Calculate(a, b int) int {
	// The + operator performs addition
	// In Go, the return statement sends the value back to the caller
	return a + b
}

// AllOperations demonstrates all basic arithmetic operators
// This function shows how to use +, -, *, /, and % operators
// Parameters:
//   - x: first integer operand
//   - y: second integer operand
// Returns: sum, difference, product, quotient, and remainder
func AllOperations(x, y int) (int, int, int, int, int) {
	// Addition operator (+)
	// Adds two numbers together
	sum := x + y

	// Subtraction operator (-)
	// Subtracts the second number from the first
	difference := x - y

	// Multiplication operator (*)
	// Multiplies two numbers
	product := x * y

	// Division operator (/)
	// Integer division - returns only the whole number part
	// Note: dividing by zero will cause a panic (runtime error)
	quotient := x / y

	// Modulo operator (%)
	// Returns the remainder after division
	// Useful for checking if a number is even/odd or for cyclic operations
	remainder := x % y

	// Go supports multiple return values
	// This is a powerful feature for returning related values together
	return sum, difference, product, quotient, remainder
}

// OperateOnFloats demonstrates arithmetic with floating-point numbers
// Float operations can have decimal results unlike integer operations
func OperateOnFloats(a, b float64) (float64, float64, float64, float64) {
	// Float addition preserves decimal precision
	sum := a + b

	// Float subtraction
	difference := a - b

	// Float multiplication
	product := a * b

	// Float division returns decimal results
	// Unlike integer division, this keeps the fractional part
	quotient := a / b

	return sum, difference, product, quotient
}

// CompoundOperations demonstrates compound assignment operators
// These operators modify a variable in place
func CompoundOperations() (int, int, int, int, int) {
	// Initialize a variable
	num := 10

	// Addition assignment (+=)
	// Equivalent to: num = num + 5
	num += 5
	afterAdd := num

	// Subtraction assignment (-=)
	// Equivalent to: num = num - 3
	num -= 3
	afterSub := num

	// Multiplication assignment (*=)
	// Equivalent to: num = num * 2
	num *= 2
	afterMul := num

	// Division assignment (/=)
	// Equivalent to: num = num / 4
	num /= 4
	afterDiv := num

	// Modulo assignment (%=)
	// Equivalent to: num = num % 3
	num %= 3
	afterMod := num

	return afterAdd, afterSub, afterMul, afterDiv, afterMod
}

// IncrementDecrement demonstrates the ++ and -- operators
// Go only supports postfix increment/decrement (i++, i--)
func IncrementDecrement(start int) (int, int) {
	// Create a copy to demonstrate increment
	inc := start
	// Increment operator (++)
	// Adds 1 to the variable
	inc++ // Equivalent to: inc = inc + 1

	// Create another copy to demonstrate decrement
	dec := start
	// Decrement operator (--)
	// Subtracts 1 from the variable
	dec-- // Equivalent to: dec = dec - 1

	// Note: In Go, ++ and -- are statements, not expressions
	// This means you cannot do: x = y++ or if (i++ > 5)
	// They must be on their own line

	return inc, dec
}

// BitwiseOperations demonstrates bitwise operators
// These work on the binary representation of numbers
func BitwiseOperations(a, b int) (int, int, int, int, int, int) {
	// Bitwise AND (&)
	// Each bit is 1 only if both corresponding bits are 1
	andResult := a & b

	// Bitwise OR (|)
	// Each bit is 1 if at least one corresponding bit is 1
	orResult := a | b

	// Bitwise XOR (^)
	// Each bit is 1 if corresponding bits are different
	xorResult := a ^ b

	// Bitwise NOT (^)
	// When used as unary operator, flips all bits
	notResult := ^a

	// Left shift (<<)
	// Shifts bits to the left, filling with zeros
	// Equivalent to multiplying by 2^n
	leftShift := a << 2 // Multiply by 4 (2^2)

	// Right shift (>>)
	// Shifts bits to the right
	// Equivalent to dividing by 2^n
	rightShift := a >> 1 // Divide by 2

	return andResult, orResult, xorResult, notResult, leftShift, rightShift
}

// ComparisonOperators demonstrates comparison operators
// These return boolean values (true/false)
func ComparisonOperators(x, y int) (bool, bool, bool, bool, bool, bool) {
	// Equal to (==)
	// Returns true if values are equal
	equal := x == y

	// Not equal to (!=)
	// Returns true if values are different
	notEqual := x != y

	// Less than (<)
	// Returns true if x is less than y
	lessThan := x < y

	// Less than or equal to (<=)
	// Returns true if x is less than or equal to y
	lessOrEqual := x <= y

	// Greater than (>)
	// Returns true if x is greater than y
	greaterThan := x > y

	// Greater than or equal to (>=)
	// Returns true if x is greater than or equal to y
	greaterOrEqual := x >= y

	return equal, notEqual, lessThan, lessOrEqual, greaterThan, greaterOrEqual
} 