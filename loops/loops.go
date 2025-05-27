// Package loops contains examples demonstrating Go loops and control flow concepts
// Go has only one looping construct: the for loop, but it's incredibly versatile
package loops

import (
	"fmt"
	"strings"
	"time"
)

// BasicForLoop demonstrates the traditional C-style for loop
// This is the most common form with initialization, condition, and post statement
func BasicForLoop() {
	fmt.Println("Basic For Loop:")
	
	// Traditional for loop with three components:
	// 1. Initialization (i := 0) - executed once before the loop starts
	// 2. Condition (i < 5) - checked before each iteration
	// 3. Post statement (i++) - executed after each iteration
	for i := 0; i < 5; i++ {
		fmt.Printf("  Iteration %d\n", i)
	}
	
	// Multiple variables can be initialized and incremented
	fmt.Println("\nFor loop with multiple variables:")
	for i, j := 0, 10; i < 5; i, j = i+1, j-2 {
		fmt.Printf("  i=%d, j=%d\n", i, j)
	}
}

// WhileStyleLoop shows how Go's for loop can act like a while loop
// By omitting the init and post statements, we get while-like behavior
func WhileStyleLoop() {
	fmt.Println("While-Style Loop:")
	
	// Only the condition is specified - similar to while loops in other languages
	count := 0
	for count < 5 {
		fmt.Printf("  Count is %d\n", count)
		count++ // Don't forget to update the condition variable!
	}
	
	// Another example with more complex condition
	fmt.Println("\nWhile loop with string building:")
	word := ""
	for len(word) < 10 {
		word += "a"
		fmt.Printf("  Word is now: '%s' (length: %d)\n", word, len(word))
	}
}

// InfiniteLoop demonstrates infinite loops and how to exit them
// Infinite loops are created by omitting the condition entirely
func InfiniteLoop() {
	fmt.Println("Infinite Loop with Break:")
	
	// Infinite loop - will run forever unless broken
	counter := 0
	for {
		// Always have an exit condition in infinite loops!
		if counter >= 5 {
			fmt.Println("  Breaking out of infinite loop")
			break // Exit the loop immediately
		}
		fmt.Printf("  Infinite loop iteration %d\n", counter)
		counter++
	}
	
	// Practical example: event loop waiting for condition
	fmt.Println("\nSimulated event loop:")
	attempts := 0
	for {
		attempts++
		// Simulate checking for an event
		if attempts > 3 {
			fmt.Printf("  Event occurred after %d attempts!\n", attempts)
			break
		}
		fmt.Printf("  Waiting for event... (attempt %d)\n", attempts)
		// In real code, you might sleep or check actual conditions here
	}
}

// BreakAndContinue demonstrates loop control statements
// break exits the loop, continue skips to the next iteration
func BreakAndContinue() {
	fmt.Println("Break Statement:")
	
	// Break example - exit when a condition is met
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Printf("  Breaking at i=%d\n", i)
			break // Exit the loop completely
		}
		fmt.Printf("  i=%d\n", i)
	}
	
	fmt.Println("\nContinue Statement:")
	
	// Continue example - skip certain iterations
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			// Skip even numbers
			continue // Jump to next iteration
		}
		fmt.Printf("  Odd number: %d\n", i)
	}
	
	// More complex example with both break and continue
	fmt.Println("\nCombined Break and Continue:")
	sum := 0
	for i := 1; i <= 20; i++ {
		// Skip multiples of 3
		if i%3 == 0 {
			fmt.Printf("  Skipping %d (multiple of 3)\n", i)
			continue
		}
		
		sum += i
		fmt.Printf("  Added %d, sum is now %d\n", i, sum)
		
		// Break if sum exceeds 50
		if sum > 50 {
			fmt.Printf("  Sum exceeded 50, breaking at i=%d\n", i)
			break
		}
	}
	fmt.Printf("  Final sum: %d\n", sum)
}

// RangeOverArraySlice demonstrates iterating over arrays and slices
// Range provides both index and value for each element
func RangeOverArraySlice() {
	fmt.Println("Range over Array:")
	
	// Fixed-size array
	fruits := [5]string{"apple", "banana", "orange", "grape", "mango"}
	
	// Range with both index and value
	for index, fruit := range fruits {
		fmt.Printf("  fruits[%d] = %s\n", index, fruit)
	}
	
	fmt.Println("\nRange over Slice:")
	
	// Dynamic slice
	numbers := []int{10, 20, 30, 40, 50}
	
	// Range with index only (value ignored with _)
	fmt.Println("  Indices only:")
	for i := range numbers {
		fmt.Printf("    Index: %d\n", i)
	}
	
	// Range with value only (index ignored with _)
	fmt.Println("  Values only:")
	sum := 0
	for _, num := range numbers {
		sum += num
		fmt.Printf("    Value: %d (running sum: %d)\n", num, sum)
	}
	
	// Modifying slice elements during iteration
	fmt.Println("\nModifying slice elements:")
	for i := range numbers {
		numbers[i] *= 2 // Double each element
	}
	fmt.Printf("  Doubled numbers: %v\n", numbers)
}

