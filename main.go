// Package main provides an overview of the Mastering Golang project
// This file serves as an entry point and guide to the learning modules
package main

import (
	"fmt"
)

// main function provides instructions on how to use this learning project
func main() {
	fmt.Println("=== Welcome to Mastering Golang ===")
	fmt.Println()
	fmt.Println("This project is organized into learning modules:")
	fmt.Println()
	fmt.Println("1. BASICS MODULE - Fundamental concepts")
	fmt.Println("   Topics: Variables, Data Types, Operators, Conditionals")
	fmt.Println("   Run: go run cmd/basics/main.go")
	fmt.Println()
	fmt.Println("2. FUNCTIONS MODULE - Function concepts and patterns")
	fmt.Println("   Topics: Function declarations, Parameters, Returns, Scope,")
	fmt.Println("           Anonymous functions, Closures, Higher-order functions,")
	fmt.Println("           Variadic functions, Defer, Recursion, Methods")
	fmt.Println("   Run: go run cmd/functions/main.go")
	fmt.Println()
	fmt.Println("PROJECT STRUCTURE:")
	fmt.Println("├── main.go                 # This file - project overview")
	fmt.Println("├── go.mod                  # Module definition")
	fmt.Println("├── README.md               # Detailed documentation")
	fmt.Println("├── basics/                 # Basic concepts package")
	fmt.Println("│   ├── variables.go       # Variables and constants")
	fmt.Println("│   ├── operators.go       # Arithmetic operators")
	fmt.Println("│   ├── conditionals.go    # Control flow")
	fmt.Println("│   └── variables_test.go  # Unit tests")
	fmt.Println("├── functions/              # Functions package")
	fmt.Println("│   ├── functions.go       # All function concepts")
	fmt.Println("│   └── functions_test.go  # Function tests")
	fmt.Println("└── cmd/                    # Executable programs")
	fmt.Println("    ├── basics/")
	fmt.Println("    │   └── main.go        # Demo runner for basics")
	fmt.Println("    └── functions/")
	fmt.Println("        └── main.go        # Demo runner for functions")
	fmt.Println()
	fmt.Println("To run tests:")
	fmt.Println("  go test ./...            # Run all tests")
	fmt.Println("  go test -v ./basics      # Run basics tests with verbose output")
	fmt.Println("  go test -v ./functions   # Run functions tests with verbose output")
	fmt.Println()
	fmt.Println("Happy learning!")
}
