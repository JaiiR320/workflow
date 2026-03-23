---
name: execute
description: Execute the next pending slice from a plan. Use when user says "execute the next slice" or wants to work on the next slice.
---

# Execute

Execute the next pending slice from a plan.

## Flow

1. **Discover context.**
   - Find plan file in `.plans/*/plan.md`
   - Read plan to understand the feature
   - Read all completed slices (if any)
   - Find next pending slice (lowest numbered file)

2. **Execute slice.**
   - Load context: plan + completed slices + current slice
   - Work through acceptance criteria
   - Implement end-to-end behavior
   - Run QA steps to verify

3. **Complete slice.**
   - Append `## Completion` section to slice file with:
     - What was built
     - Key decisions made
     - Deviations from plan and why
     - Files created or modified
     - Context for next slice
   - Move slice file from `pending/` to `completed/`
4. **Commit changes**
   - Use the commit skill to commit your changes

## Completion Section Format

```markdown
## Completion

**Built:** Brief description of what was implemented

**Decisions:** Key architectural or technical decisions made

**Deviations:** Any changes from the slice plan and why

**Files:** List of files created or modified

**Notes for next slice:** Important context for the next executor
```
