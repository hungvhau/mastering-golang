// Package functions contains examples demonstrating Go functions and related concepts
package functions

import (
	"fmt"
	"strings"
)

// Global variables - accessible throughout the package
// In Go, variables declared outside of functions have package-level scope
var (
	// Exported variable (starts with capital letter) - accessible from other packages
	GlobalCounter int = 0
	
	// Unexported variable (starts with lowercase) - only accessible within this package
	globalMessage string = "This is a package-level variable"
)

// Constants at package level - these are also global but immutable
const (
	// Mathematical constant accessible throughout the package
	GoldenRatio float64 = 1.618033988749895
	
	// String constant for demonstration
	defaultGreeting = "Hello from Go!"
)

// BasicFunction demonstrates the simplest function structure
// Functions are declared with the 'func' keyword
// This function takes no parameters and returns nothing
func BasicFunction() {
	// Local variable - only accessible within this function
	localVar := "I'm local to BasicFunction"
	fmt.Println("BasicFunction called:", localVar)
}

// FunctionWithParameters shows how to pass arguments to functions
// Parameters are declared with name followed by type
// Multiple parameters of the same type can be declared together
func FunctionWithParameters(name string, age int, height, weight float64) {
	// Parameters are local to the function scope
	fmt.Printf("%s is %d years old, %.1f cm tall, and weighs %.1f kg\n", 
		name, age, height, weight)
}

// FunctionWithReturn demonstrates a function that returns a value
// The return type is specified after the parameter list
func FunctionWithReturn(a, b int) int {
	// Local variable to store result
	sum := a + b
	// Return statement sends the value back to the caller
	return sum
}

// FunctionWithMultipleReturns shows Go's ability to return multiple values
// This is a unique and powerful feature of Go
func FunctionWithMultipleReturns(dividend, divisor int) (quotient int, remainder int) {
	// Check for division by zero
	if divisor == 0 {
		// Return zero values for error case
		return 0, 0
	}
	// Calculate and return both values
	quotient = dividend / divisor
	remainder = dividend % divisor
	return quotient, remainder
}

// FunctionWithNamedReturns demonstrates named return values
// Named returns create variables that can be used in the function
func FunctionWithNamedReturns(radius float64) (area, circumference float64) {
	// Pi constant from math package would normally be used
	const pi = 3.14159265359
	
	// Named return values are already declared
	area = pi * radius * radius
	circumference = 2 * pi * radius
	
	// 'return' without values returns the named values
	return
}

// CalculateWithError shows the Go idiom of returning an error
// This is the standard way to handle errors in Go
func CalculateWithError(a, b int, operation string) (int, error) {
	// Local variable for result
	var result int
	
	// Switch on operation type
	switch operation {
	case "add":
		result = a + b
	case "subtract":
		result = a - b
	case "multiply":
		result = a * b
	case "divide":
		if b == 0 {
			// Return zero value and error
			return 0, fmt.Errorf("division by zero")
		}
		result = a / b
	default:
		// Return error for unknown operation
		return 0, fmt.Errorf("unknown operation: %s", operation)
	}
	
	// Return result and nil error (no error)
	return result, nil
}

// DemonstrateScope shows different variable scopes in Go
func DemonstrateScope() {
	// Access global variable
	fmt.Printf("Global message: %s\n", globalMessage)
	fmt.Printf("Global counter before: %d\n", GlobalCounter)
	
	// Modify global variable
	GlobalCounter++
	fmt.Printf("Global counter after increment: %d\n", GlobalCounter)
	
	// Local variable with same name as global (shadowing)
	globalMessage := "I'm shadowing the global variable"
	fmt.Printf("Local message (shadowing): %s\n", globalMessage)
	
	// Block scope demonstration
	{
		// Variable only exists within this block
		blockVar := "I only exist in this block"
		fmt.Printf("Block variable: %s\n", blockVar)
		
		// Can still access outer scope
		fmt.Printf("Can access local message: %s\n", globalMessage)
	}
	// blockVar is not accessible here - would cause compile error
	
	// Loop scope
	for i := 0; i < 3; i++ {
		// i is local to the for loop
		loopVar := fmt.Sprintf("Loop iteration %d", i)
		fmt.Printf("%s\n", loopVar)
	}
	// i and loopVar are not accessible here
}

// HigherOrderFunction demonstrates functions as first-class citizens
// This function takes another function as a parameter
func HigherOrderFunction(numbers []int, operation func(int) int) []int {
	// Create a new slice for results
	results := make([]int, len(numbers))
	
	// Apply the operation to each number
	for i, num := range numbers {
		results[i] = operation(num)
	}
	
	return results
}

