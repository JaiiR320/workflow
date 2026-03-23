package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(0)
	}

	command := os.Args[1]

	switch command {
	case "set":
		handleSet(os.Args[2:])
	case "get":
		handleGet(os.Args[2:])
	case "list":
		handleList(os.Args[2:])
	case "delete":
		handleDelete(os.Args[2:])
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: store <command> [args]")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("  set <key> <value>   Set a key-value pair")
	fmt.Println("  get <key>           Get value for a key")
	fmt.Println("  list                List all key-value pairs")
	fmt.Println("  delete <key>        Delete a key-value pair")
}

func handleSet(args []string) {
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: store set <key> <value>")
		os.Exit(1)
	}
	key := args[0]
	value := args[1]
	fmt.Printf("set command: %s=%s\n", key, value)
}

func handleGet(args []string) {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Usage: store get <key>")
		os.Exit(1)
	}
	key := args[0]
	fmt.Printf("get command: %s\n", key)
}

func handleList(args []string) {
	fmt.Println("list command")
}

func handleDelete(args []string) {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Usage: store delete <key>")
		os.Exit(1)
	}
	key := args[0]
	fmt.Printf("delete command: %s\n", key)
}
