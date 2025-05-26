// Package basics_test - Notice the package name convention
// In Go, test files are placed in the same directory as the code they test
// Test files MUST end with _test.go suffix - this is how Go identifies test files
package basics

// Import necessary packages for testing
import (
	"fmt"      // For printing in example tests
	"strings"  // For string manipulation in our tests
	"testing" // Go's built-in testing package - required for all tests
)

// =============================================================================
// INTRODUCTION TO GO TESTING FOR JUNIOR ENGINEERS
// =============================================================================
//
// Go has a built-in testing framework that's simple yet powerful. Here's what
// you need to know:
//
// 1. Test files must end with _test.go
// 2. Test functions must start with Test (capital T)
// 3. Test functions take exactly one parameter: *testing.T
// 4. Use t.Error() or t.Errorf() to report failures
// 5. Tests run with the command: go test
//
// Basic commands:
// - go test              : Run all tests in current package
// - go test -v           : Verbose output (see each test result)
// - go test -cover       : Show test coverage
// - go test -run TestName: Run specific test
// =============================================================================

// TestConstants verifies that our constants have the expected values
// Test function names should be descriptive: Test<WhatYouAreTesting>
func TestConstants(t *testing.T) {
	// Each test should have clear sections:
	// 1. Arrange (setup)
	// 2. Act (execute)
	// 3. Assert (verify)

	// Testing MaxItems constant
	// t.Errorf() allows formatted error messages like fmt.Printf()
	if MaxItems != 100 {
		t.Errorf("MaxItems should be 100, but got %d", MaxItems)
	}

	// Testing AppName constant
	// Always provide clear error messages that help debugging
	if AppName != "GoBasics" {
		t.Errorf("AppName should be 'GoBasics', but got '%s'", AppName)
	}

	// Testing Version constant
	// For floating point comparisons, be careful about precision
	if Version != 1.0 {
		t.Errorf("Version should be 1.0, but got %f", Version)
	}

	// Testing IsDebug constant
	if IsDebug != false {
		t.Errorf("IsDebug should be false, but got %t", IsDebug)
	}

	// Testing Pi constant with floating point comparison
	// For more precise floating point comparisons, you might want to check
	// if the difference is within an acceptable range (epsilon)
	expectedPi := 3.14159
	epsilon := 0.00001
	if diff := Pi - expectedPi; diff < -epsilon || diff > epsilon {
		t.Errorf("Pi should be approximately %f, but got %f (diff: %f)", 
			expectedPi, Pi, diff)
	}
}

// TestGetTypeInfo demonstrates testing functions that return values
func TestGetTypeInfo(t *testing.T) {
	// Call the function we're testing
	result := GetTypeInfo()

	// Verify the result contains expected type information
	// We use strings.Contains to check if substrings exist
	expectedTypes := []struct {
		name     string
		expected string
	}{
		{"int type", "int: int"},
		{"float64 type", "float64: float64"},
		{"string type", "string: string"},
		{"bool type", "bool: bool"},
	}

	// Loop through test cases - this is a common pattern in Go testing
	for _, tc := range expectedTypes {
		// t.Run creates a subtest - helps organize test output
		t.Run(tc.name, func(t *testing.T) {
			if !strings.Contains(result, tc.expected) {
				t.Errorf("Expected result to contain '%s', but got: %s", 
					tc.expected, result)
			}
		})
	}
}

// TestDemonstrateVariables shows how to test functions with side effects
// Since DemonstrateVariables prints to stdout, we're mainly checking it doesn't panic
func TestDemonstrateVariables(t *testing.T) {
	// Sometimes you just want to ensure a function runs without panicking
	// This is useful for functions that primarily have side effects (like printing)
	
	// Using defer with recover to catch any panics
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemonstrateVariables panicked: %v", r)
		}
	}()

	// Call the function - if it panics, our defer will catch it
	DemonstrateVariables()

	// If we reach here, the function executed successfully
	// For functions that print output, you might want to capture stdout
	// and verify the output, but that's more advanced
}

// TestDemonstrateDataTypes tests another function with side effects
func TestDemonstrateDataTypes(t *testing.T) {
	// Similar approach - ensure the function runs without errors
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemonstrateDataTypes panicked: %v", r)
		}
	}()

	DemonstrateDataTypes()
}

// =============================================================================
// TABLE-DRIVEN TESTS - A GO BEST PRACTICE
// =============================================================================
// Table-driven tests are a common pattern in Go where you define test cases
// in a slice and iterate through them. This makes tests more maintainable
// and easier to extend.

