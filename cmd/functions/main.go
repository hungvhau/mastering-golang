// Package main demonstrates Golang function concepts
// This is an executable program showcasing various function features
package main

import (
	"fmt"
	
	// Import our functions package
	"github.com/hungvhau/mastering-golang/functions"
)

// main function is the entry point for functions demonstration
func main() {
	// Print a header
	fmt.Println("=== Mastering Golang: Functions and Scope ===")
	fmt.Println()

	// Basic function call
	fmt.Println("1. Basic Function:")
	functions.BasicFunction()
	fmt.Println()

	// Function with parameters
	fmt.Println("2. Function with Parameters:")
	functions.FunctionWithParameters("Alice", 30, 165.5, 58.2)
	fmt.Println()

	// Function with return value
	fmt.Println("3. Function with Return:")
	sum := functions.FunctionWithReturn(15, 25)
	fmt.Printf("Sum of 15 and 25 = %d\n", sum)
	fmt.Println()

	// Function with multiple returns
	fmt.Println("4. Function with Multiple Returns:")
	quotient, remainder := functions.FunctionWithMultipleReturns(17, 5)
	fmt.Printf("17 divided by 5: quotient = %d, remainder = %d\n", quotient, remainder)
	fmt.Println()

	// Function with named returns
	fmt.Println("5. Function with Named Returns:")
	area, circumference := functions.FunctionWithNamedReturns(5.0)
	fmt.Printf("Circle with radius 5.0: area = %.2f, circumference = %.2f\n", area, circumference)
	fmt.Println()

	// Function with error handling
	fmt.Println("6. Function with Error Handling:")
	result1, err := functions.CalculateWithError(10, 5, "add")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 + 5 = %d\n", result1)
	}

	// Test error case
	result2, err := functions.CalculateWithError(10, 0, "divide")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %d\n", result2)
	}
	fmt.Println()

	// Demonstrate variable scope
	fmt.Println("7. Variable Scope Demonstration:")
	functions.DemonstrateScope()
	fmt.Println()

	// Demonstrate anonymous functions
	fmt.Println("8. Anonymous Functions:")
	functions.DemonstrateAnonymousFunctions()
	fmt.Println()

	// Demonstrate higher-order functions
	fmt.Println("9. Higher-Order Functions:")
	numbers := []int{1, 2, 3, 4, 5}
	
	// Using anonymous function
	tripled := functions.HigherOrderFunction(numbers, func(n int) int {
		return n * 3
	})
	fmt.Printf("Triple %v = %v\n", numbers, tripled)
	
	// Using function returning function
	doubler := functions.FunctionReturningFunction(2)
	doubled := functions.HigherOrderFunction(numbers, doubler)
	fmt.Printf("Double %v = %v\n", numbers, doubled)
	fmt.Println()

	// Demonstrate closures
	fmt.Println("10. Closures:")
	counter := functions.Closure()
	fmt.Printf("Counter call 1: %d\n", counter())
	fmt.Printf("Counter call 2: %d\n", counter())
	fmt.Printf("Counter call 3: %d\n", counter())
	
	// Create another counter to show independence
	counter2 := functions.Closure()
	fmt.Printf("New counter call 1: %d\n", counter2())
	fmt.Printf("Original counter call 4: %d\n", counter())
	fmt.Println()

	// Demonstrate variadic functions
	fmt.Println("11. Variadic Functions:")
	functions.VariadicFunction("Numbers", 1, 2, 3, 4, 5)
	functions.VariadicFunction("More numbers", 10, 20, 30)
	functions.VariadicFunction("Single number", 42)
	functions.VariadicFunction("No numbers")
	fmt.Println()

	// Demonstrate deferred execution
	fmt.Println("12. Deferred Execution:")
	fileErr := functions.DeferredExecution("data.txt")
	if fileErr != nil {
		fmt.Printf("Error: %v\n", fileErr)
	}
	fmt.Println()

	// Test deferred execution with error
	fmt.Println("Deferred Execution with Error:")
	fileErr = functions.DeferredExecution("error.txt")
	if fileErr != nil {
		fmt.Printf("Error: %v\n", fileErr)
	}
	fmt.Println()

	// Demonstrate recursion
	fmt.Println("13. Recursive Function:")
	for i := 0; i <= 5; i++ {
		factorial := functions.RecursiveFactorial(i)
		fmt.Printf("Factorial of %d = %d\n", i, factorial)
	}
	fmt.Println()

	// Demonstrate methods
	fmt.Println("14. Methods Demonstration:")
	functions.DemonstrateMethods()

	fmt.Println("\n=== Functions Demo Complete ===")
} 