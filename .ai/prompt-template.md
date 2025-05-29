# Prompt Template for Adding Educational Content

## 🎯 Effective Prompt Structure

A good prompt for expanding this educational repository should include:
1. **Concept to teach** - Clear topic definition
2. **Prerequisites** - What learners should know first
3. **Learning objectives** - What they'll understand after
4. **Placement** - Where it fits in the progression
5. **Examples needed** - Specific scenarios to demonstrate

## 📝 Prompt Template

```
Add [CONCEPT NAME] to the mastering-golang project.

Prerequisites: [What modules/concepts should learners complete first]

Learning Objectives:
- [Objective 1]
- [Objective 2]
- [Objective 3]

Placement: [Which module - basics/functions/loops OR new module]

Examples to Include:
1. [Basic example scenario]
2. [Intermediate example]
3. [Common pitfall/error case]
4. [Real-world application]

Additional Requirements:
- [Any specific patterns to demonstrate]
- [Edge cases to cover]
- [Performance considerations if relevant]
```

## 🌟 Example Prompt: Adding Arrays

```
Add Arrays and Slices to the mastering-golang project.

Prerequisites: Variables, basic functions, and loops modules should be completed

Learning Objectives:
- Understand the difference between arrays and slices
- Learn how to declare, initialize, and manipulate both
- Understand slice capacity vs length
- Master common slice operations (append, copy, sub-slicing)

Placement: New module called "collections" (as arrays/slices are foundational for maps and structs)

Examples to Include:
1. Basic array declaration with fixed size
2. Slice creation from arrays and make()
3. Dynamic slice growth with append
4. Common pitfall: slice capacity issues
5. Real-world: Processing a list of student grades

Additional Requirements:
- Show memory efficiency comparison between arrays and slices
- Demonstrate the backing array concept
- Include range loops over slices
- Show multi-dimensional slices
```

## 🚀 Quick Prompt Examples

### Minimal Effective Prompt:
```
Add error handling patterns to the functions module. Show custom error types, 
error wrapping with fmt.Errorf, and the errors.Is/As patterns. Include examples 
of when to use each approach.
```

### Detailed Module Prompt:
```
Create a new "structs" module after learners complete functions. Cover:
- Basic struct declaration
- Struct embedding
- Methods on structs
- Pointer vs value receivers
- Constructor patterns
Include a "Student" type example building from simple to complex.
```

## 💡 What Makes a Prompt Effective

### ✅ DO Include:
- **Specific concept boundaries** - "arrays and slices" not just "collections"
- **Prerequisites** - What knowledge is assumed
- **Concrete examples** - "Student grades" not just "data"
- **Progression hints** - Where it fits in learning path

### ❌ DON'T Include:
- Vague requests - "add some data structure examples"
- Advanced patterns before basics - "add generics" before interfaces
- Multiple unrelated concepts - Keep focused on one topic
- Performance optimization as primary goal

## 📊 Prompt Complexity Levels

### Level 1: Add to Existing Module
```
Add string manipulation examples to the basics module. Include:
- String concatenation
- String formatting with fmt.Sprintf
- Basic string functions from strings package
```

### Level 2: Extend with New File
```
Add a "types.go" file to basics module covering:
- Type aliases
- Custom type definitions  
- Type methods
This builds on variables.go knowledge.
```

### Level 3: Create New Module
```
Create an "interfaces" module (after structs). Include:
- Interface declaration
- Implicit satisfaction
- Empty interface
- Type assertions
- Common interfaces (Stringer, error)
Use shapes (Circle, Rectangle) for examples.
```

## 🎯 AI Understanding Accelerators

To help AI assistants work faster, always specify:

1. **Target audience level** - "learner just finished functions module"
2. **Example domain** - "use student/classroom examples" 
3. **File structure** - "add to existing file" vs "create new file"
4. **Test requirements** - "include table-driven tests"
5. **Demo runner needs** - "update cmd/02-functions/main.go"

## 📝 Copy-Paste Template

```
Add [TOPIC] to mastering-golang.

After completing: [PREREQUISITES]
Learners will understand: [KEY CONCEPTS]
Place in: [MODULE/FILE]
Examples: [2-3 SPECIFIC SCENARIOS]
Special focus: [ANY SPECIFIC EMPHASIS]
``` 