// TestVariableTypes demonstrates table-driven testing
func TestVariableTypes(t *testing.T) {
	// Define test cases as a slice of structs
	testCases := []struct {
		name     string      // Test case name for clear output
		value    interface{} // Input value (interface{} can hold any type)
		expected string      // Expected type as string
	}{
		{
			name:     "integer type",
			value:    42,
			expected: "int",
		},
		{
			name:     "string type",
			value:    "hello",
			expected: "string",
		},
		{
			name:     "boolean type",
			value:    true,
			expected: "bool",
		},
		{
			name:     "float64 type",
			value:    3.14,
			expected: "float64",
		},
	}

	// Iterate through test cases
	for _, tc := range testCases {
		// t.Run creates a subtest for each case
		// This provides better test output and allows running specific subtests
		t.Run(tc.name, func(t *testing.T) {
			// Get the type as string using fmt.Sprintf
			gotType := strings.TrimPrefix(
				strings.Split(GetTypeInfo(), ",")[0], 
				"Types - int: ",
			)
			
			// This is a simplified example - in real tests you might
			// have a function that returns the type of a value
			// For demonstration, we're just checking that GetTypeInfo works
			if len(gotType) == 0 {
				t.Errorf("Failed to get type information")
			}
		})
	}
}

// =============================================================================
// BENCHMARK TESTS - MEASURING PERFORMANCE
// =============================================================================
// Benchmark functions must start with Benchmark and take *testing.B parameter
// Run with: go test -bench=.

func BenchmarkGetTypeInfo(b *testing.B) {
	// The testing framework will run this function b.N times
	// b.N is automatically adjusted to get reliable timing
	for i := 0; i < b.N; i++ {
		GetTypeInfo()
	}
}

// =============================================================================
// EXAMPLE TESTS - EXECUTABLE DOCUMENTATION
// =============================================================================
// Example functions serve as documentation and are verified by go test
// They must start with Example

func ExampleGetTypeInfo() {
	// This function serves as documentation
	// The output comment below is verified by go test
	result := GetTypeInfo()
	// For Example tests, we need to use fmt.Println (not println)
	// and the output must match exactly
	fmt.Println(result)
	// Output: Types - int: int, float64: float64, string: string, bool: bool
}

// =============================================================================
// TEST HELPERS - REUSABLE TEST UTILITIES
// =============================================================================
// Helper functions make tests more readable and reduce duplication

// assertStringContains is a helper function for string assertions
// Note: helper functions don't need to start with Test
func assertStringContains(t *testing.T, got, want string) {
	// t.Helper() marks this as a test helper
	// When this fails, the error points to the calling test, not this function
	t.Helper()
	
	if !strings.Contains(got, want) {
		t.Errorf("Expected string to contain '%s', but got: '%s'", want, got)
	}
}

// TestWithHelper demonstrates using test helper functions
func TestWithHelper(t *testing.T) {
	result := GetTypeInfo()
	
	// Using our helper function makes the test more readable
	assertStringContains(t, result, "int: int")
	assertStringContains(t, result, "float64: float64")
	assertStringContains(t, result, "string: string")
	assertStringContains(t, result, "bool: bool")
}

// =============================================================================
// TIPS FOR JUNIOR SOFTWARE ENGINEERS:
// =============================================================================
// 
// 1. **Write tests first (TDD)**: Consider writing tests before implementation
//    - Helps clarify requirements
//    - Ensures your code is testable
//    - Provides immediate feedback
//
// 2. **Keep tests simple and focused**: Each test should verify one thing
//    - Makes failures easier to diagnose
//    - Tests serve as documentation
//
// 3. **Use descriptive test names**: TestFunctionName_WhenCondition_ShouldResult
//    Example: TestCalculateTotal_WhenEmptyCart_ShouldReturnZero
//
// 4. **Follow the AAA pattern**:
//    - Arrange: Set up test data
//    - Act: Execute the function
//    - Assert: Verify the result
//
// 5. **Test edge cases**: Don't just test the happy path
//    - Empty inputs
//    - Nil/null values
//    - Boundary conditions
//    - Error conditions
//
// 6. **Use table-driven tests**: When testing similar scenarios
//    - Reduces code duplication
//    - Easy to add new test cases
//    - Clear test structure
//
// 7. **Mock external dependencies**: Use interfaces for testability
//    - Database calls
//    - API requests
//    - File system operations
//
// 8. **Aim for high coverage, but don't obsess**:
//    - 80% coverage is often good enough
//    - Focus on critical business logic
//    - Some code (like simple getters) may not need tests
//
// 9. **Run tests frequently**:
//    - Before committing code
//    - In CI/CD pipelines
//    - After pulling changes
//
// 10. **Keep tests fast**: Slow tests discourage running them
//     - Mock external services
//     - Use test databases/in-memory stores
//     - Parallelize when possible (t.Parallel())
// ============================================================================= 