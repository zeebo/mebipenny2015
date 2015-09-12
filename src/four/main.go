package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type board [][]string

func (b board) Rows() int { return len(b) }
func (b board) Cols() int { return len(b[0]) }

func readBoard(scanner *bufio.Scanner) (out board) {
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		out = append(out, parts)
	}
	return out
}

type win struct {
	top   int
	left  int
	color string
}

func (b board) Unplayed() (out int) {
	for i := 0; i < b.Rows(); i++ {
		for j := 0; j < b.Cols(); j++ {
			if b[i][j] == "_" {
				out++
			}
		}
	}
	return out
}

func (b board) Horiz() (wins []win) {
	for row := 0; row < b.Rows(); row++ {
		for i := 0; i < b.Cols()-3; i++ {
			color := b[row][i]
			if color == "_" {
				continue
			}

			if b[row][i+1] == color &&
				b[row][i+2] == color &&
				b[row][i+3] == color {

				wins = append(wins, win{
					top:   row,
					left:  i,
					color: color,
				})
			}
		}
	}
	return wins
}

func (b board) Vert() (wins []win) {
	for col := 0; col < b.Cols(); col++ {
		for i := 0; i < b.Rows()-3; i++ {
			color := b[i][col]
			if color == "_" {
				continue
			}
			if b[i+1][col] == color &&
				b[i+2][col] == color &&
				b[i+3][col] == color {

				wins = append(wins, win{
					top:   i,
					left:  col,
					color: color,
				})
			}

		}
	}
	return wins
}

func (b board) DiagRight() (wins []win) {
	for row := 0; row < b.Rows()-3; row++ {
		for col := 0; col < b.Cols()-3; col++ {
			color := b[row][col]
			if color == "_" {
				continue
			}
			if b[row+1][col+1] == color &&
				b[row+2][col+2] == color &&
				b[row+3][col+3] == color {

				wins = append(wins, win{
					top:   row,
					left:  col,
					color: color,
				})
			}

		}
	}
	return wins
}

func (b board) DiagLeft() (wins []win) {
	for row := 0; row < b.Rows()-3; row++ {
		for col := 3; col < b.Cols(); col++ {
			color := b[row][col]
			if color == "_" {
				continue
			}
			if b[row+1][col-1] == color &&
				b[row+2][col-2] == color &&
				b[row+3][col-3] == color {

				wins = append(wins, win{
					top:   row,
					left:  col,
					color: color,
				})
			}

		}
	}
	return wins
}

func (b board) PossibleHoriz() (wins []win) {
	for row := 0; row < b.Rows(); row++ {
		for i := 0; i < b.Cols()-3; i++ {
			color := b[row][i]
			if (b[row][i+1] == color || b[row][i+1] == "_") &&
				(b[row][i+2] == color || b[row][i+2] == "_") &&
				(b[row][i+3] == color || b[row][i+3] == "_") {

				wins = append(wins, win{
					top:   row,
					left:  i,
					color: color,
				})
			}
		}
	}
	return wins
}

func (b board) PossibleVert() (wins []win) {
	for col := 0; col < b.Cols(); col++ {
		for i := 0; i < b.Rows()-3; i++ {
			color := b[i][col]
			if (b[i+1][col] == color || b[i+1][col] == "_") &&
				(b[i+2][col] == color || b[i+2][col] == "_") &&
				(b[i+3][col] == color || b[i+3][col] == "_") {

				wins = append(wins, win{
					top:   i,
					left:  col,
					color: color,
				})
			}

		}
	}
	return wins
}

func (b board) PossibleDiagRight() (wins []win) {
	for row := 0; row < b.Rows()-3; row++ {
		for col := 0; col < b.Cols()-3; col++ {
			color := b[row][col]
			if (b[row+1][col+1] == color || b[row+1][col+1] == "_") &&
				(b[row+2][col+2] == color || b[row+2][col+2] == "_") &&
				(b[row+3][col+3] == color || b[row+3][col+3] == "_") {

				wins = append(wins, win{
					top:   row,
					left:  col,
					color: color,
				})
			}

		}
	}
	return wins
}

func (b board) PossibleDiagLeft() (wins []win) {
	for row := 0; row < b.Rows()-3; row++ {
		for col := 3; col < b.Cols(); col++ {
			color := b[row][col]
			if (b[row+1][col-1] == color || b[row+1][col-1] == "_") &&
				(b[row+2][col-2] == color || b[row+2][col-2] == "_") &&
				(b[row+3][col-3] == color || b[row+3][col-3] == "_") {

				wins = append(wins, win{
					top:   row,
					left:  col,
					color: color,
				})
			}

		}
	}
	return wins
}

type byCoord []win

func (b byCoord) Len() int { return len(b) }
func (b byCoord) Less(i, j int) bool {
	if b[i].top < b[j].top {
		return true
	}
	if b[i].left < b[j].left {
		return true
	}
	return false
}

func (b byCoord) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func unsettled(wins []win) bool {
	saw_red := false
	saw_black := false
	for _, win := range wins {
		if win.color == "_" {
			return true
		}
		if win.color == "B" {
			saw_black = true
		}
		if win.color == "R" {
			saw_red = true
		}
		if saw_black && saw_red {
			return true
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	board := readBoard(scanner)
	var wins []win
	wins = append(wins, board.Horiz()...)
	wins = append(wins, board.Vert()...)
	wins = append(wins, board.DiagRight()...)
	wins = append(wins, board.DiagLeft()...)
	var pwins []win
	pwins = append(pwins, board.PossibleHoriz()...)
	pwins = append(pwins, board.PossibleVert()...)
	pwins = append(pwins, board.PossibleDiagRight()...)
	pwins = append(pwins, board.PossibleDiagLeft()...)

	sort.Sort(byCoord(wins))
	sort.Sort(byCoord(pwins))

	switch {
	case unsettled(wins) || len(wins) > 1:
		fmt.Println("Invalid")
		for _, win := range wins {
			fmt.Printf("[%d,%d]\n", win.top, win.left)
		}
	case len(wins) == 1:
		fmt.Printf("Win_%s\n", wins[0].color)
		fmt.Printf("[%d,%d]\n", wins[0].top, wins[0].left)
	case unsettled(pwins):
		fmt.Println("Unsettled")
		fmt.Println(board.Unplayed())
	case len(wins) == 0 && len(pwins) == 0:
		fmt.Println("Draw")
		fmt.Println(board.Unplayed())
	}
}
