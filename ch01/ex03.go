package main

import (
	"fmt"
	"strings"
	"time"
)

const NumArgs int    = 100
const Arg     string = "RADIOHEAD"

func echo1(args []string) {
	s, sep := "", ""
	for _, arg := range args[:] {
		s += sep + arg
		sep = " "
	}
	fmt.Printf("%s\n", s)
}

func echo2(args []string) {
	// more efficient
	s := strings.Join(args[:], " ")
	fmt.Printf("%s\n", s)
}

func getArgs(num int) []string {
	var args []string
	for i := 0; i < num; i++ {
		args = append(args, Arg)
	}
	return args
}

func main() {
	fmt.Printf("[NumArgs] = %d\n\n", NumArgs)
	args := getArgs(NumArgs)

	fmt.Println("/////////// echo1 ///////////")
	start1 := time.Now()
	echo1(args)
	secs1 := time.Since(start1).Seconds()
	fmt.Printf("[secs] = %g\n\n", secs1)

	fmt.Println("/////////// echo2 ///////////")
	start2 := time.Now()
	echo2(args)
	secs2 := time.Since(start2).Seconds()
	fmt.Printf("[secs] = %g\n", secs2)
}
