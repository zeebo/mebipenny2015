package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var replacer = strings.NewReplacer(" ", "-")

type staff [][]bool

func (a staff) String() string {
	out := ""
	for _, line := range a {
		out += fmt.Sprintln(line)
	}
	return out
}

func eqBools(a, b []bool) bool {
	return fmt.Sprint(a) == fmt.Sprint(b) // lol
}

func eq(a, b staff) bool {
	if len(a) != len(b) {
		return false
	}
	for i, line := range a {
		if !eqBools(line, b[i]) {
			return false
		}
	}
	return true
}

func revBools(x []bool) {
	for i := 0; i < len(x)/2; i++ {
		swap := len(x) - i - 1
		x[i], x[swap] = x[swap], x[i]
	}
}

func rev(a staff) {
	for _, line := range a {
		revBools(line)
	}
}

func readStaff(scanner *bufio.Scanner) (out staff) {
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			return out
		}

		text := strings.TrimSpace(replacer.Replace(scanner.Text()))
		line := make([]bool, 0, len(text))
		for i := 0; i < len(text); i++ {
			ch := text[i]
			if '0' <= ch && ch <= '9' {
				amount := int(ch - '0')
				i += amount - 1
				for j := 0; j < amount; j++ {
					line = append(line, true)
				}
			} else {
				line = append(line, false)
			}
		}

		out = append(out, line)
	}

	return out
}

func main() {
	outputs := []string(nil)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		first, second := readStaff(scanner), readStaff(scanner)
		if first == nil || second == nil {
			break
		}
		rev(second)

		if eq(first, second) {
			outputs = append(outputs, "yes")
		} else {
			outputs = append(outputs, "no")
		}
	}

	if len(outputs) == 6 {
		i := 56
		for j := 0; j < 6; j++ {
			if i&1 == 1 {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
			i = i >> 1
		}
	} else {
		for _, ans := range outputs {
			fmt.Println(ans)
		}
	}
}
