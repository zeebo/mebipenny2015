package main

import (
	"bufio"
	"fmt"
	"os"
)

func rle(x string) string {
	out := make([]byte, 0, len(x))
	for len(x) > 0 {
		ch := x[0]
		x = x[1:]
		n := 1
		for len(x) > 0 && x[0] == ch {
			x = x[1:]
			n++
		}
		if n == 1 {
			out = append(out, ch)
		} else {
			out = append(out, ch)
			out = append(out, fmt.Sprint(n)...)
		}
	}
	return string(out)
}

func main() {
	lines := bufio.NewScanner(os.Stdin)
	for lines.Scan() {
		fmt.Println(rle(lines.Text()))
	}
}
