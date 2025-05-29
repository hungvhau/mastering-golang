// Package main demonstrates Golang collections: arrays, slices, and maps
// This is an executable program showcasing Go's fundamental data structures
package main

import (
	"fmt"
	"time"
	
	// Import our collections package
	"github.com/hungvhau/mastering-golang/collections"
)

// main function is the entry point for collections demonstration
func main() {
	// Print a header with timestamp
	fmt.Println("=== Mastering Golang: Collections (Arrays, Slices, and Maps) ===")
	fmt.Printf("Started at: %s\n", time.Now().Format("15:04:05"))
	fmt.Println()

	// 1. Arrays - Fixed-size collections
	fmt.Println("1. Arrays - Fixed-Size Collections:")
	fmt.Println("Arrays have a fixed size that's part of their type")
	collections.ArrayBasics()
	fmt.Println()

	// 2. Array Operations
	fmt.Println("2. Array Operations:")
	fmt.Println("Working with multi-dimensional arrays and understanding value semantics")
	collections.ArrayOperations()
	fmt.Println()

	// 3. Slices - Dynamic arrays
	fmt.Println("3. Slices - Dynamic Arrays:")
	fmt.Println("Slices are flexible, dynamic views into arrays")
	collections.SliceBasics()
	fmt.Println()

	// 4. Slice Operations
	fmt.Println("4. Slice Operations:")
	fmt.Println("Common operations: append, copy, insert, delete")
	collections.SliceOperations()
	fmt.Println()

	// 5. Slice Patterns
	fmt.Println("5. Slice Patterns:")
	fmt.Println("Common patterns: stacks, queues, filtering")
	collections.SlicePatterns()
	fmt.Println()

	// 6. Maps - Key-Value Pairs
	fmt.Println("6. Maps - Key-Value Pairs:")
	fmt.Println("Hash tables for fast lookups and associations")
	collections.MapBasics()
	fmt.Println()

	// 7. Map Operations
	fmt.Println("7. Map Operations:")
	fmt.Println("Advanced map usage with different key types")
	collections.MapOperations()
	fmt.Println()

	// 8. Map Patterns
	fmt.Println("8. Map Patterns:")
	fmt.Println("Common patterns: grouping, caching, defaults")
	collections.MapPatterns()
	fmt.Println()

	// 9. Collection Comparison
	fmt.Println("9. When to Use Each Collection Type:")
	collections.CollectionComparison()
	fmt.Println()

	// Additional examples in main
	fmt.Println("10. Real-World Examples:")
	
	// Example: Word frequency counter
	fmt.Println("\n  Word Frequency Counter:")
	text := "the quick brown fox jumps over the lazy dog the fox"
	wordCount := make(map[string]int)
	
	// Simple word splitting (in real code, use strings.Fields)
	word := ""
	for _, char := range text + " " {
		if char == ' ' {
			if word != "" {
				wordCount[word]++
				word = ""
			}
		} else {
			word += string(char)
		}
	}
	
	fmt.Println("  Word frequencies:")
	for word, count := range wordCount {
		if count > 1 {
			fmt.Printf("    '%s': %d times\n", word, count)
		}
	}
	
	// Example: Matrix operations with 2D slices
	fmt.Println("\n  Matrix Operations with 2D Slices:")
	// Create a 3x3 matrix
	matrix := make([][]int, 3)
	for i := range matrix {
		matrix[i] = make([]int, 3)
		for j := range matrix[i] {
			matrix[i][j] = i*3 + j + 1
		}
	}
	
	fmt.Println("  Original matrix:")
	for _, row := range matrix {
		fmt.Printf("    %v\n", row)
	}
	
	// Transpose the matrix
	transposed := make([][]int, 3)
	for i := range transposed {
		transposed[i] = make([]int, 3)
		for j := range transposed[i] {
			transposed[i][j] = matrix[j][i]
		}
	}
	
	fmt.Println("  Transposed matrix:")
	for _, row := range transposed {
		fmt.Printf("    %v\n", row)
	}
	
	// Example: Simple inventory system
	fmt.Println("\n  Inventory System Example:")
	type Item struct {
		Name     string
		Quantity int
		Price    float64
	}
	
	inventory := map[string]Item{
		"apple":  {"Red Apple", 50, 0.50},
		"banana": {"Yellow Banana", 30, 0.30},
		"orange": {"Fresh Orange", 20, 0.75},
	}
	
	// Add items
	inventory["grape"] = Item{"Purple Grape", 40, 2.00}
	
	// Update quantity
	if item, exists := inventory["apple"]; exists {
		item.Quantity += 10
		inventory["apple"] = item
	}
	
	// Calculate total value
	totalValue := 0.0
	fmt.Println("  Current inventory:")
	for id, item := range inventory {
		value := float64(item.Quantity) * item.Price
		totalValue += value
		fmt.Printf("    %s: %s (qty: %d, price: $%.2f, value: $%.2f)\n",
			id, item.Name, item.Quantity, item.Price, value)
	}
	fmt.Printf("  Total inventory value: $%.2f\n", totalValue)
	
	// Example: Implementing a simple LRU cache concept
	fmt.Println("\n  Simple Cache Example:")
	cache := make(map[string]string)
	cacheOrder := []string{} // Track order for LRU
	maxSize := 3
	
	// Helper function to add to cache
	addToCache := func(key, value string) {
		// If key exists, remove from order
		newOrder := []string{}
		for _, k := range cacheOrder {
			if k != key {
				newOrder = append(newOrder, k)
			}
		}
		cacheOrder = newOrder
		
		// Add to end (most recently used)
		cacheOrder = append(cacheOrder, key)
		cache[key] = value
		
		// Remove oldest if over capacity
		if len(cacheOrder) > maxSize {
			oldest := cacheOrder[0]
			cacheOrder = cacheOrder[1:]
			delete(cache, oldest)
			fmt.Printf("    Evicted: %s\n", oldest)
		}
		
		fmt.Printf("    Cached: %s = %s (order: %v)\n", key, value, cacheOrder)
	}
	
	// Simulate cache usage
	addToCache("user1", "Alice")
	addToCache("user2", "Bob")
	addToCache("user3", "Charlie")
	addToCache("user4", "Dave")     // This will evict user1
	addToCache("user2", "Bob Jr.")  // This updates user2 and moves it to end
	
	fmt.Println("\n  Final cache state:")
	for _, key := range cacheOrder {
		fmt.Printf("    %s: %s\n", key, cache[key])
	}

	fmt.Println("\n=== Collections Demo Complete ===")
	fmt.Printf("Finished at: %s\n", time.Now().Format("15:04:05"))
	
	// Key takeaways
	fmt.Println("\nKey Takeaways:")
	fmt.Println("• Arrays: Fixed size, value semantics, good for known sizes")
	fmt.Println("• Slices: Dynamic, reference semantics, most versatile for lists")
	fmt.Println("• Maps: Fast key-value lookups, perfect for associations and caching")
	fmt.Println("• Choose the right tool: Arrays for fixed data, Slices for lists, Maps for lookups")
} 