// FunctionReturningFunction shows how functions can return other functions
// This creates a closure that captures the multiplier value
func FunctionReturningFunction(multiplier int) func(int) int {
	// Return an anonymous function that captures 'multiplier'
	return func(x int) int {
		// This function has access to 'multiplier' from outer scope
		return x * multiplier
	}
}

// DemonstrateAnonymousFunctions shows different ways to use anonymous functions
func DemonstrateAnonymousFunctions() {
	// Anonymous function assigned to a variable
	square := func(x int) int {
		return x * x
	}
	fmt.Printf("Square of 5: %d\n", square(5))
	
	// Immediately invoked anonymous function
	result := func(a, b int) int {
		return a + b
	}(10, 20) // Note the (10, 20) - this invokes the function immediately
	fmt.Printf("Immediate result: %d\n", result)
	
	// Anonymous function with closure
	counter := 0
	increment := func() int {
		// This function captures and modifies the outer 'counter' variable
		counter++
		return counter
	}
	
	fmt.Printf("Counter: %d\n", increment()) // 1
	fmt.Printf("Counter: %d\n", increment()) // 2
	fmt.Printf("Counter: %d\n", increment()) // 3
	
	// Using anonymous function as argument
	numbers := []int{1, 2, 3, 4, 5}
	doubled := HigherOrderFunction(numbers, func(n int) int {
		return n * 2
	})
	fmt.Printf("Doubled numbers: %v\n", doubled)
}

// Closure creates a function that captures external variables
// This is a powerful feature for creating stateful functions
func Closure() func() int {
	// Local variable that will be captured
	count := 0
	
	// Return a function that captures 'count'
	return func() int {
		// Each call modifies and returns the captured variable
		count++
		return count
	}
}

// VariadicFunction demonstrates functions with variable number of arguments
// The ... syntax allows accepting any number of arguments of the same type
func VariadicFunction(prefix string, values ...int) {
	// values is treated as a slice of int inside the function
	fmt.Printf("%s: ", prefix)
	
	// Calculate sum of all values
	sum := 0
	for _, v := range values {
		sum += v
	}
	
	fmt.Printf("Sum of %v = %d\n", values, sum)
}

// DeferredExecution demonstrates the defer statement
// Deferred functions execute after the surrounding function returns
func DeferredExecution(filename string) error {
	fmt.Printf("Opening file: %s\n", filename)
	
	// Defer statements execute in LIFO order (last in, first out)
	defer fmt.Println("3. Final cleanup")
	defer fmt.Println("2. Closing file")
	defer fmt.Println("1. First deferred call")
	
	// Simulate some work
	fmt.Println("Processing file...")
	
	// Deferred functions run even if there's an error
	if strings.Contains(filename, "error") {
		return fmt.Errorf("simulated error in file processing")
	}
	
	fmt.Println("File processed successfully")
	return nil
}

// RecursiveFactorial demonstrates recursion in Go
// Recursive functions call themselves with modified parameters
func RecursiveFactorial(n int) int {
	// Base case - stop recursion
	if n <= 1 {
		return 1
	}
	
	// Recursive case - function calls itself
	// Each call reduces n by 1, moving toward the base case
	return n * RecursiveFactorial(n-1)
}

// Calculator demonstrates methods in Go
// Methods are functions with a receiver
type Calculator struct {
	// Struct field to store state
	LastResult int
}

// Add method with value receiver - receives a copy of the struct
func (c Calculator) Add(a, b int) int {
	result := a + b
	// This won't affect the original struct (it's a copy)
	c.LastResult = result
	return result
}

// AddAndStore method with pointer receiver - can modify the struct
func (c *Calculator) AddAndStore(a, b int) int {
	result := a + b
	// This will modify the original struct
	c.LastResult = result
	return result
}

// GetLastResult is a getter method
func (c Calculator) GetLastResult() int {
	return c.LastResult
}

// DemonstrateMethods shows how methods work in Go
func DemonstrateMethods() {
	// Create a Calculator instance
	calc := Calculator{LastResult: 0}
	
	// Call method with value receiver
	sum1 := calc.Add(5, 3)
	fmt.Printf("Add(5, 3) = %d, LastResult = %d\n", sum1, calc.LastResult)
	
	// Call method with pointer receiver
	sum2 := calc.AddAndStore(10, 7)
	fmt.Printf("AddAndStore(10, 7) = %d, LastResult = %d\n", sum2, calc.LastResult)
} 