# Common Tasks Reference

## 🎯 Adding New Concepts

### Add a New Topic to Existing Module
1. **Identify the module**: basics/, functions/, loops/, or collections/
2. **Add to existing file** or create new file if topic is large
3. **Follow documentation pattern**:
   ```go
   // FunctionName demonstrates [concept being taught]
   // This function shows how to [specific learning objective]
   func FunctionName() {
       // Step-by-step explanation
       code := "with inline comments"
   }
   ```
4. **Update the demo runner** in corresponding cmd/ directory
5. **Add tests** if in basics/ or functions/ module

### Create a New Learning Module
1. **Create module directory**: e.g., `structs/`
2. **Create main teaching file**: e.g., `structs/structs.go`
3. **Create test file**: e.g., `structs/structs_test.go`
4. **Create demo runner**: `cmd/05-structs/main.go`
5. **Update**:
   - `main.go` with new module info
   - `README.md` with module description
   - `.ai/context.json` with module metadata

### Working with Collections Examples
1. **Arrays**: Show fixed vs dynamic sizing needs
   ```go
   // Fixed size when size is known
   rgb := [3]byte{255, 0, 0}  // Red color
   
   // Show value semantics
   copy := rgb  // Full copy made
   ```

2. **Slices**: Demonstrate growth and operations
   ```go
   // Show capacity growth
   data := make([]int, 0, 2)
   data = append(data, 1, 2, 3) // Capacity doubles
   
   // Common operations
   filtered := data[:0]  // Reuse backing array
   ```

3. **Maps**: Show practical use cases
   ```go
   // Counting pattern
   counts := make(map[string]int)
   counts[word]++
   
   // Caching pattern
   if val, ok := cache[key]; ok {
       return val  // Cache hit
   }
   ```

## 📝 Documentation Standards

### Function Documentation Template
```go
// DemonstrateConceptName shows [what it demonstrates]
// Key concepts covered:
// - Point 1
// - Point 2
// This is useful when [practical application]
func DemonstrateConceptName() {
    // Inline comments explain each step
}
```

### Test Documentation Template
```go
func TestConceptName(t *testing.T) {
    // Arrange - setup test data
    
    // Act - perform the operation
    
    // Assert - verify results
}
```

## 🔧 Common Modifications

### Add Example to Existing Function
- Locate the function in the appropriate module
- Add new example with explanatory comments
- Ensure it builds on existing examples

### Improve Error Handling Examples
- Add to `functions/functions.go` in error handling section
- Show both error creation and handling
- Demonstrate error wrapping

### Add Performance Examples
- Add to relevant module (likely loops/)
- Include benchmarks in test files
- Comment on why one approach is faster

## 🚀 Running and Testing

### Quick Commands
```bash
# Run all demos
go run main.go                    # See overview
go run cmd/01-basics/main.go     # Run basics
go run cmd/02-functions/main.go  # Run functions
go run cmd/03-loops/main.go      # Run loops
go run cmd/04-collections/main.go # Run collections

# Run tests
go test ./...                     # All tests
go test -v ./basics              # Specific module
go test -bench=. ./...           # Run benchmarks
```

### Adding New Tests
1. Place in `*_test.go` file in same directory
2. Follow naming: `Test[FunctionName][Scenario]`
3. Use table-driven tests for multiple cases
4. Include both positive and negative test cases

## ⚡ Performance Considerations

For educational code:
- **Clarity > Performance** - Always prioritize understanding
- **Show both ways** - Simple version first, optimized second
- **Explain tradeoffs** - Comment why optimization matters
- **Benchmark** - Prove performance claims with benchmarks

## 🎨 Style Guidelines

1. **Function Length**: Keep demo functions focused on one concept
2. **Comments**: Every non-obvious line should have explanation
3. **Variable Names**: Descriptive names over short names
4. **Error Messages**: Educational - explain what went wrong
5. **Output**: Use fmt.Printf with clear labels

## 🔍 Debugging Educational Code

When something doesn't work as expected:
1. Check the demo runner output
2. Verify imports are correct
3. Ensure examples are self-contained
4. Test in isolation with `go run`
5. Add debug prints to show variable states 