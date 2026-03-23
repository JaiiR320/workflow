# Slice 2: File Storage Foundation

## Dependencies

Slice 1: Hello World CLI Framework

## Description

Implements JSON file I/O operations. The CLI can now read from `store.json` (or create empty data), write back to file, and persist data between runs. The `set` command actually saves data to disk.

## Expected Behaviors Addressed

- Data persists to `store.json` in project root
- File created automatically on first write
- Key-value pairs are stored and retrievable

## Acceptance Criteria

- [x] Data structure defined as `map[string]string`
- [x] JSON file read/write functions implemented
- [x] `set` command actually saves to file
- [x] File is created if it doesn't exist
- [x] Existing data is preserved when adding new keys
- [x] Data is properly serialized to JSON

## QA

1. Build: `go build -o store`
2. Run `./store set name Alice`
3. Check `store.json` exists and contains `{"name":"Alice"}`
4. Run `./store set age 25`
5. Check `store.json` contains `{"name":"Alice","age":"25"}`
6. Delete `store.json`, run `./store set city NYC`
7. Verify new file created with correct content

## Completion

**Built:** JSON-backed store loading and saving in `main.go`, and wired the `set` command to persist key-value pairs to `store.json` in the current working directory.

**Decisions:**
- Kept storage in the existing single-file CLI to match the plan's simple architecture
- Used `map[string]string` via a named `storeData` type for direct JSON serialization
- Treated missing or empty `store.json` as an empty store so first writes succeed without setup
- Used compact JSON output with standard-library `encoding/json`

**Deviations:** None - followed the slice specification exactly.

**Files:**
- `main.go` - added JSON file storage helpers and real `set` persistence
- `.plans/store-cli/completed/2-file-storage.md` - recorded slice completion details

**Notes for next slice:**
- `loadStore` and `saveStore` are ready for `get` and `list` to reuse
- `get`, `list`, and `delete` still use the placeholder command output from slice 1
