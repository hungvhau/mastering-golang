// Package collections demonstrates Go's built-in collection types:
// arrays, slices, and maps. These are fundamental data structures
// that every Go programmer must master.
package collections

import (
	"fmt"
	"sort"
)

// ArrayBasics demonstrates declaring and using arrays in Go
// Arrays have a fixed size that's part of their type
func ArrayBasics() {
	// Arrays are fixed-size sequences of elements of the same type
	// The size is part of the type, so [5]int and [10]int are different types
	
	// Method 1: Declare and zero-initialize
	var numbers [5]int
	fmt.Println("  Zero-initialized array:", numbers) // [0 0 0 0 0]
	
	// Method 2: Declare and initialize with values
	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println("  Initialized array:", primes)
	
	// Method 3: Let compiler count the elements
	colors := [...]string{"red", "green", "blue"}
	fmt.Printf("  Auto-sized array (len=%d): %v\n", len(colors), colors)
	
	// Method 4: Initialize specific indices
	sparse := [10]int{1: 10, 5: 50, 9: 90}
	fmt.Println("  Sparse array:", sparse)
	
	// Accessing elements (0-indexed)
	fmt.Printf("  First prime: %d, Last prime: %d\n", primes[0], primes[len(primes)-1])
	
	// Modifying elements
	numbers[2] = 42
	fmt.Println("  After modification:", numbers)
	
	// Arrays are value types (copied when assigned)
	copyOfPrimes := primes
	copyOfPrimes[0] = 100
	fmt.Println("  Original primes:", primes)      // Unchanged
	fmt.Println("  Copy of primes:", copyOfPrimes) // Modified
	
	// Iterating over arrays
	fmt.Print("  Squares: ")
	for i, v := range primes {
		fmt.Printf("%d:%d ", i, v*v)
	}
	fmt.Println()
}

// ArrayOperations demonstrates common operations with arrays
func ArrayOperations() {
	// Multi-dimensional arrays
	var matrix [3][3]int
	// Initialize as identity matrix
	for i := 0; i < 3; i++ {
		matrix[i][i] = 1
	}
	fmt.Println("  Identity matrix:")
	for _, row := range matrix {
		fmt.Printf("    %v\n", row)
	}
	
	// Array comparison (arrays are comparable if elements are comparable)
	a1 := [3]int{1, 2, 3}
	a2 := [3]int{1, 2, 3}
	a3 := [3]int{3, 2, 1}
	fmt.Printf("  a1 == a2: %v\n", a1 == a2) // true
	fmt.Printf("  a1 == a3: %v\n", a1 == a3) // false
	
	// Passing arrays to functions (by value - entire array is copied!)
	bigArray := [1000000]int{} // 1 million integers
	bigArray[0] = 42
	fmt.Printf("  Large array first element before: %d\n", bigArray[0])
	modifyArrayCopy(bigArray) // This won't change the original
	fmt.Printf("  Large array first element after: %d\n", bigArray[0])
	
	// Using array pointers for efficiency
	modifyArrayPointer(&bigArray)
	fmt.Printf("  Large array first element after pointer modification: %d\n", bigArray[0])
}

// Helper function that receives array by value
func modifyArrayCopy(arr [1000000]int) {
	arr[0] = 999 // This modifies the copy, not the original
}

// Helper function that receives array by pointer
func modifyArrayPointer(arr *[1000000]int) {
	arr[0] = 999 // This modifies the original array
}

