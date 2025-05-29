# Mastering Golang - Learning Project

This project demonstrates fundamental Golang concepts with detailed comments for learning purposes. The code is organized into modules that teach specific aspects of the Go programming language.

## Project Structure

```
.
├── go.mod                  # Go module definition
├── README.md               # Project documentation (this file)
├── main.go                 # Entry point with project overview
├── basics/                 # Basic concepts package
│   ├── variables.go       # Variables and constants examples
│   ├── operators.go       # Basic operators examples
│   ├── conditionals.go    # Conditional statements examples
│   └── variables_test.go  # Unit tests
├── functions/              # Functions package
│   ├── functions.go       # Comprehensive function concepts
│   └── functions_test.go  # Function tests
├── loops/                  # Loops and control flow package
│   └── loops.go           # Comprehensive loop concepts
├── collections/            # Collections package
│   └── collections.go     # Arrays, slices, and maps
└── cmd/                    # Executable programs
    ├── 01-basics/
    │   └── main.go        # Demo runner for basics
    ├── 02-functions/
    │   └── main.go        # Demo runner for functions
    ├── 03-loops/
    │   └── main.go        # Demo runner for loops and control flow
    └── 04-collections/
        └── main.go        # Demo runner for collections
```

## Learning Modules

### 1. Basics Module
Located in the `basics/` directory, this module covers fundamental Go concepts:

- **Variables and Constants**: Declaration and usage of `var` and `const`
- **Data Types**: Working with `int`, `string`, `bool`, and `float64`
- **Basic Operators**: Arithmetic operators (+, -, *, /, %)
- **Conditionals**: Control flow with `if`, `else`, and `switch`

### 2. Functions Module
Located in the `functions/` directory, this module covers all aspects of functions in Go:

- **Basic Functions**: Simple function declarations with the `func` keyword
- **Parameters and Returns**: Passing arguments and returning single/multiple values
- **Named Returns**: Functions with pre-declared return variables
- **Error Handling**: The Go idiom of returning errors
- **Variable Scope**: Local vs global variables, shadowing, and block scope
- **Anonymous Functions**: Function literals and immediately invoked functions
- **Higher-Order Functions**: Functions as first-class citizens
- **Closures**: Functions that capture external variables
- **Variadic Functions**: Functions with variable number of arguments
- **Deferred Execution**: Using `defer` for cleanup operations
- **Recursion**: Functions that call themselves
- **Methods**: Functions with receivers attached to types

### 3. Loops Module
Located in the `loops/` directory, this module covers Go's versatile for loop and control flow:

- **Basic For Loop**: Traditional C-style loops with initialization, condition, and post statement
- **While-Style Loops**: Using for loops without init and post statements
- **Infinite Loops**: Creating and breaking out of endless loops
- **Break and Continue**: Controlling loop execution flow
- **Range Keyword**: Iterating over arrays, slices, maps, strings, and channels
- **Nested Loops**: Working with multi-dimensional data structures
- **Labeled Break/Continue**: Controlling outer loops from inner loops
- **Loop Patterns**: Common patterns like filtering, transformation, and accumulation
- **Performance Considerations**: Writing efficient loops with pre-allocation and caching

### 4. Collections Module
Located in the `collections/` directory, this module covers Go's built-in collection types:

- **Arrays**: Fixed-size sequences with value semantics
  - Declaration methods (zero-initialized, literal, auto-sized)
  - Multi-dimensional arrays and array operations
  - Array comparison and performance considerations
- **Slices**: Dynamic arrays with reference semantics
  - Creating slices (nil, make, literal, from arrays)
  - Slice operations (append, copy, insert, delete)
  - Capacity management and growth patterns
  - Common patterns (stacks, queues, filtering)
- **Maps**: Hash tables for key-value associations
  - Declaration and initialization
  - CRUD operations (Create, Read, Update, Delete)
  - Different key types and nested maps
  - Common patterns (caching, counting, grouping, sets)
- **Collection Comparison**: When to use arrays vs slices vs maps

## Running the Examples

### Quick Start
```bash
# Run the overview/guide
go run main.go

# Run basic concepts demo
go run cmd/01-basics/main.go

# Run functions demo
go run cmd/02-functions/main.go

# Run loops and control flow demo
go run cmd/03-loops/main.go

# Run collections demo
go run cmd/04-collections/main.go
```

### Running Tests
```bash
# Run all tests in the project
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests for a specific package
go test -v ./basics
```

## Learning Path

1. **Start with the overview**: Run `go run main.go` to understand the project structure
2. **Learn the basics**: Run `go run cmd/01-basics/main.go` to see fundamental concepts in action
3. **Explore functions**: Run `go run cmd/02-functions/main.go` to understand Go's function features
4. **Master loops**: Run `go run cmd/03-loops/main.go` to learn Go's powerful for loop and control flow
5. **Understand collections**: Run `go run cmd/04-collections/main.go` to work with arrays, slices, and maps
6. **Study the code**: Read through the well-commented source files:
   - Start with `basics/variables.go` for variable declarations
   - Move to `basics/operators.go` for arithmetic operations
   - Study `basics/conditionals.go` for control flow
   - Deep dive into `functions/functions.go` for comprehensive function concepts
   - Master `loops/loops.go` for all loop patterns and control flow techniques
   - Explore `collections/collections.go` for data structure fundamentals
7. **Run the tests**: Use `go test -v ./basics` to see how the functions are tested
8. **Experiment**: Modify the code and create your own examples

## Go Best Practices Demonstrated

### Package Organization
- Clear separation of concerns with `basics` and `functions` packages
- Each package has its own executable `main.go` for demonstrations
- Test files are co-located with the code they test

### Code Quality
- Comprehensive inline documentation explaining concepts
- Meaningful variable and function names
- Consistent formatting following `gofmt` standards
- Proper error handling patterns

### Project Structure
- Module-based organization for easy navigation
- Separate executable programs for different learning topics
- Central entry point (`main.go`) providing guidance
- Clean import paths using the module name

### Testing
- Test-driven development approach
- Comprehensive test coverage for basic concepts
- Clear test function names describing what is being tested

## Requirements

- Go 1.21 or higher
- Basic understanding of command-line operations
- Text editor or IDE with Go support (recommended: VS Code with Go extension)

## Contributing

This is a learning project. Feel free to:
- Add more examples to existing modules
- Create new modules for advanced topics
- Improve documentation and comments
- Add more comprehensive tests

## Next Steps

After mastering these basics, consider exploring:
- Arrays and Slices
- Maps and Structs
- Interfaces
- Goroutines and Channels
- Error handling patterns
- Package management
- Building web services

Happy learning! 