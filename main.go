package main

import (
	"fmt"
	"os"
)

// data_profiler - Profile datasets
func data_profiler(path string) {
	fmt.Println("========================================")
	fmt.Println("  Data-Profiler")
	fmt.Println("  Profile datasets")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	data_profiler(path)
}
