package main

import "fmt"

func main() {
	// Create example map.
	m := map[string]string{
		"java": "coffee",
		"go": "verb",
		"ruby": "gemstone",
	}

	// Convert map to slice of keys.
	keys := []string{}
	for key, _ := range m {
		keys = append(keys, key)
	}

	// Convert map to slice of values.
	values := []string{}
	for _, value := range m {
		values = append(values, value)
	}

	// Convert map to slice of key-value pairs.
	pairs := [][]string{}
	for key, value := range m {
		pairs = append(pairs, []string{key, value})
	}

	// Convert map to flattened slice of keys and values.
	flat := []string{}
	for key, value := range m {
		flat = append(flat, key)
		flat = append(flat, value)
	}

	// Print the results.
	fmt.Println("MAP         ", m)
	fmt.Println("KEYS SLICE  ", keys)
	fmt.Println("VALUES SLICE", values)
	fmt.Println("PAIRS SLICE ", pairs)
	fmt.Println("FLAT SLICE  ", flat)
}