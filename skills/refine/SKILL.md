---
name: refine
description: Refine a feature idea through focused questions, then produce a plan document. Use when user wants to think through a feature, work through a design, says "refine this", or wants to go from a rough idea to a concrete plan.
---

# Refine

Refine the user's feature or design idea into a shared understanding, then produce a plan document.

## Rules

1. **One question at a time. Always.** Never present a list of questions.

2. **Stay at the right altitude.** Architecture, interfaces, data flow, state management, expected behavior, coding patterns and idioms. Not edge cases, defensive coding, or implementation minutiae.

3. **Let context guide you.** Don't ask things already resolved or inferable from previous answers.

4. **Know when to stop.** When the design space is covered, say so and summarize what was decided. Don't drill for the sake of it.

5. **Explore the codebase when it helps.** If a question can be answered by looking at existing code, look at the code instead of asking.

## Flow

### Phase 1: Refine

1. The user describes their feature or problem, then invokes this skill.
2. Ask questions one at a time, working through the design space.
3. When done, present a summary of decisions in chat. Wait for confirmation.
4. If the user raises new concerns or changes a decision, continue asking questions until the new branches are resolved, then re-summarize.

### Phase 2: Plan

5. Read the template at `template.md` in this skill's directory.
6. Write the plan following the template structure. Present it in chat first.
7. Wait for the user to approve or request changes.
8. Once approved, write the plan to `.plans/<feature-name>/plan.md`, creating the directory if needed. Ask the user for the feature name if it isn't obvious from context.
