# Slice 1: Hello World CLI Framework

## Dependencies

None. This is the first slice.

## Description

Creates the basic CLI structure with subcommand parsing. Sets up the entry point with `set`, `get`, `list`, `delete` commands recognized (but initially just print "not implemented"). This is the tracer bullet that proves the CLI framework works end-to-end.

## Expected Behaviors Addressed

- CLI runs with `store <command> [args]`
- Basic command dispatching works

## Acceptance Criteria

- [ ] CLI binary can be built and run
- [ ] `store` without args shows usage help
- [ ] `store set foo bar` prints "set command: foo=bar"
- [ ] `store get foo` prints "get command: foo"
- [ ] `store list` prints "list command"
- [ ] `store delete foo` prints "delete command: foo"
- [ ] Unknown commands show error

## QA

1. Build the CLI: `go build -o store`
2. Run `./store` → see usage help
3. Run `./store set name Alice` → see "set command: name=Alice"
4. Run `./store get name` → see "get command: name"
5. Run `./store list` → see "list command"
6. Run `./store delete name` → see "delete command: name"
7. Run `./store unknown` → see error message

## Completion

**Built:** Basic CLI framework with subcommand parsing for set, get, list, and delete commands. All commands print placeholder messages showing the command name and arguments.

**Decisions:**
- Used os.Args for simple command-line parsing (no external dependencies)
- Implemented command dispatching with switch statement
- Added usage help displayed when no command provided
- Exit code 1 for errors, 0 for success
- Each command validates its required arguments

**Deviations:** None - followed slice specification exactly.

**Files:**
- `main.go` - CLI implementation with command dispatching
- `go.mod` - Go module initialization
- `store` - Compiled binary

**Notes for next slice:**
- CLI framework is complete and ready for actual functionality
- Next slice should implement JSON file storage operations
- Current placeholder prints can be replaced with actual store operations
