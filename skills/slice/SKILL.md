---
name: slice
description: Break a plan into vertical slices that can be executed independently. Use when user has a plan and wants to break it into small, independently verifiable units of work, says "slice this", or wants to prepare work for execution.
---

# Slice

Break a plan into vertical slices.

## Gate Phrases

Do NOT write files until user says:
- "slice it up"
- "approved, write the slices"
- "save the slices"
- "write them to files"

Stay in discussion mode otherwise.

## Rules

1. **Vertical, not horizontal.** Each slice goes through every layer — schema, API, UI, tests. Never a slice that's just "set up database" or "build frontend."

2. **Independently verifiable.** Every slice must have a meaningful QA step. If you can't describe how to verify it, it's too small. If QA covers more than one distinct user flow, it's too big.

3. **Thin but not trivial.** Renaming variables or adding a single if statement is not a slice. Deliver a narrow but complete behavior.

4. **First slice is the simplest end-to-end path.** The "hello world" tracer bullet. Later slices add breadth — more user stories, edge cases, polish.

5. **Keep the count reasonable.** More than 10 slices means the feature is too big. Split into multiple plans. One or two slices is fine for straightforward features.

## Flow

### Phase 1: Draft

1. Use conversation context as primary source of truth. Reference `.plans/<feature-name>/plan.md` only if needed to fill gaps.
2. Explore codebase to understand existing patterns, integration layers, and natural seams.
3. Draft slices. For each slice, show:
   - Title
   - What it does end-to-end
   - Dependencies (if any)
   - Expected behaviors addressed
4. Present full breakdown in chat.
5. Say: "Draft complete. Say 'slice it up' to save slices to files."
6. Wait for gate phrase.

### Phase 2: Write

7. Read template at `template.md`.
8. Write each slice as numbered file in `.plans/<feature-name>/pending/`.

## Next Step

Slices written. Next: execute slices. Say "execute the next slice" when ready.

## Completion

When a slice finishes during execution:

1. Append `## Completion` section describing what was built, decisions made, deviations, and context for next slice.
2. Move file from `pending/` to `completed/`.