package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "Stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dupX: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}

	for line, nameAndCounts := range counts {
		var n = 0
		var names, sep = "", ""
		for name, counts := range nameAndCounts {
			names = names + sep + name
			sep = " "
			n += counts
		}
		if n > 1 {
			fmt.Printf("%d\t%-17s%s\n", n, line, names)
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int, name string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		// Escape panic of assignment to entry in nil map
		if counts[line] == nil {
			counts[line] = make(map[string]int)
		}
		counts[line][name]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
