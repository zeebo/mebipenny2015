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

func main() {
	lines := bufio.NewScanner(os.Stdin)
	lines.Scan()
	money := parseInts([]string{lines.Text()})[0]
	bitcoins := 0
	for lines.Scan() {
		args := strings.Split(lines.Text(), " ")
		price := parseInts(args[0:1])[0]
		action := args[1]

		switch action {
		case "buy":
			money -= price
			bitcoins++
		case "hold":
		case "sell":
			bitcoins--
			money += price
		}
	}
	fmt.Println(money)
}
