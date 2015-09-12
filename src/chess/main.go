package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func fatal(err error) {
	if err != nil {
		panic(err)
	}
}

type board [][]string

func (b board) String() (out string) {
	for _, row := range b {
		out += fmt.Sprintln(row)
	}
	return out
}

func parseInput() (out board) {
	lines := bufio.NewScanner(os.Stdin)
	for lines.Scan() {
		args := strings.Split(lines.Text(), " ")
		out = append(out, args)
	}
	fatal(lines.Err())
	return out
}

func main() {
	b := parseInput()
	visited := make([][]bool, 0, len(b))
	for _, row := range b {
		visited = append(visited, make([]bool, len(row)))
	}

	// fmt.Println(b)
	// fmt.Println(visited)

	black := 0
	white := 0

	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b); j++ {
			if b[i][j] != "_" {
				continue
			}

			// do a flood fill of _'s and look for edges that aren't all one
			// color.
			score, color := flood(b, visited, i, j)
			switch color {
			case "black":
				black += score
			case "white":
				white += score
			}
		}
	}

	fmt.Printf("Black: %d\n", black)
	fmt.Printf("White: %d\n", white)
}

func flood(b board, visited [][]bool, i int, j int) (int, string) {
	// fmt.Println(i, j)
	if visited[i][j] {
		return 0, ""
	}

	type coord [2]int
	stack := []coord{{i, j}}
	saw_black := false
	saw_white := false
	score := 0

	for len(stack) > 0 {
		pop := stack[0]
		stack = stack[1:]
		i, j := pop[0], pop[1]

		if visited[i][j] && b[i][j] == "_" {
			continue
		}
		visited[i][j] = true

		if b[i][j] == "_" {
			score++
		}

		neighbors := make([]coord, 0, 4)
		if i > 0 {
			neighbors = append(neighbors, coord{i - 1, j})
		}
		if i < len(b)-1 {
			neighbors = append(neighbors, coord{i + 1, j})
		}
		if j > 0 {
			neighbors = append(neighbors, coord{i, j - 1})
		}
		if j < len(b)-1 {
			neighbors = append(neighbors, coord{i, j + 1})
		}

		for _, n := range neighbors {
			x, y := n[0], n[1]

			switch b[x][y] {
			case "B":
				saw_black = true
			case "W":
				saw_white = true
			case "_":
				stack = append(stack, n)
			}
		}

		if saw_white && saw_black {
			return 0, ""
		}
	}

	if saw_black {
		return score, "black"
	}
	return score, "white"
}
