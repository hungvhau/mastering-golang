// Package main demonstrates Golang loops and control flow concepts
// This is an executable program showcasing Go's versatile for loop
package main

import (
	"fmt"
	"time"
	
	// Import our loops package
	"github.com/hungvhau/mastering-golang/loops"
)

// main function is the entry point for loops demonstration
func main() {
	// Print a header with timestamp
	fmt.Println("=== Mastering Golang: Loops and Control Flow ===")
	fmt.Printf("Started at: %s\n", time.Now().Format("15:04:05"))
	fmt.Println()

	// 1. Basic for loop demonstration
	fmt.Println("1. Basic For Loop:")
	fmt.Println("The traditional for loop with initialization, condition, and post statement")
	loops.BasicForLoop()
	fmt.Println()

	// 2. While-style loop demonstration
	fmt.Println("2. While-Style Loop:")
	fmt.Println("Go's for loop can mimic while loops from other languages")
	loops.WhileStyleLoop()
	fmt.Println()

	// 3. Infinite loop demonstration
	fmt.Println("3. Infinite Loop:")
	fmt.Println("Loops without conditions run forever unless explicitly broken")
	loops.InfiniteLoop()
	fmt.Println()

	// 4. Break and Continue demonstration
	fmt.Println("4. Break and Continue:")
	fmt.Println("Control flow statements to exit or skip iterations")
	loops.BreakAndContinue()
	fmt.Println()

	// 5. Range over arrays and slices
	fmt.Println("5. Range over Arrays and Slices:")
	fmt.Println("The range keyword provides a clean way to iterate over collections")
	loops.RangeOverArraySlice()
	fmt.Println()

	// 6. Range over maps
	fmt.Println("6. Range over Maps:")
	fmt.Println("Iterating over key-value pairs in maps (unordered)")
	loops.RangeOverMap()
	fmt.Println()

	// 7. Range over strings
	fmt.Println("7. Range over Strings:")
	fmt.Println("String iteration works with Unicode code points (runes)")
	loops.RangeOverString()
	fmt.Println()

	// 8. Range over channels
	fmt.Println("8. Range over Channels:")
	fmt.Println("Receiving values from channels until closed")
	loops.RangeOverChannel()
	fmt.Println()

	// 9. Nested loops
	fmt.Println("9. Nested Loops:")
	fmt.Println("Loops within loops for multi-dimensional operations")
	loops.NestedLoops()
	fmt.Println()

	// 10. Labeled break and continue
	fmt.Println("10. Labeled Break and Continue:")
	fmt.Println("Using labels to control outer loops from inner loops")
	loops.LabeledBreakContinue()
	fmt.Println()

	// 11. Common loop patterns
	fmt.Println("11. Common Loop Patterns:")
	fmt.Println("Practical patterns you'll use frequently in Go")
	loops.LoopPatterns()
	fmt.Println()

	// 12. Loop performance considerations
	fmt.Println("12. Loop Performance Considerations:")
	fmt.Println("Tips for writing efficient loops")
	loops.LoopPerformanceConsiderations()
	fmt.Println()

	// Additional examples in main
	fmt.Println("13. Additional Examples:")
	
	// Example: Do-while equivalent
	fmt.Println("\n  Do-While Equivalent:")
	fmt.Println("  Go doesn't have do-while, but we can simulate it")
	attemptCount := 0
	for {
		attemptCount++
		fmt.Printf("    Attempt %d\n", attemptCount)
		
		// Condition checked at the end (do-while behavior)
		if attemptCount >= 3 {
			fmt.Println("    Maximum attempts reached")
			break
		}
	}
	
	// Example: Countdown timer
	fmt.Println("\n  Countdown Timer:")
	for i := 5; i > 0; i-- {
		fmt.Printf("    %d...\n", i)
		time.Sleep(200 * time.Millisecond) // Small delay for effect
	}
	fmt.Println("    Liftoff! 🚀")
	
	// Example: Fibonacci sequence
	fmt.Println("\n  Fibonacci Sequence (first 10 numbers):")
	a, b := 0, 1
	for i := 0; i < 10; i++ {
		fmt.Printf("    F(%d) = %d\n", i, a)
		a, b = b, a+b // Parallel assignment
	}
	
	// Example: Password validation loop
	fmt.Println("\n  Password Validation Example:")
	passwords := []string{"abc", "password123", "SecureP@ss1", "12345"}
	
	for _, pwd := range passwords {
		fmt.Printf("    Checking '%s': ", pwd)
		
		// Multiple validation checks
		hasUpper := false
		hasLower := false
		hasDigit := false
		hasSpecial := false
		
		// Check each character
		for _, char := range pwd {
			switch {
			case char >= 'A' && char <= 'Z':
				hasUpper = true
			case char >= 'a' && char <= 'z':
				hasLower = true
			case char >= '0' && char <= '9':
				hasDigit = true
			case char == '@' || char == '!' || char == '#' || char == '$':
				hasSpecial = true
			}
		}
		
		// Determine if password is strong
		if len(pwd) >= 8 && hasUpper && hasLower && hasDigit && hasSpecial {
			fmt.Println("✓ Strong")
		} else {
			fmt.Println("✗ Weak")
		}
	}

	fmt.Println("\n=== Loops Demo Complete ===")
	fmt.Printf("Finished at: %s\n", time.Now().Format("15:04:05"))
} 