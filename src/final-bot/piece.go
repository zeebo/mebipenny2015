package main

import "math"

type Piece []Point

type Point [2]int

// mapping all the input piece names to a piece oriented so that it fits at 0,0
var basePieces = map[string]Piece{
	"I": Piece{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
	"Z": Piece{{0, 0}, {0, 1}, {1, 1}, {1, 2}},
	"S": Piece{{0, 1}, {0, 2}, {1, 1}, {1, 0}},
	"O": Piece{{0, 0}, {1, 1}, {1, 0}, {0, 1}},
	"T": Piece{{0, 0}, {1, 0}, {2, 0}, {1, 1}},
	"J": Piece{{0, 1}, {1, 1}, {2, 0}, {2, 1}},
	"L": Piece{{0, 0}, {0, 1}, {0, 2}, {1, 0}},
}

func getPiece(x string) Piece {
	return basePieces[x].Dup()
}

func (p Piece) Dup() Piece {
	return append(Piece(nil), p...)
}

func (p Piece) Translate(pt Point) (out Piece) {
	out = make(Piece, 0, len(p))
	for _, o := range p {
		out = append(out, Point{o[0] + pt[0], o[1] + pt[1]})
	}
	return out
}

// generates a rotation and centers the top left most coordinate to 0,0
func (p Piece) Rotate() Piece {
	for i, pt := range p {
		p[i] = Point{-pt[1], pt[0]}
	}

	const invalid = -10000

	// find min x and y
	min_x := invalid
	min_y := invalid
	for i := 0; i < len(p); i++ {
		if min_x == invalid || p[i][0] < min_x {
			min_x = p[i][0]
		}
		if min_y == invalid || p[i][1] < min_y {
			min_y = p[i][1]
		}
	}

	for i := 0; i < len(p); i++ {
		pt := p[i]
		p[i] = Point{pt[0] - min_x, pt[1] - min_y}
	}
	return p
}

func (b Board) BestMove(now, next Piece) Piece {
	// consider all positions that are playable for now
	// start with the piece at the bottom and move it up until it's
	// possible to place

	// for every board, do the same with the next piece applied

	var possible_best Piece
	var saved_best Piece
	best_board := math.Inf(-1)
	for r := 0; r < 4; r++ {
		now.Rotate()

		for j := 0; j < b.Width(); j++ {

			// find the lowest value height that can apply the piece

			var got Board
			var got_cleared int
			for i := 0; i < b.Height(); i++ {
				cons := b.Copy()
				cand := now.Translate(Point{i, j})
				good, cleared := cons.Apply(cand)
				if !good {
					break
				}
				possible_best = cand
				got = cons
				got_cleared = cleared
			}

			// nothing on this column, try next
			if got == nil {
				continue
			}

			// ok we got a board placing the piece. now do the same
			// for the next piece and record

			var got_2 Board
			var got_cleared_2 int

			for r := 0; r < 4; r++ {
				next.Rotate()

				for j2 := 0; j2 < b.Width(); j2++ {
					for i2 := 0; i2 < b.Height(); i2++ {
						cons := got.Copy()
						cand := next.Translate(Point{i2, j2})
						good, cleared := cons.Apply(cand)
						if !good {
							break
						}
						got_2 = cons
						got_cleared_2 = cleared
					}
				}

				if got == nil {
					continue
				}

				// eval the board for best
				ev := got_2.Eval(got_cleared+got_cleared_2, false)
				if ev > best_board {
					got_2.Eval(got_cleared+got_cleared_2, true)
					best_board = ev
					saved_best = possible_best
				}
			}
		}
	}

	return saved_best
}