// RangeOverMap shows how to iterate over map key-value pairs
// Maps are unordered, so iteration order is not guaranteed
func RangeOverMap() {
	fmt.Println("Range over Map:")
	
	// Create a map of string to int
	ages := map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
		"Diana":   28,
	}
	
	// Range over map - returns key and value
	fmt.Println("  All entries:")
	for name, age := range ages {
		fmt.Printf("    %s is %d years old\n", name, age)
	}
	
	// Range with keys only
	fmt.Println("\n  Keys only:")
	for name := range ages {
		fmt.Printf("    Name: %s\n", name)
	}
	
	// Filtering during iteration
	fmt.Println("\n  People over 30:")
	for name, age := range ages {
		if age <= 30 {
			continue // Skip people 30 or younger
		}
		fmt.Printf("    %s (%d years)\n", name, age)
	}
	
	// Safe deletion during iteration
	fmt.Println("\n  Removing people under 30:")
	for name, age := range ages {
		if age < 30 {
			delete(ages, name) // Safe to delete during range
			fmt.Printf("    Removed %s\n", name)
		}
	}
	fmt.Printf("  Remaining: %v\n", ages)
}

// RangeOverString demonstrates iterating over string characters
// Range over string iterates over Unicode code points (runes), not bytes
func RangeOverString() {
	fmt.Println("Range over String:")
	
	// ASCII string
	ascii := "Hello"
	fmt.Printf("  ASCII string: %s\n", ascii)
	for index, char := range ascii {
		// char is a rune (int32), representing a Unicode code point
		fmt.Printf("    [%d] = %c (Unicode: U+%04X)\n", index, char, char)
	}
	
	// Unicode string with emojis and non-ASCII characters
	unicode := "Hello, 世界! 🌍"
	fmt.Printf("\n  Unicode string: %s\n", unicode)
	for index, char := range unicode {
		// Notice index jumps by more than 1 for multi-byte characters
		fmt.Printf("    [%d] = %c (Unicode: U+%04X)\n", index, char, char)
	}
	
	// Counting actual characters vs bytes
	fmt.Printf("\n  String analysis:\n")
	fmt.Printf("    Byte length: %d\n", len(unicode))
	charCount := 0
	for range unicode {
		charCount++
	}
	fmt.Printf("    Character count: %d\n", charCount)
	
	// Building a new string character by character
	fmt.Println("\n  Reversing string (Unicode-safe):")
	runes := []rune{}
	for _, r := range unicode {
		runes = append([]rune{r}, runes...) // Prepend each rune
	}
	reversed := string(runes)
	fmt.Printf("    Original:  %s\n", unicode)
	fmt.Printf("    Reversed:  %s\n", reversed)
}

// RangeOverChannel shows how to iterate over channel values
// Range over channel receives values until the channel is closed
func RangeOverChannel() {
	fmt.Println("Range over Channel:")
	
	// Create a buffered channel
	numbers := make(chan int, 5)
	
	// Send values to channel in a goroutine
	go func() {
		// Send some values
		for i := 1; i <= 5; i++ {
			numbers <- i * 10
			fmt.Printf("  Sent: %d\n", i*10)
		}
		// Important: close the channel when done sending
		close(numbers)
	}()
	
	// Small delay to see the sending messages first
	time.Sleep(100 * time.Millisecond)
	
	// Range over channel - receives until channel is closed
	fmt.Println("  Receiving values:")
	for num := range numbers {
		fmt.Printf("    Received: %d\n", num)
		time.Sleep(50 * time.Millisecond) // Simulate processing
	}
	fmt.Println("  Channel closed, range loop ended")
}

// NestedLoops demonstrates loops within loops
// Common for working with 2D data structures or combinations
func NestedLoops() {
	fmt.Println("Nested Loops:")
	
	// Simple nested loop - multiplication table
	fmt.Println("  Multiplication table (3x3):")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("    %d × %d = %2d  ", i, j, i*j)
		}
		fmt.Println() // New line after each row
	}
	
	// Nested loops with break
	fmt.Println("\n  Breaking from nested loops:")
	for i := 0; i < 5; i++ {
		fmt.Printf("  Outer loop i=%d:\n", i)
		for j := 0; j < 5; j++ {
			if i*j > 6 {
				fmt.Printf("    Breaking inner loop at j=%d (i*j=%d > 6)\n", j, i*j)
				break // Only breaks the inner loop
			}
			fmt.Printf("    Inner loop j=%d (i*j=%d)\n", j, i*j)
		}
	}
	
	// 2D slice traversal
	fmt.Println("\n  2D slice traversal:")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	
	for row := range matrix {
		for col := range matrix[row] {
			fmt.Printf("    matrix[%d][%d] = %d\n", row, col, matrix[row][col])
		}
	}
}

