Perfect — below is a **FULL END-TO-END MASTER PRD** that includes:

• Product vision
• Engineering specs
• Architecture
• Coding standards
• Build + packaging
• Distribution
• Launch checklist
• Shippable definition
• Post-launch roadmap

This is written so **Codex (or any AI dev agent)** can start from **literally nothing → shippable macOS tool**.

I am going to assume final product name:

# ⭐ Orion (CLI alias: `o`)

If you change the name later, this PRD still holds.

---

# 📄 MASTER PRODUCT REQUIREMENTS DOCUMENT

# Project: ORION

# Version: 1.0 (Codex Implementation Spec)

---

# 1. PRODUCT OVERVIEW

## 1.1 Vision

Orion is a natural-language terminal assistant that converts conversational user commands into safe, structured system automation.

The tool allows users to:

* Open applications
* Navigate workflows
* Execute shell tasks
* Automate repetitive commands
* Perform system operations safely
* Use optional AI-based command interpretation

---

## 1.2 Product Mission

Reduce friction between user intent and system execution while maintaining strong safety guarantees.

---

## 1.3 Example Usage

```
o msu d2l
o open spotify
o search discrete math cheat sheet
o ai clean node_modules folders
```

---

# 2. TARGET PLATFORMS

## Phase 1

macOS only

## Phase 2

Linux + Windows

---

# 3. CORE VALUE PROPOSITION

Orion eliminates repetitive system navigation by acting as:

* Shortcut launcher
* Intelligent command router
* Safe automation layer
* AI-enhanced terminal interface

---

# 4. SUCCESS CRITERIA

Orion is considered shippable when:

✔ Installable via download or Homebrew
✔ Executes shortcuts reliably
✔ AI commands route safely
✔ Risk confirmation blocks dangerous commands
✔ Config system stable
✔ `o doctor` validates environment
✔ Crash-free execution in common workflows

---

# 5. TECH STACK REQUIREMENTS

## Core Language

Go (Golang)

## CLI Framework

Cobra

## Config Management

Viper

## Storage

SQLite

## AI Providers

Ollama (default offline)
Gemini (cloud fallback)

## macOS Integration

* `open`
* AppleScript when necessary

---

# 6. SYSTEM ARCHITECTURE

---

## 6.1 High-Level Flow

```
User Input
   ↓
Shortcut Resolver
   ↓
Domain Detector
   ↓
Search Fallback
   ↓
AI Router (if needed)
   ↓
Risk Gate
   ↓
Executor
```

---

## 6.2 Directory Structure

```
orion/
 ├ cmd/
 │   ├ root.go
 │   ├ ai.go
 │   ├ add.go
 │   ├ list.go
 │   ├ doctor.go
 │   ├ update.go
 │
 ├ internal/
 │   ├ shortcuts/
 │   ├ router/
 │   ├ executor/
 │   ├ safety/
 │   ├ ranking/
 │   ├ config/
 │
 ├ providers/
 │   ├ ollama.go
 │   ├ gemini.go
 │
 ├ models/
 │   ├ intent.go
 │   ├ history.go
 │
 ├ storage/
 │   ├ sqlite.go
```

---

# 7. FUNCTIONAL REQUIREMENTS

---

## 7.1 Shortcut Resolution

### Input

```
o msu d2l
```

### Behavior

Match fuzzy shortcut from:

```
~/.config/orion/shortcuts.yaml
```

### Output

Execute mapped command

---

## 7.2 Domain Detection

If input matches domain pattern:

```
*.com
*.edu
*.org
```

Execute:

```
open https://domain
```

---

## 7.3 Search Fallback

If no other match:

```
open https://google.com/search?q=query
```

---

## 7.4 Application Launcher

```
open -a "<App Name>"
```

---

## 7.5 AI Intent Router

Triggered when:

* No shortcut match
* Not domain
* Not search
* OR user explicitly calls:

```
o ai <task>
```

---

# 8. AI INTENT SPECIFICATION

---

## 8.1 LLM Output Schema

Codex MUST enforce:

```
{
  "action": string,
  "args": object,
  "risk": "low" | "medium" | "high"
}
```

---

## 8.2 Supported Actions