// SliceBasics demonstrates declaring and using slices
// Slices are dynamic, flexible views into arrays
func SliceBasics() {
	// Slices are references to underlying arrays
	// They have length and capacity
	
	// Method 1: Declare a nil slice
	var nilSlice []int
	fmt.Printf("  Nil slice: %v, len=%d, cap=%d, is nil: %v\n", 
		nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
	
	// Method 2: Make a slice with make()
	numbers := make([]int, 5)     // length 5, capacity 5
	fmt.Printf("  Made slice: %v, len=%d, cap=%d\n", numbers, len(numbers), cap(numbers))
	
	// Method 3: Make with different length and capacity
	reserved := make([]int, 3, 10) // length 3, capacity 10
	fmt.Printf("  Reserved slice: %v, len=%d, cap=%d\n", reserved, len(reserved), cap(reserved))
	
	// Method 4: Slice literal
	fruits := []string{"apple", "banana", "cherry"}
	fmt.Printf("  Fruit slice: %v\n", fruits)
	
	// Method 5: Slice from array
	primeArray := [...]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	primeSlice := primeArray[1:4] // Elements at index 1, 2, 3 (not 4!)
	fmt.Printf("  Prime slice [1:4]: %v\n", primeSlice)
	
	// Slice expressions
	fmt.Printf("  primeArray[:5]: %v\n", primeArray[:5])   // First 5 elements
	fmt.Printf("  primeArray[5:]: %v\n", primeArray[5:])   // From index 5 to end
	fmt.Printf("  primeArray[:]: %v\n", primeArray[:])     // All elements
	
	// Modifying slice modifies underlying array
	primeSlice[0] = 999
	fmt.Printf("  After modifying slice: array=%v, slice=%v\n", primeArray, primeSlice)
}

// SliceOperations demonstrates common slice operations
func SliceOperations() {
	// Appending to slices
	var dynamic []int
	fmt.Println("  Building a dynamic slice:")
	
	// Append single elements
	dynamic = append(dynamic, 1)
	dynamic = append(dynamic, 2, 3, 4)
	fmt.Printf("    After appending: %v, len=%d, cap=%d\n", 
		dynamic, len(dynamic), cap(dynamic))
	
	// Append another slice (note the ... operator)
	more := []int{5, 6, 7}
	dynamic = append(dynamic, more...)
	fmt.Printf("    After appending slice: %v\n", dynamic)
	
	// Capacity growth demonstration
	fmt.Println("\n  Capacity growth pattern:")
	growth := make([]int, 0)
	prevCap := cap(growth)
	for i := 0; i < 20; i++ {
		growth = append(growth, i)
		if cap(growth) != prevCap {
			fmt.Printf("    len=%d, new cap=%d (was %d)\n", 
				len(growth), cap(growth), prevCap)
			prevCap = cap(growth)
		}
	}
	
	// Copying slices
	original := []int{1, 2, 3, 4, 5}
	
	// Method 1: Using copy()
	copySlice := make([]int, len(original))
	n := copy(copySlice, original)
	fmt.Printf("\n  Copied %d elements: %v\n", n, copySlice)
	
	// Method 2: Using append (creates new backing array)
	cloneSlice := append([]int(nil), original...)
	fmt.Printf("  Cloned slice: %v\n", cloneSlice)
	
	// Slicing slices (reslicing)
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("\n  Original: %v\n", numbers)
	
	// Various slice operations
	evens := numbers[0:10:10] // slice[low:high:max] - controls capacity
	fmt.Printf("  Full slice with capacity limit: %v, cap=%d\n", evens, cap(evens))
	
	middle := numbers[3:7]
	fmt.Printf("  Middle portion: %v\n", middle)
	
	// Removing elements (by creating new slice)
	// Remove element at index 5
	indexToRemove := 5
	removed := append(numbers[:indexToRemove], numbers[indexToRemove+1:]...)
	fmt.Printf("  After removing index %d: %v\n", indexToRemove, removed)
	
	// Inserting elements
	indexToInsert := 3
	valueToInsert := 999
	inserted := append(numbers[:indexToInsert], 
		append([]int{valueToInsert}, numbers[indexToInsert:]...)...)
	fmt.Printf("  After inserting %d at index %d: %v\n", 
		valueToInsert, indexToInsert, inserted)
}

// SlicePatterns demonstrates common slice patterns and tricks
func SlicePatterns() {
	// Stack operations using slice
	fmt.Println("  Stack operations:")
	stack := []string{}
	
	// Push
	stack = append(stack, "first")
	stack = append(stack, "second")
	stack = append(stack, "third")
	fmt.Printf("    Stack after pushes: %v\n", stack)
	
	// Pop
	if len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("    Popped: %s, Stack now: %v\n", top, stack)
	}
	
	// Queue operations
	fmt.Println("\n  Queue operations:")
	queue := []string{"a", "b", "c", "d"}
	
	// Dequeue (remove from front - less efficient)
	if len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		fmt.Printf("    Dequeued: %s, Queue now: %v\n", front, queue)
	}
	
	// Filter pattern
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := []int{}
	for _, n := range numbers {
		if n%2 == 0 {
			evens = append(evens, n)
		}
	}
	fmt.Printf("\n  Filtered evens: %v\n", evens)
	
	// In-place filtering (more efficient)
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := 0
	for _, x := range nums {
		if x%2 == 1 { // Keep odd numbers
			nums[n] = x
			n++
		}
	}
	nums = nums[:n]
	fmt.Printf("  In-place filtered odds: %v\n", nums)
	
	// Reversing a slice
	toReverse := []int{1, 2, 3, 4, 5}
	for i, j := 0, len(toReverse)-1; i < j; i, j = i+1, j-1 {
		toReverse[i], toReverse[j] = toReverse[j], toReverse[i]
	}
	fmt.Printf("\n  Reversed slice: %v\n", toReverse)
}

