# 🚀 Bootstrap Prompt - Start Here!

## Copy-Paste This First (For New AI Sessions)

```
I'm working on the mastering-golang educational project. 
Please read .ai/context.json for project structure and .ai-prompt.md for guidelines.
This is a learning-focused Go project with modules in basics/, functions/, and loops/.
```

## Why This Works

This minimal prompt:
1. Identifies the project type (educational)
2. Points to key context files
3. Mentions the main directories
4. Sets the "learning-focused" tone

## Platform-Specific Instructions

### Cursor
- May auto-detect `.ai-prompt.md` and `.cursorrules`
- Still good to mention the project type

### ChatGPT/Claude (Web)
- Always paste the bootstrap prompt first
- Can follow with: "Now read .ai/context.json"

### GitHub Copilot Chat
- Include in first message: "Check .ai/ folder"

### VS Code + AI Extensions
- Mention: "Project has .ai/ documentation"

## Even Shorter Version

If you're in a hurry, just prefix your prompt with:
```
[mastering-golang edu project, see .ai/ docs] 
```

## Full Context Loading Sequence

For complex tasks, load context in this order:
1. This bootstrap message
2. `.ai/context.json` (structure)
3. `.ai-prompt.md` (principles)
4. `.ai/codemap.md` (navigation)
5. Your specific request

## Example Full Prompt

```
I'm working on mastering-golang (educational Go project).
Read: .ai/context.json, .ai-prompt.md, .ai/learning-roadmap.md

Add pointers to basics module after functions prerequisites.
Show pointer declaration, dereferencing, and pointer parameters.
Include swap example and memory efficiency benefits.
``` 