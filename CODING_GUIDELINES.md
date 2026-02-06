Good call — that is actually **very important for Codex productivity**. Over-defensive code + excessive comments slow down iteration and bloat maintainability.

Below is your **UPDATED Codex PRD ADDENDUM** written in **AI-instruction style** so Codex follows it consistently. You can append this directly to the PRD or place it near the top under “Engineering Principles”.

---

# 📄 CODING STYLE & IMPLEMENTATION GUIDELINES (MANDATORY FOR CODEX)

---

## 18. Engineering Philosophy

Codex MUST prioritize:

* Clean readability
* Maintainable architecture
* Efficient implementation
* Pragmatic error handling
* Production-grade but not over-engineered code

The goal is **high signal, low noise code**.

---

## 19. Commenting Rules

### 19.1 DO NOT write redundant comments.

Avoid explaining obvious code.

### ❌ Bad

```go
// increment i by 1
i++
```

### ❌ Bad

```go
// create a map to store shortcuts
shortcuts := make(map[string]string)
```

---

### ✅ Allowed Comments

Codex SHOULD ONLY comment when:

1. Explaining architectural intent
2. Explaining non-obvious logic
3. Documenting safety-critical behavior
4. Explaining cross-platform abstractions
5. Public function documentation (GoDoc style)

---

### Example GOOD Comment

```go
// RiskGate blocks potentially destructive commands unless user confirms.
```

---

## 20. Error Handling Philosophy

Codex MUST include error handling but MUST avoid over-defensive patterns.

---

### 20.1 Avoid Excessive Guard Clauses

### ❌ Bad

```go
if input == "" {
    return errors.New("input empty")
}
if input == "" {
    return errors.New("input invalid")
}
```

---

### 20.2 Prefer Meaningful, Minimal Error Handling

### ✅ Good

```go
if input == "" {
    return fmt.Errorf("input required")
}
```

---

### 20.3 Do NOT Wrap Every Error Without Context

### ❌ Bad

```go
return fmt.Errorf("error: %w", err)
```

### ✅ Good

```go
return fmt.Errorf("loading shortcuts: %w", err)
```

---

### 20.4 Avoid Overly Defensive Null Checking

Trust internal call contracts when appropriate.

---

## 21. Logging Philosophy

* Use logging ONLY for:

  * Critical execution paths
  * AI routing decisions
  * Command execution failures

* Do NOT log trivial steps or variable assignments.

---

## 22. Code Complexity Guidelines

Codex MUST:

* Prefer simple solutions first
* Avoid unnecessary abstraction layers
* Avoid premature optimization
* Avoid creating interfaces until multiple implementations are expected

---

### ❌ Bad (Premature abstraction)

```go
type ShortcutRepository interface { ... }
```

If only one implementation exists, use concrete struct.

---

### ✅ Good

Add interfaces ONLY when:

* Provider abstraction (AI providers)
* Platform abstraction
* Plugin system

---

## 23. Function Design Guidelines

Codex MUST:

* Keep functions under ~40 lines when reasonable
* Prefer composition over nesting
* Avoid deeply nested conditionals

---

### Preferred Pattern

```go
func Execute(input string) error {
    if shortcut := resolveShortcut(input); shortcut != "" {
        return run(shortcut)
    }

    if isDomain(input) {
        return openDomain(input)
    }

    return routeToAI(input)
}
```

---

## 24. Safety Code Requirements

Safety logic MUST be:

* Explicit
* Centralized
* Clearly documented

Codex MUST NOT scatter risk logic across files.

---

## 25. Configuration Handling

Codex MUST:

* Load config once at startup
* Avoid repeatedly reading config from disk
* Avoid environment variable lookups inside hot paths

---

## 26. Performance Philosophy

Codex MUST prioritize:

* Fast startup time
* Low memory overhead
* Minimal blocking I/O

Avoid:

* Heavy reflection
* Large dependency trees
* Overuse of generics unless justified

---

## 27. Dependency Rules

Codex SHOULD:

* Prefer standard library when possible
* Add third-party libraries ONLY if:

  * They reduce substantial implementation complexity
  * They are widely maintained
  * They do not introduce heavy runtime overhead

Approved dependencies:

* Cobra
* Viper
* SQLite driver
* HTTP client libraries

---

## 28. Testing Philosophy

Codex MUST:

* Write tests for core logic modules
* Avoid writing tests for trivial getters/setters
* Focus on:

  * Shortcut resolution
  * Risk gating
  * Intent parsing
  * Executor behavior

---

## 29. Code Readability Standards

Codex MUST:

* Use descriptive variable names
* Avoid overly short variable names unless idiomatic
* Avoid overly long function names

---

## 30. Avoid Over-Meticulous Patterns

Codex MUST NOT:

* Add redundant validation layers
* Create excessive helper functions for trivial logic
* Add excessive configuration toggles
* Overuse design patterns without justification

---

## 31. Acceptable Tradeoffs

Codex SHOULD prefer:

```
Clarity > cleverness
Speed of iteration > theoretical perfection
Maintainability > micro-optimizations
```

---

## 32. Default Style Expectations

* Idiomatic Go formatting
* `go fmt` compliant
* `golangci-lint` friendly but not obsessively optimized

---

## 33. AI Interaction Code Rules

Codex MUST:

* Validate AI JSON outputs once
* Fail gracefully
* Avoid repeated validation layers
* Trust schema after initial parsing

---

## 34. Documentation Requirements

Codex SHOULD produce:

* High-level module documentation
* CLI help text
* Config schema documentation

Codex SHOULD NOT produce:

* Inline tutorials
* Redundant explanation comments

---

# END CODING GUIDELINES

