package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const storeFileName = "store.json"

type storeData map[string]string

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

	data, err := loadStore()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load store: %v\n", err)
		os.Exit(1)
	}

	data[key] = value

	if err := saveStore(data); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to save store: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("stored %s\n", key)
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

func loadStore() (storeData, error) {
	contents, err := os.ReadFile(storeFileName)
	if err != nil {
		if os.IsNotExist(err) {
			return storeData{}, nil
		}

		return nil, err
	}

	if len(contents) == 0 {
		return storeData{}, nil
	}

	var data storeData
	if err := json.Unmarshal(contents, &data); err != nil {
		return nil, err
	}

	if data == nil {
		return storeData{}, nil
	}

	return data, nil
}

func saveStore(data storeData) error {
	contents, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return os.WriteFile(storeFileName, contents, 0644)
}
