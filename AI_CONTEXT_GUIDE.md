# 🤖 How to Ensure AI Uses the Context Files

## The Problem
AI assistants don't automatically know about `.ai/` folders or special documentation files unless you tell them!

## The Solution

### 🚀 Method 1: Bootstrap Prompt (Recommended)
Start every new AI conversation with:
```
I'm working on the mastering-golang educational project. 
Please read .ai/context.json for project structure and .ai-prompt.md for guidelines.
```

### ⚡ Method 2: Quick Prefix
Add this to any prompt:
```
[mastering-golang edu project, see .ai/ docs] Your request here...
```

### 📋 Method 3: Copy Context
For web-based AI (ChatGPT, Claude), you can:
1. Copy the contents of `.ai-prompt.md`
2. Paste it at the start of your conversation
3. Then make your request

## Platform-Specific Behavior

| Platform | Auto-detects? | What to do |
|----------|--------------|------------|
| **Cursor** | Sometimes (`.cursorrules`, `.ai-prompt.md`) | Still mention project type |
| **GitHub Copilot** | No | Always include context |
| **ChatGPT/Claude** | No | Paste bootstrap or mention files |
| **VS Code AI** | Varies | Include context to be safe |

## Examples

### ❌ Without Context (AI doesn't understand project):
```
Add error handling to the project
```

### ✅ With Context (AI knows exactly what to do):
```
[mastering-golang edu project, see .ai/ docs]
Add error handling examples to functions module after the multiple returns section.
Show custom errors, wrapping, and when to use errors vs panic.
```

## Why This Matters

With context, AI will:
- ✅ Maintain educational tone
- ✅ Follow your module structure  
- ✅ Create appropriate tests
- ✅ Update demo runners
- ✅ Use consistent examples

Without context, AI might:
- ❌ Create production-style code
- ❌ Miss the educational focus
- ❌ Put files in wrong locations
- ❌ Skip important comments
- ❌ Break the learning progression

## Quick Test

Try these two prompts in a new AI session and see the difference:

**Without context:**
```
Add interfaces to the Go project
```

**With context:**
```
[mastering-golang edu project, see .ai/ docs]
Create interfaces module after structs (which comes after functions).
Include interface declaration, satisfaction, and common patterns.
Use Shape/Animal examples with clear educational comments.
```

The second prompt will generate much better educational content!

## Pro Tip

Create a text expander or snippet for:
```
[mastering-golang edu project, see .ai/ docs] 
```
Then you can quickly prefix any prompt! 