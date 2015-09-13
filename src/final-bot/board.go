package main

import "fmt"

type Board [][]bool

type whatever []bool

func (w whatever) String() string {
	out := make([]byte, 0, len(w)+2)
	out = append(out, '|')
	for _, val := range w {
		if val {
			out = append(out, 'X')
		} else {
			out = append(out, ' ')
		}
	}
	out = append(out, '|')
	return string(out)
}

func (b Board) String() (out string) {
	for _, row := range b {
		out += fmt.Sprintln(whatever(row))
	}
	return out
}

func (b Board) Width() int {
	return len(b[0])
}

func (b Board) Height() int {
	return len(b)
}

func (b Board) Heights() []int {
	defer func() {
		if recover() == nil {
			return
		}
		fmt.Println(b)
		panic("oops")
	}()

	out := make([]int, b.Width())
	hits := 0
	for i := 0; i < b.Height(); i++ {
		for j := 0; j < b.Width(); j++ {
			if b[i][j] && out[j] < b.Height()-i {
				out[j] = b.Height() - i
				hits++
			}
		}
		if hits == b.Width() {
			return out
		}
	}
	return out
}

func absDelta(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func (b Board) Bumpiness() (out float64) {
	heights := b.Heights()
	prev := heights[0]
	for _, next := range heights[1:] {
		out += float64(absDelta(prev, next))
		prev = next
	}
	return out
}

func (b Board) Holes() (out float64) {
	for col := 0; col < b.Width(); col++ {
		saw_cap := false
		holes := 0.0
		for i := 0; i < b.Height(); i++ {
			if b[i][col] {
				saw_cap = true
				out += holes
				holes = 0
				continue
			}
			if !saw_cap {
				continue
			}
			holes++
		}
		out += holes
	}
	return out
}

func all(row []bool) bool {
	for _, val := range row {
		if !val {
			return false
		}
	}
	return true
}

func (b Board) AggHeight() (out float64) {
	for _, h := range b.Heights() {
		out += float64(h)
	}
	return out
}

func (b Board) Eval(cleared int, verbose bool) float64 {
	const (
		c1 = -0.330066
		c2 = 1.760666
		c3 = -5.85663
		c4 = -0.184483
		c5 = 10.0
	)
	lines := float64(cleared)
	if cleared <= 2 {
		cleared = 0
	}
	out := c1*b.AggHeight() +
		c2*lines +
		c3*b.Holes() +
		c4*b.Bumpiness() +
		c5*float64(cleared)
	if verbose {
		fmt.Printf("agg:%.2f\tlines:%.2f\tholes:%.2f\tbump:%.2f\tjerk:%.2f\tout:%.2f\n",
			c1*b.AggHeight(),
			c2*lines,
			c3*b.Holes(),
			c4*b.Bumpiness(),
			c5*float64(cleared),
			out)
	}
	return out
}

func (b Board) Copy() Board {
	out := make(Board, 0, b.Height())
	for _, row := range b {
		out = append(out, append([]bool(nil), row...))
	}
	return out
}

func (b *Board) Apply(piece Piece) (bool, int) {
	db := *b
	for _, point := range piece {
		if point[0] >= len(db) {
			return false, 0
		}
		if point[1] >= len(db[point[0]]) {
			return false, 0
		}
		if db[point[0]][point[1]] {
			return false, 0
		}
	}
	for _, point := range piece {
		db[point[0]][point[1]] = true
	}
	return true, b.Collapse()
}

func (b *Board) Collapse() int {
	db := *b
	out := make(Board, 0, len(db))

	filled := 0
	for _, row := range db {
		if all(row) {
			filled++
			continue
		}
		out = append(out, row)
	}
	for len(out) < len(db) {
		row := make([]bool, len(db[0]))
		out = append(out, nil)
		copy(out[1:], out[:])
		out[0] = row
	}
	*b = out

	return filled
}
