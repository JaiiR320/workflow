---
name: repro
description: Reproduce a bug by understanding it thoroughly then writing a failing test. Use when user reports a bug, says "repro this", or wants to write a test that captures broken behavior.
---

# Repro

Understand a bug, then write a failing test that reproduces it. Nothing else.

## Rules

1. **Understand the bug first.** Ask clarifying questions one at a time until you know exactly how to reproduce it.

2. **Do NOT read implementation files.** Do not open, read, grep, or explore any source files. Do not read git logs or diffs. The only files you may read are existing test files to understand testing patterns and conventions.

3. **The bug description and your questions are your only input.** Write the test based solely on what the user tells you and what you learn from existing test files.

4. **One bug, one test.** The test should capture the exact broken behavior.

5. **Test behavior, not implementation.** Assert on observable outcomes — API responses, UI state, return values.

## Flow

1. The user describes the bug.
2. Ask clarifying questions one at a time: what input triggers it, what happens, what should happen. Stop when reproduction steps are clear.
3. If needed, read existing test files only to understand testing patterns and conventions.
4. Write the failing test. Run it to confirm it fails.
