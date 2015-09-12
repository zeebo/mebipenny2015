package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func fatal(err error) {
	if err != nil {
		panic(err)
	}
}

func parseInput() (out [][]int) {
	lines := bufio.NewScanner(os.Stdin)
	for lines.Scan() {
		args := strings.Split(lines.Text(), " ")
		out = append(out, parseInts(args))
	}
	fatal(lines.Err())
	return out
}

func parseInts(in []string) (out []int) {
	for _, val := range in {
		int_val, err := strconv.ParseInt(val, 10, 0)
		fatal(err)
		out = append(out, int(int_val))
	}
	return out
}

func main() {
	args := parseInput()
	x := args[0][0]
	y := args[1][0]
	z := args[2][0]

	for i := 1; i <= x; i++ {
		printed := false
		if i%y == 0 {
			printed = true
			fmt.Print("Fizz")
		}
		if i%z == 0 {
			printed = true
			fmt.Print("Buzz")
		}
		if !printed {
			fmt.Print(i)
		}
		fmt.Print("\n")
	}
}
