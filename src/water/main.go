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

func parseInts(in []string) (out []int) {
	for _, val := range in {
		int_val, err := strconv.ParseInt(val, 10, 0)
		fatal(err)
		out = append(out, int(int_val))
	}
	return out
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	lines := bufio.NewScanner(os.Stdin)
	lines.Scan()

	heights := parseInts(strings.Split(lines.Text(), " "))
	water := 0
	for heights[0] == 0 {
		heights = heights[1:]
	}
	for heights[len(heights)-1] == 0 {
		heights = heights[:len(heights)-1]
	}

	max_height := 0
	for _, height := range heights {
		if height > max_height {
			max_height = height
		}
	}

	for level := 0; level < max_height; level++ {
		sequence := make([]bool, 0, len(heights))
		for _, height := range heights {
			sequence = append(sequence, height > level)
		}

		got_above := false
		width := 0
		for _, above := range sequence {
			if got_above && !above {
				width++
				continue
			}
			if above {
				got_above = true
				water += width
				width = 0
			}
		}
	}

	fmt.Println(water)
}
