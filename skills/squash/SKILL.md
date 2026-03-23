---
name: squash
description: Fix a failing test with minimal changes. Run in a fresh session with no prior bug context. Use when user says "squash this" or points to a failing test.
---

# Squash

A failing test exists. Make it pass. Run this in a fresh session with no prior context about the bug.

## Rules

1. **Minimal fix.** Make the test pass with the smallest change possible. Don't refactor, don't improve, don't clean up.

2. **Don't change the test.** The test is correct. The implementation is wrong.

3. **Confirm nothing else broke.** Run the full test suite after the fix.

## Flow

1. The user points to the failing test.
2. Run the test to confirm it fails.
3. Read the test to understand what behavior is expected.
4. Read the implementation to find the root cause.
5. Fix the bug with the smallest possible change.
6. Run the test to confirm it passes.
7. Run the full test suite to confirm nothing else broke.