// LabeledBreakContinue shows how to break/continue outer loops
// Labels allow breaking out of nested loops
func LabeledBreakContinue() {
	fmt.Println("Labeled Break:")
	
	// Label the outer loop
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("  i=%d, j=%d\n", i, j)
			if i == 2 && j == 2 {
				fmt.Println("  Breaking outer loop!")
				break outerLoop // Break the outer loop, not just the inner
			}
		}
	}
	fmt.Println("  After labeled break")
	
	fmt.Println("\nLabeled Continue:")
	
	// Label for continue
nextNumber:
	for i := 0; i < 5; i++ {
		// Check if number should be skipped
		for j := 2; j <= 4; j++ {
			if i%j == 0 && i != j {
				fmt.Printf("  Skipping %d (divisible by %d)\n", i, j)
				continue nextNumber // Continue the outer loop
			}
		}
		fmt.Printf("  %d passed all checks\n", i)
	}
}

// LoopPatterns demonstrates common loop patterns and idioms
func LoopPatterns() {
	fmt.Println("Common Loop Patterns:")
	
	// Pattern 1: Filtering
	fmt.Println("\n  1. Filtering pattern:")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := []int{}
	for _, n := range numbers {
		if n%2 == 0 {
			evens = append(evens, n)
		}
	}
	fmt.Printf("    Original: %v\n", numbers)
	fmt.Printf("    Evens: %v\n", evens)
	
	// Pattern 2: Transformation
	fmt.Println("\n  2. Transformation pattern:")
	words := []string{"hello", "world", "golang"}
	upper := make([]string, len(words))
	for i, word := range words {
		upper[i] = strings.ToUpper(word)
	}
	fmt.Printf("    Original: %v\n", words)
	fmt.Printf("    Uppercase: %v\n", upper)
	
	// Pattern 3: Accumulation/Reduction
	fmt.Println("\n  3. Accumulation pattern:")
	values := []int{10, 20, 30, 40, 50}
	sum := 0
	product := 1
	for _, v := range values {
		sum += v
		product *= v
	}
	fmt.Printf("    Values: %v\n", values)
	fmt.Printf("    Sum: %d, Product: %d\n", sum, product)
	
	// Pattern 4: Finding
	fmt.Println("\n  4. Finding pattern:")
	fruits := []string{"apple", "banana", "cherry", "date"}
	target := "cherry"
	found := false
	index := -1
	for i, fruit := range fruits {
		if fruit == target {
			found = true
			index = i
			break // Stop once found
		}
	}
	if found {
		fmt.Printf("    Found '%s' at index %d\n", target, index)
	} else {
		fmt.Printf("    '%s' not found\n", target)
	}
	
	// Pattern 5: Sliding window
	fmt.Println("\n  5. Sliding window pattern:")
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	windowSize := 3
	fmt.Printf("    Data: %v\n", data)
	fmt.Printf("    Windows of size %d:\n", windowSize)
	for i := 0; i <= len(data)-windowSize; i++ {
		window := data[i : i+windowSize]
		windowSum := 0
		for _, v := range window {
			windowSum += v
		}
		fmt.Printf("      %v -> sum: %d\n", window, windowSum)
	}
}

// LoopPerformanceConsiderations demonstrates performance aspects of loops
func LoopPerformanceConsiderations() {
	fmt.Println("Loop Performance Considerations:")
	
	// Pre-allocating slices
	fmt.Println("\n  1. Pre-allocation vs dynamic growth:")
	size := 5
	
	// Dynamic growth (less efficient)
	dynamic := []int{}
	for i := 0; i < size; i++ {
		dynamic = append(dynamic, i*10) // Slice may need to be reallocated
	}
	fmt.Printf("    Dynamic result: %v\n", dynamic)
	
	// Pre-allocated (more efficient)
	prealloc := make([]int, size)
	for i := 0; i < size; i++ {
		prealloc[i] = i * 10 // No reallocation needed
	}
	fmt.Printf("    Pre-allocated result: %v\n", prealloc)
	
	// Length caching
	fmt.Println("\n  2. Length caching:")
	text := "Hello, Golang!"
	
	// Without caching (len called each iteration)
	count1 := 0
	for i := 0; i < len(text); i++ {
		if text[i] == 'o' {
			count1++
		}
	}
	
	// With caching (len called once)
	count2 := 0
	length := len(text) // Cache the length
	for i := 0; i < length; i++ {
		if text[i] == 'o' {
			count2++
		}
	}
	fmt.Printf("    Found %d 'o' characters in '%s'\n", count1, text)
	
	// Early termination
	fmt.Println("\n  3. Early termination optimization:")
	numbers := []int{1, 3, 5, 7, 9, 11, 13, 15}
	threshold := 10
	
	// Find first number greater than threshold
	var result int
	for _, n := range numbers {
		fmt.Printf("    Checking %d...\n", n)
		if n > threshold {
			result = n
			fmt.Printf("    Found %d > %d, stopping early\n", n, threshold)
			break // Don't check remaining elements
		}
	}
	
	// Display the result
	if result > 0 {
		fmt.Printf("    Result: First number > %d is %d\n", threshold, result)
	} else {
		fmt.Printf("    No number found greater than %d\n", threshold)
	}
} 