| Action     | Description          |
| ---------- | -------------------- |
| open_url   | Browser open         |
| open_app   | App launcher         |
| search     | Web search           |
| run_shell  | Safe shell execution |
| file_find  | File system search   |
| git_helper | Git workflow         |

---

## 8.3 Risk Gate

| Risk   | Behavior                   |
| ------ | -------------------------- |
| Low    | Execute                    |
| Medium | Confirm                    |
| High   | Block or explicit override |

---

# 9. COMMAND EXECUTOR REQUIREMENTS

Executor MUST:

* Run only validated commands
* Use platform abstraction
* Return execution result + error status

---

# 10. CONFIGURATION SYSTEM

---

## Config File

```
~/.config/orion/config.yaml
```

### Fields

```
ai_provider
search_engine
risk_threshold
features
```

---

# 11. HISTORY & RANKING

SQLite database:

```
~/.config/orion/history.db
```

Tracked data:

* command text
* success/failure
* timestamp
* usage count

Used for fuzzy ranking.

---

# 12. CLI COMMAND SET

```
o <input>
o add "<phrase>" "<command>"
o list
o edit
o ai "<task>"
o doctor
o update
```

---

# 13. DOCTOR COMMAND

Validates:

* Config integrity
* AI provider availability
* Shortcut validity
* Environment readiness

---

# 14. SAFETY POLICY

Orion MUST:

* Never execute raw LLM shell output
* Require confirmation for destructive commands
* Centralize risk logic in safety module

---

# 15. CODING STYLE RULES (MANDATORY)

---

## Comment Rules

Do NOT write redundant comments.
Only document architecture or complex logic.

---

## Error Handling

Include meaningful but not excessive validation.
Avoid redundant guard clauses.

---

## Complexity Rules

Prefer simple implementation.
Avoid premature abstraction.

---

## Function Limits

Target < 40 lines when possible.

---

## Dependency Rules

Prefer standard library.
Allow only approved dependencies.

---

# 16. PERFORMANCE REQUIREMENTS

* Startup time < 50ms
* Non-AI command < 100ms
* Minimal memory footprint

---

# 17. SECURITY REQUIREMENTS

* Structured LLM outputs only
* Centralized risk evaluation
* Confirmation UX mandatory

---

# 18. PACKAGING REQUIREMENTS

---

## Build Targets

```
darwin_arm64
darwin_amd64
```

---

## Release Artifacts

```
orion_darwin_arm64
orion_darwin_amd64
checksums.txt
```

---

## Installation Methods

### Method 1: Direct Download

```
curl → chmod → move to /usr/local/bin
```

### Method 2: Homebrew Tap

---

# 19. WEBSITE DISTRIBUTION SPEC

Website MUST include:

* Product description
* Download links
* Install instructions
* Demo
* Known limitations

---

# 20. VERSIONING STRATEGY

Semantic Versioning:

```
v0.1.0 → Public Beta
v1.0.0 → Stable Release
```

---

# 21. TESTING REQUIREMENTS

Codex MUST implement tests for:

* Shortcut matching
* Domain detection
* Risk gating
* Intent parsing
* Executor behavior

---

# 22. SHIPPABLE DEFINITION

Orion is shippable when:

✔ Fresh mac install works
✔ Core commands functional
✔ AI routing safe
✔ Packaging install works
✔ Documentation available

---

# 23. LAUNCH CHECKLIST

---

### Build

* Compile binaries
* Generate checksums

### Release

* Upload GitHub release
* Update Homebrew tap

### Website

* Upload install instructions
* Publish download links

---

# 24. FUTURE ROADMAP

* Plugin SDK
* Cross-platform support
* Background daemon
* GUI launcher
* License system
* Mobile companion

---

# 25. DEVELOPMENT MILESTONES

---

### Milestone 1

CLI + shortcuts + config

### Milestone 2

History + ranking

### Milestone 3

AI router + Ollama

### Milestone 4

Gemini support

### Milestone 5

Packaging + website

---

# 26. ACCEPTANCE CRITERIA

Codex must deliver:

* Fully compiling project
* Working CLI commands
* Packaging scripts
* Documentation

---

# END MASTER PRD

---