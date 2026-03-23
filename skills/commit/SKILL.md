---
name: commit
description: Commit the changes made in the current session with a message that matches the project's commit style. Use only when explicitly asked to commit.
---

# Commit

Commit changes from this session.

## Rules

1. **Only commit changes from this session.** Do not stage unrelated files. Use `git diff` and `git status` to identify what changed.

2. **Match the project's commit style.** Read the last 10 commit messages with `git log --oneline -10` and match the tone, format, and conventions.

3. **One commit.** Stage all session changes and commit them together.

## Flow

1. Run `git status` and `git diff` to identify session changes.
2. Run `git log --oneline -10` to see the project's commit style.
3. Stage the changed files.
4. Commit with a message that matches the project's conventions.