package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("os.Args[%d] = %s\n", i, os.Args[i])
	}
}
