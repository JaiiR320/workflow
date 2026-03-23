# Slice 3: Get and List Commands

## Dependencies

Slice 2: File Storage Foundation

## Description

Implements the `get` and `list` commands for data retrieval. `get` retrieves a value by key and displays it. `list` shows all key-value pairs in a readable, formatted table. Handles "key not found" cases gracefully.

## Expected Behaviors Addressed

- Run `store get <key>` to retrieve the value for a specific key
- Run `store list` to display all stored key-value pairs

## Acceptance Criteria

- [x] `get` command reads from file and displays value
- [x] `get` with non-existent key shows "not found" message
- [x] `list` command displays all key-value pairs
- [x] Output is formatted for readability (table or aligned columns)
- [x] Empty store shows appropriate message for `list`

## QA

1. Set some data:
   ```
   ./store set name Alice
   ./store set age 25
   ./store set city "New York"
   ```
2. Run `./store get name` → see "Alice"
3. Run `./store get missing` → see "not found" message
4. Run `./store list` → see formatted table with all 3 pairs
5. Delete `store.json`, run `./store list` → see "no entries" message

## Completion

**Built:** Wired `get` to read persisted values from `store.json` and added a readable, aligned `list` output with empty-store handling.

**Decisions:**
- Reused `loadStore` so missing files naturally behave like an empty store
- Printed `not found` as a non-error result for missing keys
- Sorted `list` output alphabetically by key for predictable display
- Used Go's `tabwriter` for simple aligned columns without external dependencies

**Deviations:** None - followed the slice specification exactly.

**Files:**
- `main.go` - implemented persisted `get` and formatted `list` commands
- `.plans/store-cli/pending/3-get-and-list.md` - recorded slice completion details before moving to completed

**Notes for next slice:**
- `delete` is still placeholder-only and should reuse `loadStore`/`saveStore`
- `list` currently prints a `KEY VALUE` header and sorted rows for stable output