// MapBasics demonstrates declaring and using maps
// Maps are Go's built-in hash table implementation
func MapBasics() {
	// Maps are reference types that store key-value pairs
	
	// Method 1: Declare a nil map
	var nilMap map[string]int
	fmt.Printf("  Nil map: %v, len=%d, is nil: %v\n", 
		nilMap, len(nilMap), nilMap == nil)
	
	// Method 2: Make a map
	ages := make(map[string]int)
	ages["Alice"] = 30
	ages["Bob"] = 25
	fmt.Printf("  Ages map: %v\n", ages)
	
	// Method 3: Map literal
	scores := map[string]int{
		"Alice":   95,
		"Bob":     87,
		"Charlie": 92,
	}
	fmt.Printf("  Scores map: %v\n", scores)
	
	// Accessing values
	aliceScore := scores["Alice"]
	fmt.Printf("  Alice's score: %d\n", aliceScore)
	
	// Non-existent key returns zero value
	daveScore := scores["Dave"]
	fmt.Printf("  Dave's score (not in map): %d\n", daveScore)
	
	// Two-value assignment to check existence
	score, exists := scores["Eve"]
	fmt.Printf("  Eve's score: %d, exists: %v\n", score, exists)
	
	// Checking existence without needing the value
	if _, ok := scores["Charlie"]; ok {
		fmt.Println("  Charlie is in the map")
	}
	
	// Adding and updating
	scores["Dave"] = 88      // Add new
	scores["Alice"] = 97     // Update existing
	fmt.Printf("  After add/update: %v\n", scores)
	
	// Deleting from map
	delete(scores, "Bob")
	fmt.Printf("  After deleting Bob: %v\n", scores)
	
	// Iterating over maps (order is not guaranteed!)
	fmt.Println("  Iterating over map:")
	for name, score := range scores {
		fmt.Printf("    %s: %d\n", name, score)
	}
}

// MapOperations demonstrates advanced map operations
func MapOperations() {
	// Maps with different key types
	// Any comparable type can be a key
	
	// Integer keys
	fibonacci := map[int]int{
		0: 0, 1: 1, 2: 1, 3: 2, 4: 3, 5: 5,
	}
	fmt.Printf("  Fibonacci map: %v\n", fibonacci)
	
	// Struct keys
	type Point struct {
		X, Y int
	}
	
	grid := map[Point]string{
		{0, 0}:   "origin",
		{1, 0}:   "east",
		{0, 1}:   "north",
		{-1, 0}:  "west",
		{0, -1}:  "south",
	}
	fmt.Printf("  Grid map: %v\n", grid)
	
	// Map of slices (common pattern)
	studentCourses := map[string][]string{
		"Alice":   {"Math", "Physics", "Chemistry"},
		"Bob":     {"English", "History"},
		"Charlie": {"Computer Science", "Math"},
	}
	fmt.Println("\n  Student courses:")
	for student, courses := range studentCourses {
		fmt.Printf("    %s: %v\n", student, courses)
	}
	
	// Adding to slice in map
	studentCourses["Bob"] = append(studentCourses["Bob"], "Art")
	fmt.Printf("  Bob's courses after adding: %v\n", studentCourses["Bob"])
	
	// Map as a set (using bool values)
	seen := make(map[string]bool)
	words := []string{"hello", "world", "hello", "go", "world", "go", "hello"}
	
	fmt.Println("\n  Using map as set to find unique words:")
	for _, word := range words {
		if !seen[word] {
			seen[word] = true
			fmt.Printf("    First occurrence of: %s\n", word)
		}
	}
	
	// Counting with maps
	letterCount := make(map[rune]int)
	text := "hello golang"
	
	for _, char := range text {
		if char != ' ' {
			letterCount[char]++
		}
	}
	
	fmt.Println("\n  Letter frequency:")
	// Sort keys for consistent output
	var letters []rune
	for letter := range letterCount {
		letters = append(letters, letter)
	}
	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})
	
	for _, letter := range letters {
		fmt.Printf("    %c: %d\n", letter, letterCount[letter])
	}
}

