// Package basics contains fundamental Go concepts for learning
// Package names should be short, concise, and lowercase
package basics

// Import the fmt package for printing output
import "fmt"

// Constants are declared using the const keyword
// They must be assigned at compile time and cannot be changed
const (
	// Typed constant - explicitly declaring the type
	MaxItems int = 100

	// Untyped constants - Go infers the type
	AppName = "GoBasics"    // string constant
	Version = 1.0           // float64 constant
	IsDebug = false         // bool constant
)

// Single constant declaration
const Pi float64 = 3.14159

// DemonstrateVariables shows different ways to declare and use variables in Go
// Function names in Go should be CamelCase (exported) or camelCase (unexported)
func DemonstrateVariables() {
	// Method 1: var keyword with explicit type
	// This is the most explicit way to declare a variable
	var age int = 25
	fmt.Printf("Age (explicit type): %d\n", age)

	// Method 2: var keyword with type inference
	// Go can infer the type from the assigned value
	var name = "John Doe" // Go infers this as string
	fmt.Printf("Name (type inference): %s\n", name)

	// Method 3: Short variable declaration using :=
	// This is the most common way inside functions
	city := "New York" // Only works inside functions
	fmt.Printf("City (short declaration): %s\n", city)

	// Multiple variable declaration
	var x, y, z int = 1, 2, 3
	fmt.Printf("Multiple vars: x=%d, y=%d, z=%d\n", x, y, z)

	// Multiple variables with different types
	var (
		firstName string = "Alice"
		lastName  string = "Smith"
		salary    float64 = 75000.50
		isActive  bool   = true
	)
	fmt.Printf("Employee: %s %s, Salary: $%.2f, Active: %t\n", 
		firstName, lastName, salary, isActive)

	// Zero values - variables without initialization get zero values
	var (
		zeroInt    int     // 0
		zeroString string  // "" (empty string)
		zeroBool   bool    // false
		zeroFloat  float64 // 0.0
	)
	fmt.Printf("Zero values - int: %d, string: '%s', bool: %t, float: %f\n",
		zeroInt, zeroString, zeroBool, zeroFloat)

	// Using constants
	fmt.Printf("Constants - App: %s v%.1f, Max Items: %d, Debug: %t\n",
		AppName, Version, MaxItems, IsDebug)
}

// DemonstrateDataTypes shows the basic data types in Go
// This function helps understand Go's type system
func DemonstrateDataTypes() {
	// Integer types
	// int - platform dependent (32 or 64 bit)
	var regularInt int = 42
	// Specific size integers
	var smallInt int8 = 127        // -128 to 127
	var mediumInt int16 = 32767    // -32768 to 32767
	var largeInt int32 = 2147483647 // -2147483648 to 2147483647
	var hugeInt int64 = 9223372036854775807

	fmt.Printf("Integers - int: %d, int8: %d, int16: %d, int32: %d, int64: %d\n",
		regularInt, smallInt, mediumInt, largeInt, hugeInt)

	// Unsigned integers (only positive values)
	var unsignedInt uint = 42
	var byteVal byte = 255 // byte is alias for uint8
	var bigUint uint64 = 18446744073709551615

	fmt.Printf("Unsigned - uint: %d, byte: %d, uint64: %d\n",
		unsignedInt, byteVal, bigUint)

	// Floating point types
	var price float32 = 19.99      // 32-bit floating point
	var precision float64 = 3.141592653589793 // 64-bit floating point (default for decimals)

	fmt.Printf("Floats - float32: %.2f, float64: %.15f\n", price, precision)

	// String type
	var message string = "Hello, Go!"
	var multiline string = `This is a
multiline string
using backticks`

	fmt.Printf("Strings - message: %s\n", message)
	fmt.Printf("Multiline:\n%s\n", multiline)

	// Boolean type
	var isReady bool = true
	var isComplete bool = false

	fmt.Printf("Booleans - isReady: %t, isComplete: %t\n", isReady, isComplete)

	// Type conversion (casting)
	// Go requires explicit type conversion
	var intValue int = 42
	var floatValue float64 = float64(intValue) // Must explicitly convert
	var backToInt int = int(floatValue)

	fmt.Printf("Type conversion - int: %d -> float: %f -> int: %d\n",
		intValue, floatValue, backToInt)
}

// GetTypeInfo returns a string describing the type of common variables
// This is a helper function to understand types
func GetTypeInfo() string {
	// Using fmt.Sprintf to create formatted strings
	var i int = 42
	var f float64 = 3.14
	var s string = "hello"
	var b bool = true

	// %T is the format verb for printing the type
	return fmt.Sprintf("Types - int: %T, float64: %T, string: %T, bool: %T",
		i, f, s, b)
} 