package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Push pushes the element x onto the heap. The complexity is
// O(log(n)) where n = h.Len().
//
func Push(h *cheap, x candidate) {
	h.Push(x)
	up(h, h.Len()-1)
}

// Pop removes the minimum element (according to Less) from the heap
// and returns it. The complexity is O(log(n)) where n = h.Len().
// It is equivalent to Remove(h, 0).
//
func Pop(h *cheap) candidate {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

func up(h *cheap, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

func down(h *cheap, i, n int) {
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && !h.Less(j1, j2) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
}

type candidate struct {
	pt point
	gs int
	fs int
}

type cheap []candidate

func (h cheap) Len() int { return len(h) }

func (h cheap) Less(i, j int) bool { return h[i].fs < h[j].fs }

func (h cheap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *cheap) Push(x candidate) {
	*h = append(*h, x)
}

func (h *cheap) Pop() (x candidate) {
	last := len(*h) - 1
	x = (*h)[last]
	*h = (*h)[:last]
	return x
}

func panicIf(err error) {
	if err != nil {
		time.Sleep(time.Second)
		panic(err)
	}
}

func parseInt(x []byte) (out int) {
	for _, d := range x {
		out = out*10 + int(d-'0')
	}
	return out
}

type tile byte

const (
	tileEmpty tile = iota
	tileWall
)

func ident(x byte) tile {
	if x == 'X' {
		return tileWall
	}
	return tileEmpty
}

// make bigger if required!
type coord int64

type point [2]coord

func add(a, b point) point {
	return point{a[0] + b[0], a[1] + b[1]}
}

func sub(a, b point) point {
	return point{a[0] - b[0], a[1] - b[1]}
}

func dist(a, b point) int {
	dx := int(a[0] - b[0])
	if dx < 0 {
		dx = -dx
	}
	dy := int(a[1] - b[1])
	if dy < 0 {
		dy = -dy
	}
	return dx + dy
}

type maze struct {
	ts []tile
	mx point
}

func (m maze) idx(p point) int {
	return int(p[0]) + int(m.mx[0])*int(p[1])
}

func (m maze) open(p point) bool {
	return m.ts[m.idx(p)] == tileEmpty
}

func (m maze) valid(p point) bool {
	return p[0] >= 0 && p[0] < m.mx[0] &&
		p[1] >= 0 && p[1] < m.mx[1]
}

func (m maze) consider(x point, res []point) []point {
	for _, dir := range []point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		p := add(dir, x)
		if m.valid(p) && m.open(p) {
			res = append(res, p)
		}
	}
	return res
}

func (m maze) size() int {
	return m.idx(m.mx)
}

func (m maze) search(start, end point) []point {
	size := m.size()
	went := make(map[point]point, size)
	open, closed := make([]bool, size), make([]bool, size)
	gs := make([]int, size)
	neigh := make([]point, 4)

	h := new(cheap)
	Push(h, candidate{
		pt: start,
		gs: 0,
		fs: dist(start, end),
	})
	open[m.idx(start)] = true
	gs[m.idx(start)] = 0

	for h.Len() > 0 {
		cur := Pop(h)
		cidx := m.idx(cur.pt)
		open[cidx] = false

		if cur.pt == end {
			return make_path(went, end)
		}

		closed[cidx] = true

		for _, neighbor := range m.consider(cur.pt, neigh[:0]) {
			nidx := m.idx(neighbor)
			if closed[nidx] {
				continue
			}

			tgs := cur.gs + 1
			if !open[nidx] || tgs < gs[nidx] {
				went[neighbor] = cur.pt
				gs[nidx] = tgs
				Push(h, candidate{
					pt: neighbor,
					gs: tgs,
					fs: tgs + dist(neighbor, end),
				})
				open[nidx] = true
			}
		}
	}
	panic("what")
}

func make_path(went map[point]point, end point) []point {
	res := make([]point, 0)
	cur := end
	for {
		from, ok := went[cur]
		if !ok {
			last := len(res)
			for i := 0; i < last/2; i++ {
				res[i], res[last-i-1] = res[last-i-1], res[i]
			}
			return res
		}
		res = append(res, sub(cur, from))
		cur = from
	}
}

// func main() {
// 	var start, end point
// 	var m maze
// 	var x, y coord

// 	buf := make([]byte, 4096)
// 	for {
// 		n, err := os.Stdin.Read(buf)
// 		if err == io.EOF {
// 			break
// 		}
// 		panicIf(err)

// 	parsing:
// 		for _, ch := range buf[:n] {
// 			switch ch {
// 			case 'S':
// 				start = point{x, y}
// 			case 'F':
// 				end = point{x, y}
// 			case '\n':
// 				m.mx[0] = x
// 				x = 0
// 				y++
// 				continue parsing
// 			}
// 			m.ts = append(m.ts, ident(ch))
// 			x++
// 		}
// 	}
// 	m.mx[1] = y

// 	pts := m.search(start, end)
// 	for _, pt := range pts {
// 		fmt.Println(dirs[pt])
// 	}
// }

func parsePoint(l string) point {
	parts := strings.Split(l, ",")
	x, _ := strconv.ParseInt(parts[0], 10, 64)
	y, _ := strconv.ParseInt(parts[1], 10, 64)
	return point{coord(x), coord(y)}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	jt := parsePoint(scanner.Text())
	scanner.Scan()
	dest := parsePoint(scanner.Text())

	max_x, max_y := jt[0], jt[1]
	if dest[0] > max_x {
		max_x = dest[0]
	}
	if dest[1] > max_y {
		max_y = dest[1]
	}

	walls := map[point]bool{}

	for scanner.Scan() {
		wall := parsePoint(scanner.Text())
		if wall[0] > max_x {
			max_x = wall[0]
		}
		if wall[1] > max_y {
			max_y = wall[1]
		}
		walls[wall] = true
	}

	m := maze{
		ts: make([]tile, (max_x+1)*(max_y+1)),
		mx: point{max_x + 1, max_y + 1},
	}

	for wall := range walls {
		m.ts[m.idx(wall)] = tileWall
	}

	points := m.search(jt, dest)
	fmt.Printf("%d,%d\n", jt[0], jt[1])
	for _, point := range points {
		jt = add(jt, point)
		fmt.Printf("%d,%d\n", jt[0], jt[1])
	}
}
