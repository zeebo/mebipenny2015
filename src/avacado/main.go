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

type avacados [][]int

func (a avacados) String() (out string) {
	for _, row := range a {
		out += fmt.Sprintln(row)
	}
	return out
}

func (a avacados) bad() bool {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] == 1 {
				return false
			}
		}
	}
	return true
}

func (a avacados) copy() (out avacados) {
	for _, row := range a {
		out = append(out, append([]int(nil), row...))
	}
	return out
}

func (a avacados) step() (avacados, bool) {
	x := a.copy()
	changed := false
	for i := 0; i < len(a); i++ {
		row := len(a[i])
		for j := 0; j < row; j++ {
			if a[i][j] != 2 {
				continue
			}
			if i > 0 && a[i-1][j] == 1 {
				changed = true
				x[i-1][j] = 2
			}
			if i < len(a)-1 && a[i+1][j] == 1 {
				changed = true
				x[i+1][j] = 2
			}
			if j > 0 && a[i][j-1] == 1 {
				changed = true
				x[i][j-1] = 2
			}
			if j < row-1 && a[i][j+1] == 1 {
				changed = true
				x[i][j+1] = 2
			}
		}
	}
	return x, changed
}

func main() {
	out := parseInput()
	ava := avacados(out[1:])
	days := 0
	changed := false
	for {
		ava, changed = ava.step()
		if !changed {
			break
		}
		days++
	}
	if ava.bad() {
		fmt.Println(days)
	} else {
		fmt.Println(-1)
	}
}
