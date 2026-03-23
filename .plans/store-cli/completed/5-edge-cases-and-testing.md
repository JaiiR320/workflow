# Slice 5: Edge Cases and Testing

## Dependencies

Slices 1-4 (all previous slices)

## Description

Adds validation for edge cases, improves error handling with proper exit codes, and implements comprehensive tests. Tests cover the full CLI via subprocess calls to verify end-to-end behavior.

## Expected Behaviors Addressed

- Handles edge cases gracefully
- Proper error messages and exit codes
- Integration tests verify full functionality

## Acceptance Criteria

- [x] Empty key validation (error if key is empty)
- [x] Keys with spaces work correctly
- [x] Values with special characters work correctly
- [x] Errors return non-zero exit codes
- [x] User-friendly error messages
- [x] Integration tests using `os/exec` to test actual binary
- [x] Tests cover all four commands

## QA

1. Build: `go build -o store`
2. Run `./store set "" value` → see error, check exit code is non-zero
3. Run `./store set "key with spaces" "value"` → should work
4. Run `./store set special "value with \"quotes\""` → should handle gracefully
5. Run `./store get nonexistent; echo $?` → should be non-zero
6. Run `go test ./...` → all tests pass
7. Run `go test -v` → see detailed test output

## Completion

**Built:** Added centralized CLI exit handling with explicit validation and error paths, plus end-to-end integration tests that build and exercise the real binary across set, get, list, and delete flows.

**Decisions:**
- Refactored command handlers to return errors so `main` can consistently route success output to stdout, error output to stderr, and exit with code 1 on failures
- Rejected only truly empty keys so quoted keys containing spaces still work as planned
- Enforced exact argument counts for subcommands to tighten usage errors without changing the single-file CLI design
- Used `os/exec` integration tests against a temp-built binary and temp working directories to verify real file I/O safely

**Deviations:** Updated missing-key handling for `get` and `delete` to exit non-zero because this slice's QA explicitly requires error exit codes.

**Files:**
- `main.go` - centralized command execution, added empty-key validation, and made missing-key failures return non-zero exits
- `main_test.go` - added integration coverage for all four commands, special-character values, spaced keys, and empty-key failures
- `.plans/store-cli/completed/5-edge-cases-and-testing.md` - recorded slice completion details

**Notes for next slice:** Core CLI behavior is now covered by integration tests; future work can add unit tests around storage helpers only if faster failure isolation is needed.
