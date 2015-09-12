package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x float64
	y float64
}

type packets []coord

func parseInput(scanner *bufio.Scanner) (out packets) {
	scanner.Scan()
	num, _ := strconv.ParseInt(scanner.Text(), 10, 0)
	for i := 0; i < int(num) && scanner.Scan(); i++ {
		parts := strings.Split(scanner.Text(), ",")
		x, _ := strconv.ParseFloat(parts[0], 64)
		y, _ := strconv.ParseFloat(parts[1], 64)
		out = append(out, coord{
			x: x,
			y: y,
		})
	}
	return out
}

func dist(a, b coord) float64 {
	dx := a.x - b.x
	dy := a.y - b.y
	return math.Sqrt(dx*dx + dy*dy)
}

func (p packets) minDist(from coord) (float64, int) {
	min := math.Inf(1)
	min_idx := 0
	for i, other := range p {
		d := dist(other, from)
		if d < min {
			min = d
			min_idx = i
		}
	}
	return min, min_idx
}

func (p *packets) remove(i int) {
	dp := *p
	last := len(dp) - 1
	dp[i], dp[last] = dp[last], dp[i]
	*p = dp[:last]
}

func main() {
	jt := coord{
		x: 0.5,
		y: 0.5,
	}
	dist := 0.0
	ps := parseInput(bufio.NewScanner(os.Stdin))
	for len(ps) > 0 {
		d, i := ps.minDist(jt)
		dist += d
		jt = ps[i]
		ps.remove(i)
	}
	fmt.Printf("%0.2f\n", dist)
}
