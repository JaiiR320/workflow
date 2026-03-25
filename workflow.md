# Workflow

Go from idea to implementation end-to-end.

## Overview

This workflow uses skills that interview first, then write files when explicitly approved. Each skill produces files that serve as handoffs to the next skill.

## Main Flow

1. **Refine** - Say "refine this" and describe your idea. Interview clarifies the feature. Say "looks good, write it" to produce `.plans/<feature>/plan.md`.

2. **Slice** - Say "slice this" to break the plan into vertical, independently testable slices. Discuss the breakdown, then say "slice it up" to write files to `.plans/<feature>/pending/slice-1.md`, etc.

3. **Execute Loop** - For each pending slice:
   - Say "execute the next slice for plan <feature>". Skill loads only that plan's context (`plan.md` + completed slices + current pending slice) and then implements end-to-end.
   - Say "commit" to commit the completed slice.
   - Repeat until all slices are in `completed/`.
   - Can be used in a "ralph" style loop, or manually in a new session.

## Bug Fixing

Use these when bugs are discovered after slices are completed, and you are QAing:

- **Repro** - Say "repro this" and describe the bug. Interview clarifies reproduction steps. Say "write the repro test" to create a failing test. **Forbidden from reading source code** - only test files.

- **Squash** - Say "squash this" to fix the failing test with minimal changes.

## Key Concepts

- **Gate phrases:** Interview skills (refine, slice, repro) never write files until you say specific phrases.
- **Vertical slices:** Each slice cuts through all layers (schema, API, UI, tests) and is independently verifiable.
- **Execution loop:** Execute slice for the target plan → Commit → Next slice. Fresh sessions recommended per slice to avoid context bloat.

## Directory Structure

```
.plans/
  <feature>/
    plan.md
    pending/
      slice-1.md
      slice-2.md
    completed/
      slice-1.md
```

## Example Usage

### Refine and Slice

https://opncd.ai/share/e0rSA4xg

### Execution and Committing

https://opncd.ai/share/Nm2pelmB
