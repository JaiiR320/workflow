# Slice 4: Delete Command

## Dependencies

Slice 3: Get and List Commands

## Description

Implements the `delete` command to remove key-value pairs from storage. Handles cases where the key doesn't exist. File is updated immediately after deletion.

## Expected Behaviors Addressed

- Run `store delete <key>` to remove a key-value pair

## Acceptance Criteria

- [x] `delete` command removes key from storage
- [x] File is updated after deletion
- [x] Deleting non-existent key shows "not found" message
- [x] Other keys remain intact after deletion

## QA

1. Set up data:
   ```
   ./store set name Alice
   ./store set age 25
   ./store set city NYC
   ```
2. Run `./store delete age`
3. Check `store.json` - should only have `name` and `city`
4. Run `./store list` - confirm only 2 entries
5. Run `./store delete missing` -> see "not found" message
6. Run `./store get age` -> see "not found"

## Completion

**Built:** Implemented persisted `delete` behavior in `main.go` so keys are removed from `store.json`, existing entries remain intact, and missing keys print `not found`.

**Decisions:**
- Reused `loadStore` and `saveStore` so delete follows the same storage path as `set`, `get`, and `list`
- Validated `delete` expects exactly one key argument, matching the stricter `get` behavior
- Treated missing keys as a non-error result and skipped rewriting the file when nothing changed
- Printed `deleted <key>` on success for explicit CLI feedback

**Deviations:** None - followed the slice specification exactly.

**Files:**
- `main.go` - implemented persisted delete logic and argument validation
- `.plans/store-cli/completed/4-delete-command.md` - recorded slice completion details

**Notes for next slice:**
- Core CRUD commands are now implemented in the single-file CLI
- A temp-dir QA run verified delete updates `store.json`, preserves other keys, and reports missing keys cleanly
- Next slice can focus on edge cases and automated tests without needing more command-surface changes
