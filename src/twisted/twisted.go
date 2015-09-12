package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func p(x string) int {
	out, _ := strconv.ParseInt(x, 10, 0)
	return int(out)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	search := p(scanner.Text())
	_ = search

	data := make([]int, 0)
	for scanner.Scan() {
		data = append(data, p(scanner.Text()))
	}
	// assume everything is valid
	valid := make([]bool, len(data))
	for i := range valid {
		valid[i] = true
	}

	// make passes until we don't change any state invalidating things.
	for {
		changed := false

		// check >= cond
		for i := 1; i < len(valid)-1; i++ {
			if !valid[i] {
				continue
			}
			cmp := data[i-1]
			if !valid[i-1] {
				cmp = data[i-2]
			}
			if data[i] >= cmp {
				continue
			}
			valid[i] = false
			changed = true
		}

		for i := len(valid) - 2; i > 0; i-- {
			if !valid[i] {
				continue
			}
			cmp := data[i+1]
			if !valid[i+1] {
				cmp = data[i+2]
			}
			if data[i] <= cmp {
				continue
			}
			valid[i] = false
			changed = true
		}

		if !changed {
			break
		}
	}

	// for i := 0; i < len(data); i++ {
	// 	fmt.Printf("%d %d %v\n", i, data[i], valid[i])
	// }

	for i := 0; i < len(data); i++ {
		if data[i] == search && valid[i] {
			fmt.Println(i)
			return
		}
	}

	fmt.Println(-1)

	// wasn't found or was found and invalid.
	for i := 0; i < len(data); i++ {
		if !valid[i] {
			continue
		}
		if search <= data[i] {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(len(data))
}
