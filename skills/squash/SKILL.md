---
name: squash
description: Fix a failing test with minimal changes. Run in a fresh session with no prior bug context. Use when user says "squash this" or points to a failing test.
---

# Squash

Fix a failing test with minimal changes.

## Rules

1. **Minimal fix.** Make the test pass with the smallest change possible. Don't refactor, don't improve, don't clean up.

2. **Don't change the test.** The test is correct. The implementation is wrong.

3. **Confirm nothing else broke.** Run the full test suite after the fix.

## Flow

1. Run the test to confirm it fails.
2. Read the test to understand expected behavior.
3. Read the implementation to find the root cause.
4. Fix the bug with the smallest possible change.
5. Run the test to confirm it passes.
6. Run the full test suite to confirm nothing else broke.