---
name: refine
description: Refine a feature idea through focused questions, then produce a plan document. Use when user wants to think through a feature, work through a design, says "refine this", or wants to go from a rough idea to a concrete plan.
---

# Refine

Refine a feature idea into a plan document.

## Gate Phrases

Do NOT write files until user says:
- "looks good, write it"
- "approved, save the plan"
- "write the plan"
- "save to file"

Stay in discussion mode otherwise.

## Rules

1. **One question per response.** Never present multiple questions, never use bullet lists of questions, never ask "should it be X or Y or Z?" Ask exactly ONE question, wait for the answer, then ask the next.

2. **Stay at the right altitude.** Architecture, interfaces, data flow, state management, expected behavior. Not edge cases or implementation minutiae.

3. **Let context guide you.** Don't ask things already resolved or inferable from previous answers.

4. **Know when to stop.** When the design space is covered, present the summary. Don't drill for the sake of it.

5. **Explore the codebase when it helps.** If a question can be answered by looking at existing code, look instead of asking.

## Flow

### Phase 1: Interview

1. Ask questions one at a time, working through the design space.
2. When done, present summary of decisions in chat.
3. Say: "Design space covered. Say 'looks good, write it' to create the plan."
4. Wait for gate phrase.

### Phase 2: Write

5. Read template at `template.md`.
6. Write plan following template structure. Present in chat first.
7. Once approved, write to `.plans/<feature-name>/plan.md`. Ask for feature name if not obvious.

## Next Step

Plan written. Next: slice this plan. Say "slice this" when ready.