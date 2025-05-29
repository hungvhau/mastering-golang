# Quick Prompt Guide for Mastering Golang

## ⚠️ IMPORTANT: Ensuring AI Reads Context

**The AI won't automatically know about .ai/ files!** Here's how to ensure it does:

### Option 1: Quick Context (Add to any prompt)
```
[In an educational Go project with .ai/ docs. Check .ai/context.json for structure.]
Your actual request here...
```

### Option 2: Explicit Mention
```
This is the mastering-golang educational project. Read .ai/context.json first.
Add [TOPIC] to [MODULE] after [PREREQUISITES]...
```

### Option 3: Full Context (For new conversations)
```
I'm working on mastering-golang, an educational Go project organized as:
- basics/ (variables, operators, conditionals)
- functions/ (all function concepts)  
- loops/ (iteration patterns)
- .ai/ (AI documentation)
[Your request]
```

---

## ⚡ Lightning-Fast Prompt Formula

```
Add [WHAT] to [WHERE] after [PREREQUISITES].
Show [EXAMPLE 1], [EXAMPLE 2], and [PITFALL].
Focus on [KEY LEARNING POINT].
```

## 🎯 Real Examples You Can Copy

### Adding to Existing Module:
```
Add type conversion examples to basics/operators.go after variables section.
Show int to float64, string to int parsing, and common conversion errors.
Focus on why Go requires explicit conversion.
```

### Creating New Concept:
```
Add pointers to basics module after learners understand functions.
Show pointer declaration, dereferencing, and passing pointers to functions.
Include the classic "swap values" example and explain memory efficiency.
Focus on when to use pointers vs values.
```

### Full Module Creation:
```
Create a "structs" module after functions module completion.
Cover struct basics, embedding, methods with receivers, and constructor patterns.
Use a Student/Course system for examples showing progression from simple to nested structs.
Focus on how structs enable object-oriented patterns in Go.
```

## 🚀 The 5-Second Checklist

Before sending your prompt, ensure it has:
- [ ] **WHAT**: Specific concept (not vague)
- [ ] **WHERE**: Module/file location
- [ ] **WHEN**: Prerequisites clear
- [ ] **HOW**: 2-3 concrete examples
- [ ] **WHY**: Learning objective stated

## 💡 Power Words That Help AI

Use these to get better results:
- "Show progression from simple to complex"
- "Include common beginner mistakes"
- "Use [domain] examples throughout" (student/grade, shop/product, etc.)
- "Demonstrate why [concept] matters with practical example"
- "Build on existing [module] knowledge"

## 🎨 Example Domains That Work Well

For consistency across examples:
- **Academic**: Student, Course, Grade, Assignment
- **E-commerce**: Product, Cart, Order, Customer  
- **Games**: Player, Score, Level, Achievement
- **Real-world**: Person, Address, BankAccount, Transaction

## ⚠️ Instant Red Flags to Avoid

- ❌ "Add some examples about..." (too vague)
- ❌ "Implement advanced..." (breaks progression)
- ❌ Multiple topics in one prompt
- ❌ No prerequisite context
- ❌ No concrete examples specified

## 📋 Copy-Paste Templates

### Minimal (1 line):
```
Add [maps/dictionaries] to new module after loops, showing [student grades], [word counting], and [nil map pitfalls].
```

### Standard (3 lines):
```
Add goroutines introduction to new "concurrency" module.
Prerequisites: All basic modules, especially functions.
Examples: Simple goroutine, wait groups, basic channel communication.
```

### Detailed (5 lines):
```
Create "errors" module extending the functions module.
Prerequisites: Functions with multiple returns.
Objectives: Custom errors, wrapping, sentinel errors, error types.
Examples: File operations, API calls, validation errors.
Special focus: When to panic vs return error.
```

## 🏃 Speed Tips

1. **Reference existing code**: "Like the pattern in functions.go"
2. **Specify test needs**: "Include table-driven tests"
3. **Mention demo runner**: "Update cmd/XX/main.go"
4. **Use module numbers**: "After module 03-loops"
5. **State complexity**: "Basic examples only" or "Include advanced"

## 📝 The Universal Template

```
Add [TOPIC] to mastering-golang.
Location: [MODULE/FILE]
Prerequisites: [WHAT THEY KNOW]
Show: [EXAMPLE 1], [EXAMPLE 2], [EDGE CASE]
Goal: Learner understands [CORE CONCEPT]
```

Just fill in the brackets and go! 🚀 