// MapPatterns demonstrates common map patterns and idioms
func MapPatterns() {
	// Pattern 1: Grouping with maps
	people := []struct {
		Name string
		Age  int
		City string
	}{
		{"Alice", 25, "New York"},
		{"Bob", 30, "London"},
		{"Charlie", 25, "New York"},
		{"Dave", 30, "London"},
		{"Eve", 35, "Paris"},
	}
	
	// Group by city
	byCity := make(map[string][]string)
	for _, person := range people {
		byCity[person.City] = append(byCity[person.City], person.Name)
	}
	
	fmt.Println("  People grouped by city:")
	for city, names := range byCity {
		fmt.Printf("    %s: %v\n", city, names)
	}
	
	// Pattern 2: Cache/Memoization
	fmt.Println("\n  Fibonacci with memoization:")
	cache := make(map[int]int)
	var fib func(int) int
	
	fib = func(n int) int {
		if n <= 1 {
			return n
		}
		
		// Check cache
		if val, ok := cache[n]; ok {
			fmt.Printf("    Cache hit for fib(%d) = %d\n", n, val)
			return val
		}
		
		// Calculate and cache
		result := fib(n-1) + fib(n-2)
		cache[n] = result
		fmt.Printf("    Calculated fib(%d) = %d\n", n, result)
		return result
	}
	
	fmt.Printf("  fib(6) = %d\n", fib(6))
	fmt.Printf("  fib(7) = %d (uses cached values)\n", fib(7))
	
	// Pattern 3: Default values
	config := map[string]string{
		"host": "localhost",
		"port": "8080",
	}
	
	// Helper function for default values
	getConfig := func(key, defaultValue string) string {
		if val, ok := config[key]; ok {
			return val
		}
		return defaultValue
	}
	
	fmt.Println("\n  Configuration with defaults:")
	fmt.Printf("    host: %s\n", getConfig("host", "0.0.0.0"))
	fmt.Printf("    port: %s\n", getConfig("port", "3000"))
	fmt.Printf("    timeout: %s\n", getConfig("timeout", "30s"))
	
	// Pattern 4: Two-level map (nested maps)
	// Useful for representing tables or matrices
	userPermissions := make(map[string]map[string]bool)
	
	// Initialize nested maps
	userPermissions["alice"] = map[string]bool{
		"read":   true,
		"write":  true,
		"delete": false,
	}
	
	userPermissions["bob"] = map[string]bool{
		"read":  true,
		"write": false,
	}
	
	// Safe access to nested maps
	hasPermission := func(user, action string) bool {
		if perms, ok := userPermissions[user]; ok {
			return perms[action] // returns false if action not in map
		}
		return false
	}
	
	fmt.Println("\n  Permission checks:")
	fmt.Printf("    alice can write: %v\n", hasPermission("alice", "write"))
	fmt.Printf("    bob can delete: %v\n", hasPermission("bob", "delete"))
	fmt.Printf("    charlie can read: %v\n", hasPermission("charlie", "read"))
}

// CollectionComparison shows when to use arrays vs slices vs maps
func CollectionComparison() {
	fmt.Println("  When to use each collection type:")
	fmt.Println()
	
	fmt.Println("  Arrays - Use when:")
	fmt.Println("    • Size is known at compile time and won't change")
	fmt.Println("    • You need a value type (arrays are copied)")
	fmt.Println("    • Working with fixed-size data (e.g., RGB colors [3]byte)")
	fmt.Println("    • Performance is critical and size is small")
	
	fmt.Println("\n  Slices - Use when:")
	fmt.Println("    • Size might change or is unknown at compile time")
	fmt.Println("    • You need to pass collections to functions efficiently")
	fmt.Println("    • Working with subsets of data (slicing)")
	fmt.Println("    • Building collections dynamically")
	fmt.Println("    • This is the most common choice for sequences")
	
	fmt.Println("\n  Maps - Use when:")
	fmt.Println("    • You need key-value associations")
	fmt.Println("    • Fast lookup by key is important (O(1) average)")
	fmt.Println("    • Keys are not sequential integers")
	fmt.Println("    • Implementing sets, caches, or lookups")
	fmt.Println("    • Grouping or indexing data")
	
	fmt.Println("\n  Performance characteristics:")
	fmt.Println("    Arrays:  O(1) access, O(n) search, fixed size")
	fmt.Println("    Slices:  O(1) access, O(n) search, O(1) amortized append")
	fmt.Println("    Maps:    O(1) average access/insert/delete, no ordering")
} 