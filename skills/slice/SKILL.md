---
name: slice
description: Break a plan into vertical slices that can be executed independently. Use when user has a plan and wants to break it into small, independently verifiable units of work, says "slice this", or wants to prepare work for execution.
---

# Slice

Break a plan into thin vertical slices. Each slice cuts through all layers of the system end-to-end and is independently verifiable.

## Rules

1. **Vertical, not horizontal.** Each slice goes through every layer — schema, API, UI, tests. Never a slice that's just "set up the database" or "build the frontend."

2. **Independently verifiable.** Every slice must have a meaningful QA step. If you can't describe how to verify it, it's too small. If the QA step covers more than one distinct user flow, it's too big.

3. **Thin but not trivial.** Renaming variables or adding a single if statement is not a slice. A slice delivers a narrow but complete behavior.

4. **First slice is the simplest end-to-end path.** The "hello world" tracer bullet. Later slices add breadth — more user stories, edge cases, polish.

5. **Keep the count reasonable.** If you're looking at more than 10 slices, the feature is probably too big and should be split into multiple plans. One or two slices is fine for a straightforward feature.

## Flow

1. Use the conversation context as the primary source of truth. Reference `.plans/<feature-name>/plan.md` only if needed to fill gaps.
2. Explore the codebase to understand existing patterns, integration layers, and natural seams.
3. Draft the slices. For each slice, show:
   - Title
   - What it does end-to-end
   - Which slices it depends on (if any)
   - Which expected behaviors from the plan it addresses
4. Present the full breakdown in chat. Wait for the user to approve, adjust, merge, split, or reorder.
5. Once approved, read the template at `template.md` in this skill's directory.
6. Write each slice as a numbered file in `.plans/<feature-name>/pending/`, creating the directory if needed.

## Completion

When a slice is finished during execution (by the user or an agent), the executor should:

1. Append a `## Completion` section to the slice file describing what was built, decisions made, deviations from the plan, and anything the next slice should know.
2. Move the file from `pending/` to `completed/`.
