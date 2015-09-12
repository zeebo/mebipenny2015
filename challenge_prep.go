package main

import (
	"bufio"
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
}
