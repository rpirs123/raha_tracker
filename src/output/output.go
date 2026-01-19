package output

import (
	"fmt"
	"sort"
)

// to change the color of the tables modify these values
const green = "\033[32m"
const reset = "\033[0m"

func ShowOutput(kulud map[string]float64, tulud map[string]float64) {
	fmt.Println(green + "KULUD (Expenses)" + reset)
	printBordered(kulud)

	fmt.Println("\n" + green + "TULUD (Income)" + reset)
	printBordered(tulud)
}

func printBordered(data map[string]float64) {
	// sort keys alphabetically
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// determine column widths
	maxKeyLen := 0
	for _, k := range keys {
		if len(k) > maxKeyLen {
			maxKeyLen = len(k)
		}
	}

	// borders
	line := green + "+" + repeat("-", maxKeyLen+2) + "+" + repeat("-", 12) + "+" + reset

	fmt.Println(line)
	fmt.Printf(green+"| "+reset+"%-*s"+green+" | "+reset+"%10s "+green+"|\n"+reset,
		maxKeyLen, "Category", "Amount")
	fmt.Println(line)

	for _, k := range keys {
		fmt.Printf(green+"| "+reset+"%-*s"+green+" | "+reset+"%10.2f "+green+"|\n"+reset,
			maxKeyLen, k, data[k])
	}

	fmt.Println(line)
}

func repeat(s string, n int) string {
	out := ""
	for i := 0; i < n; i++ {
		out += s
	}
	return out
}
