// Package main demonstrates basic Golang concepts
// This is an executable program showcasing variables, operators, and conditionals
package main

import (
	"fmt"
	
	// Import our basics package with basic concepts
	"github.com/hungvhau/mastering-golang/basics"
)

// main function is the entry point for basic concepts demonstration
func main() {
	// Print a header
	fmt.Println("=== Mastering Golang: Basic Concepts ===")
	fmt.Println()

	// Demonstrate variables and constants
	fmt.Println("1. Variables and Constants Demo:")
	basics.DemonstrateVariables()
	fmt.Println()

	// Demonstrate data types
	fmt.Println("2. Data Types Demo:")
	basics.DemonstrateDataTypes()
	fmt.Println()

	// Demonstrate operators
	fmt.Println("3. Basic Operators Demo:")
	// Call our calculator function with sample values
	result := basics.Calculate(10, 3)
	fmt.Printf("Calculate(10, 3) = %d\n", result)

	// Demonstrate all arithmetic operations
	sum, diff, prod, quot, rem := basics.AllOperations(15, 4)
	fmt.Printf("AllOperations(15, 4):\n")
	fmt.Printf("  Sum: %d, Difference: %d, Product: %d\n", sum, diff, prod)
	fmt.Printf("  Quotient: %d, Remainder: %d\n", quot, rem)
	fmt.Println()

	// Demonstrate conditionals
	fmt.Println("4. Conditionals Demo:")

	// Test age category function
	age := 25
	category := basics.GetAgeCategory(age)
	fmt.Printf("Age %d is categorized as: %s\n", age, category)

	// Test grade evaluation
	score := 85
	grade := basics.EvaluateGrade(score)
	fmt.Printf("Score %d receives grade: %s\n", score, grade)

	// Test day type
	dayNumber := 3
	dayType := basics.GetDayType(dayNumber)
	fmt.Printf("Day %d is a: %s\n", dayNumber, dayType)

	fmt.Println("\n=== Basic Concepts Demo Complete ===")
} 