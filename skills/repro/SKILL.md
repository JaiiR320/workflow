---
name: repro
description: Reproduce a bug by understanding it thoroughly then writing a failing test. Use when user reports a bug, says "repro this", or wants to write a test that captures broken behavior.
---

# Repro

Write a failing test that reproduces a bug.

## Gate Phrases

Do NOT write the test until user says:
- "write the repro test"
- "create the test"
- "write the test now"

Stay in discussion mode otherwise.

## Rules

1. **Understand the bug first.** Ask clarifying questions one at a time until reproduction steps are clear.

2. **FORBIDDEN: Reading source implementation files.** Do not open, read, grep, or explore ANY source files. Do not read git logs or diffs. The ONLY files you may read are existing test files to understand testing patterns. No exceptions.

3. **Bug description and questions are your only input.** Write the test based solely on what the user tells you and what you learn from existing test files.

4. **One bug, one test.** Capture the exact broken behavior.

5. **Test behavior, not implementation.** Assert on observable outcomes — API responses, UI state, return values.

## Flow

### Phase 1: Interview

1. Ask clarifying questions one at a time: what input triggers it, what happens, what should happen.
2. Stop when reproduction steps are clear.

### Phase 2: Prepare

3. Read existing test files ONLY to understand testing patterns.
4. Present test plan in chat.
5. Say: "Reproduction steps clear. Say 'write the repro test' to create the failing test."
6. Wait for gate phrase.

### Phase 3: Write

7. Write failing test.
8. Run it to confirm it fails.

## Next Step

Failing test created. Next: fix the bug. Say "squash this" when ready.