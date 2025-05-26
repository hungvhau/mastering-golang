// Package functions contains tests for function concepts
package functions

import (
	"fmt"
	"strings"
	"testing"
)

// TestBasicFunction verifies that BasicFunction executes without errors
func TestBasicFunction(t *testing.T) {
	// This is a simple test to ensure the function runs
	// In a real scenario, we might capture stdout to verify output
	BasicFunction() // Should not panic
}

// TestFunctionWithReturn verifies function return values
func TestFunctionWithReturn(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 5, 3, 8},
		{"negative numbers", -5, -3, -8},
		{"mixed numbers", -5, 10, 5},
		{"with zero", 0, 5, 5},
		{"both zero", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FunctionWithReturn(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("FunctionWithReturn(%d, %d) = %d; want %d",
					tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestFunctionWithMultipleReturns verifies functions returning multiple values
func TestFunctionWithMultipleReturns(t *testing.T) {
	tests := []struct {
		name      string
		dividend  int
		divisor   int
		expQuot   int
		expRem    int
	}{
		{"normal division", 17, 5, 3, 2},
		{"exact division", 20, 4, 5, 0},
		{"dividend smaller", 3, 5, 0, 3},
		{"division by zero", 10, 0, 0, 0}, // Special case
		{"negative dividend", -17, 5, -3, -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quotient, remainder := FunctionWithMultipleReturns(tt.dividend, tt.divisor)
			if quotient != tt.expQuot || remainder != tt.expRem {
				t.Errorf("FunctionWithMultipleReturns(%d, %d) = (%d, %d); want (%d, %d)",
					tt.dividend, tt.divisor, quotient, remainder, tt.expQuot, tt.expRem)
			}
		})
	}
}

// TestFunctionWithNamedReturns verifies named return values
func TestFunctionWithNamedReturns(t *testing.T) {
	tests := []struct {
		name           string
		radius         float64
		expectedArea   float64
		expectedCircum float64
		tolerance      float64
	}{
		{"unit circle", 1.0, 3.14159265359, 6.28318530718, 0.00001},
		{"radius 5", 5.0, 78.53981633975, 31.4159265359, 0.00001},
		{"radius 0", 0.0, 0.0, 0.0, 0.00001},
		{"large radius", 100.0, 31415.9265359, 628.318530718, 0.0001},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			area, circumference := FunctionWithNamedReturns(tt.radius)
			
			// Check area within tolerance
			areaDiff := area - tt.expectedArea
			if areaDiff < 0 {
				areaDiff = -areaDiff
			}
			if areaDiff > tt.tolerance {
				t.Errorf("Area for radius %.2f = %.10f; want %.10f (tolerance %.10f)",
					tt.radius, area, tt.expectedArea, tt.tolerance)
			}
			
			// Check circumference within tolerance
			circumDiff := circumference - tt.expectedCircum
			if circumDiff < 0 {
				circumDiff = -circumDiff
			}
			if circumDiff > tt.tolerance {
				t.Errorf("Circumference for radius %.2f = %.10f; want %.10f (tolerance %.10f)",
					tt.radius, circumference, tt.expectedCircum, tt.tolerance)
			}
		})
	}
}

// TestCalculateWithError verifies error handling in functions
func TestCalculateWithError(t *testing.T) {
	tests := []struct {
		name      string
		a, b      int
		operation string
		expected  int
		wantErr   bool
		errMsg    string
	}{
		{"addition", 10, 5, "add", 15, false, ""},
		{"subtraction", 10, 5, "subtract", 5, false, ""},
		{"multiplication", 10, 5, "multiply", 50, false, ""},
		{"division", 10, 5, "divide", 2, false, ""},
		{"division by zero", 10, 0, "divide", 0, true, "division by zero"},
		{"unknown operation", 10, 5, "modulo", 0, true, "unknown operation: modulo"},
		{"empty operation", 10, 5, "", 0, true, "unknown operation: "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := CalculateWithError(tt.a, tt.b, tt.operation)
			
			// Check error expectation
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateWithError() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			
			// If we expect an error, check the error message
			if tt.wantErr && err != nil {
				if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("Error message = %q, want to contain %q", err.Error(), tt.errMsg)
				}
			}
			
			// Check result if no error expected
			if !tt.wantErr && result != tt.expected {
				t.Errorf("CalculateWithError() = %d, want %d", result, tt.expected)
			}
		})
	}
}

