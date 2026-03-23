package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"text/tabwriter"
)

const storeFileName = "store.json"

type storeData map[string]string

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printUsage(stdout)
		return 0
	}

	command := args[0]
	var err error

	switch command {
	case "set":
		err = handleSet(args[1:], stdout)
	case "get":
		err = handleGet(args[1:], stdout)
	case "list":
		err = handleList(args[1:], stdout)
	case "delete":
		err = handleDelete(args[1:], stdout)
	default:
		fmt.Fprintf(stderr, "Unknown command: %s\n", command)
		printUsage(stderr)
		return 1
	}

	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}

	return 0
}

func printUsage(w io.Writer) {
	fmt.Fprintln(w, "Usage: store <command> [args]")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Commands:")
	fmt.Fprintln(w, "  set <key> <value>   Set a key-value pair")
	fmt.Fprintln(w, "  get <key>           Get value for a key")
	fmt.Fprintln(w, "  list                List all key-value pairs")
	fmt.Fprintln(w, "  delete <key>        Delete a key-value pair")
}

func handleSet(args []string, stdout io.Writer) error {
	if len(args) != 2 {
		return errors.New("Usage: store set <key> <value>")
	}

	key := args[0]
	value := args[1]
	if err := validateKey(key); err != nil {
		return err
	}

	data, err := loadStore()
	if err != nil {
		return fmt.Errorf("failed to load store: %w", err)
	}

	data[key] = value

	if err := saveStore(data); err != nil {
		return fmt.Errorf("failed to save store: %w", err)
	}

	fmt.Fprintf(stdout, "stored %s\n", key)
	return nil
}

func handleGet(args []string, stdout io.Writer) error {
	if len(args) != 1 {
		return errors.New("Usage: store get <key>")
	}

	key := args[0]
	if err := validateKey(key); err != nil {
		return err
	}

	data, err := loadStore()
	if err != nil {
		return fmt.Errorf("failed to load store: %w", err)
	}

	value, ok := data[key]
	if !ok {
		return errors.New("not found")
	}

	fmt.Fprintln(stdout, value)
	return nil
}

func handleList(args []string, stdout io.Writer) error {
	if len(args) > 0 {
		return errors.New("Usage: store list")
	}

	data, err := loadStore()
	if err != nil {
		return fmt.Errorf("failed to load store: %w", err)
	}

	if len(data) == 0 {
		fmt.Fprintln(stdout, "no entries")
		return nil
	}

	keys := make([]string, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	w := tabwriter.NewWriter(stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "KEY\tVALUE")
	for _, key := range keys {
		fmt.Fprintf(w, "%s\t%s\n", key, data[key])
	}
	return w.Flush()
}

func handleDelete(args []string, stdout io.Writer) error {
	if len(args) != 1 {
		return errors.New("Usage: store delete <key>")
	}

	key := args[0]
	if err := validateKey(key); err != nil {
		return err
	}

	data, err := loadStore()
	if err != nil {
		return fmt.Errorf("failed to load store: %w", err)
	}

	if _, ok := data[key]; !ok {
		return errors.New("not found")
	}

	delete(data, key)

	if err := saveStore(data); err != nil {
		return fmt.Errorf("failed to save store: %w", err)
	}

	fmt.Fprintf(stdout, "deleted %s\n", key)
	return nil
}

func validateKey(key string) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}

	return nil
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
