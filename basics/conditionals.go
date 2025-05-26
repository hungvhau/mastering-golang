// Package basics - Conditional statements demonstration
package basics

// GetAgeCategory demonstrates basic if-else statements
// This function categorizes age into different groups
// Parameter: age - the age to categorize
// Returns: a string describing the age category
func GetAgeCategory(age int) string {
	// Basic if statement
	// The condition must be a boolean expression (no parentheses needed)
	if age < 0 {
		// This block executes if the condition is true
		return "Invalid age"
	}

	// if-else statement
	// else must be on the same line as the closing brace of if
	if age < 13 {
		return "Child"
	} else if age < 20 {
		// else if allows multiple conditions to be checked in sequence
		return "Teenager"
	} else if age < 60 {
		return "Adult"
	} else {
		// else block executes when all previous conditions are false
		return "Senior"
	}
}

// CheckNumber demonstrates if statements with initialization
// This shows Go's unique feature of initializing variables in if statements
func CheckNumber(num int) string {
	// If statement with initialization
	// The variable 'doubled' is only available within this if-else block
	if doubled := num * 2; doubled > 100 {
		// 'doubled' is accessible here
		return "Double of number is greater than 100"
	} else if doubled == 100 {
		// 'doubled' is still accessible in else if
		return "Double of number equals 100"
	} else {
		// 'doubled' is accessible in else too
		return "Double of number is less than 100"
	}
	// 'doubled' is NOT accessible here - it's out of scope
}

// EvaluateGrade demonstrates switch statements
// Switch is cleaner than multiple if-else statements for multiple conditions
func EvaluateGrade(score int) string {
	// Basic switch statement
	// Go's switch doesn't need break statements (implicit break)
	switch {
	// Switch without expression evaluates boolean conditions
	case score >= 90:
		return "A - Excellent!"
	case score >= 80:
		return "B - Good job!"
	case score >= 70:
		return "C - Satisfactory"
	case score >= 60:
		return "D - Needs improvement"
	default:
		// default case executes when no other case matches
		return "F - Failed"
	}
}

// GetDayType demonstrates switch with expression
// This shows a more traditional switch usage
func GetDayType(dayNumber int) string {
	// Switch with expression
	// Each case is compared with the expression value
	switch dayNumber {
	case 1, 2, 3, 4, 5:
		// Multiple values can be in one case
		return "Weekday"
	case 6, 7:
		return "Weekend"
	default:
		return "Invalid day number"
	}
}

// GetMonthQuarter demonstrates switch with fallthrough
// fallthrough is rarely used but important to understand
func GetMonthQuarter(month int) string {
	var quarter string
	
	// Switch to determine quarter
	switch month {
	case 1, 2, 3:
		quarter = "Q1"
	case 4, 5, 6:
		quarter = "Q2"
	case 7, 8, 9:
		quarter = "Q3"
	case 10, 11, 12:
		quarter = "Q4"
	default:
		return "Invalid month"
	}
	
	return quarter
}

// AnalyzeCharacter demonstrates type switch
// Type switches are useful when working with interfaces
func AnalyzeCharacter(char rune) string {
	// Switch on multiple conditions
	switch {
	case char >= 'a' && char <= 'z':
		return "Lowercase letter"
	case char >= 'A' && char <= 'Z':
		return "Uppercase letter"
	case char >= '0' && char <= '9':
		return "Digit"
	case char == ' ':
		return "Space"
	case char == '\n':
		return "Newline"
	case char == '\t':
		return "Tab"
	default:
		return "Special character"
	}
}

// ProcessValue demonstrates switch with initialization
// Similar to if statements, switches can have initialization
func ProcessValue(value int) string {
	// Switch with initialization statement
	// remainder is only available within the switch block
	switch remainder := value % 3; remainder {
	case 0:
		return "Divisible by 3"
	case 1:
		return "Remainder 1 when divided by 3"
	case 2:
		return "Remainder 2 when divided by 3"
	default:
		// This default case will never execute for modulo 3
		// But it's good practice to include it
		return "Unexpected value"
	}
}

// CheckConditions demonstrates logical operators in conditions
// Shows how to combine multiple conditions
func CheckConditions(x, y int) string {
	// Logical AND (&&)
	// Both conditions must be true
	if x > 0 && y > 0 {
		return "Both numbers are positive"
	}
	
	// Logical OR (||)
	// At least one condition must be true
	if x < 0 || y < 0 {
		return "At least one number is negative"
	}
	
	// Logical NOT (!)
	// Inverts the boolean value
	if !(x == 0 && y == 0) {
		return "Not both zeros"
	}
	
	return "Both are zeros"
}

// ValidateInput demonstrates nested conditions
// Shows how conditions can be nested for complex logic
func ValidateInput(username string, age int) string {
	// Check username length first
	if len(username) < 3 {
		return "Username too short"
	} else if len(username) > 20 {
		return "Username too long"
	} else {
		// Username length is valid, now check age
		if age < 13 {
			return "Too young to register"
		} else if age > 120 {
			return "Invalid age"
		} else {
			// Both username and age are valid
			if age >= 18 {
				return "Registration complete - Full access granted"
			} else {
				return "Registration complete - Limited access (under 18)"
			}
		}
	}
}

// GetSeason demonstrates switch without fallthrough (Go default)
// Each case automatically breaks unless fallthrough is used
func GetSeason(month string) string {
	// Switch on string values
	switch month {
	case "December", "January", "February":
		return "Winter"
	case "March", "April", "May":
		return "Spring"
	case "June", "July", "August":
		return "Summer"
	case "September", "October", "November":
		return "Fall"
	default:
		return "Invalid month"
	}
	// No break needed - Go automatically breaks after each case
} 