// TestHigherOrderFunction verifies functions as parameters
func TestHigherOrderFunction(t *testing.T) {
	tests := []struct {
		name      string
		numbers   []int
		operation func(int) int
		expected  []int
	}{
		{
			name:      "double",
			numbers:   []int{1, 2, 3, 4, 5},
			operation: func(n int) int { return n * 2 },
			expected:  []int{2, 4, 6, 8, 10},
		},
		{
			name:      "square",
			numbers:   []int{1, 2, 3, 4, 5},
			operation: func(n int) int { return n * n },
			expected:  []int{1, 4, 9, 16, 25},
		},
		{
			name:      "add one",
			numbers:   []int{0, 1, 2, 3, 4},
			operation: func(n int) int { return n + 1 },
			expected:  []int{1, 2, 3, 4, 5},
		},
		{
			name:      "negative",
			numbers:   []int{1, -2, 3, -4, 5},
			operation: func(n int) int { return -n },
			expected:  []int{-1, 2, -3, 4, -5},
		},
		{
			name:      "empty slice",
			numbers:   []int{},
			operation: func(n int) int { return n * 2 },
			expected:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HigherOrderFunction(tt.numbers, tt.operation)
			
			// Check length
			if len(result) != len(tt.expected) {
				t.Errorf("Result length = %d, want %d", len(result), len(tt.expected))
				return
			}
			
			// Check each element
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Result[%d] = %d, want %d", i, result[i], tt.expected[i])
				}
			}
		})
	}
}

// TestFunctionReturningFunction verifies closure creation
func TestFunctionReturningFunction(t *testing.T) {
	tests := []struct {
		name       string
		multiplier int
		input      int
		expected   int
	}{
		{"multiply by 2", 2, 5, 10},
		{"multiply by 0", 0, 5, 0},
		{"multiply by -1", -1, 5, -5},
		{"multiply by 10", 10, 7, 70},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := FunctionReturningFunction(tt.multiplier)
			result := fn(tt.input)
			if result != tt.expected {
				t.Errorf("Function with multiplier %d: fn(%d) = %d, want %d",
					tt.multiplier, tt.input, result, tt.expected)
			}
		})
	}
}

// TestClosure verifies stateful functions
func TestClosure(t *testing.T) {
	// Test that each closure maintains independent state
	counter1 := Closure()
	counter2 := Closure()
	
	// Test counter1
	for i := 1; i <= 3; i++ {
		result := counter1()
		if result != i {
			t.Errorf("counter1() call %d = %d, want %d", i, result, i)
		}
	}
	
	// Test counter2 (should start from 1, not continue from counter1)
	for i := 1; i <= 2; i++ {
		result := counter2()
		if result != i {
			t.Errorf("counter2() call %d = %d, want %d", i, result, i)
		}
	}
	
	// Verify counter1 continues from where it left off
	result := counter1()
	if result != 4 {
		t.Errorf("counter1() after counter2 calls = %d, want 4", result)
	}
}

// TestRecursiveFactorial verifies recursive function
func TestRecursiveFactorial(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
		{7, 5040},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("factorial(%d)", tt.n), func(t *testing.T) {
			result := RecursiveFactorial(tt.n)
			if result != tt.expected {
				t.Errorf("RecursiveFactorial(%d) = %d, want %d", tt.n, result, tt.expected)
			}
		})
	}
}

// TestCalculatorMethods verifies struct methods
func TestCalculatorMethods(t *testing.T) {
	t.Run("value receiver", func(t *testing.T) {
		calc := Calculator{LastResult: 0}
		
		// Add with value receiver shouldn't modify the original
		result := calc.Add(5, 3)
		if result != 8 {
			t.Errorf("Add(5, 3) = %d, want 8", result)
		}
		if calc.LastResult != 0 {
			t.Errorf("LastResult after Add = %d, want 0 (unchanged)", calc.LastResult)
		}
	})
	
	t.Run("pointer receiver", func(t *testing.T) {
		calc := Calculator{LastResult: 0}
		
		// AddAndStore with pointer receiver should modify the original
		result := calc.AddAndStore(10, 7)
		if result != 17 {
			t.Errorf("AddAndStore(10, 7) = %d, want 17", result)
		}
		if calc.LastResult != 17 {
			t.Errorf("LastResult after AddAndStore = %d, want 17", calc.LastResult)
		}
		
		// Verify GetLastResult
		if calc.GetLastResult() != 17 {
			t.Errorf("GetLastResult() = %d, want 17", calc.GetLastResult())
		}
	})
}

// TestDeferredExecution verifies defer behavior
func TestDeferredExecution(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		{"normal file", "test.txt", false},
		{"data file", "data.txt", false},
		{"error file", "error.txt", true},
		{"another error", "error_log.txt", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := DeferredExecution(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeferredExecution(%q) error = %v, wantErr %v", 
					tt.filename, err, tt.wantErr)
			}
		})
	}
}

// Benchmark example for performance testing
func BenchmarkRecursiveFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RecursiveFactorial(10)
	}
}

func BenchmarkHigherOrderFunction(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	operation := func(n int) int { return n * n }
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HigherOrderFunction(numbers, operation)
	}
} 