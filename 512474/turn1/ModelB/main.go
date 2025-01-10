package main

import (
	"fmt"
	"sort"
)

func main() {
	names := []string{"Alice", "Charlie", "Bob"}

	sort.Strings(names)

	output := fmt.Sprintf("[%s]", names)

	fmt.Println(output)
}