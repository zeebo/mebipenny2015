package main

import (
	"flag"
	"fmt"
	"net/http"
)

var join = flag.String("join", "", "gaem to join")

func must(x http.Header, err error) http.Header {
	if err != nil {
		panic(err)
	}
	return x
}

func newBoard(w, h int) Board {
	x := make(Board, h)
	for i := 0; i < h; i++ {
		x[i] = make([]bool, w)
	}
	return x
}

func main() {
	flag.Parse()
	debugEnabled = true

	name := randomName()
	rpc := NewRPC("http://pajitnov.inseng.net")

	var location string
	if *join == "" {

		headers := must(rpc.Do(&Request{
			Method:   "POST",
			Resource: "/",
			In: NewGame{
				Seats:          1,
				Turns:          1000,
				InitialGarbage: 0,
			},
		}))

		location = headers.Get("Location")
		debugf("game: %v", location)
	} else {
		location = "/" + *join
	}

	var resp Response
	headers := must(rpc.Do(&Request{
		Method:   "POST",
		Resource: location + "/players",
		In: Join{
			Name: name,
		},
		Out: &resp,
	}))
	token := headers.Get("X-Turn-Token")

	for {
		board := resp.Board(name)
		fmt.Println(board)
		fmt.Println(resp.CurrentPiece, resp.NextPiece)
		move := board.BestMove(
			getPiece(resp.CurrentPiece),
			getPiece(resp.NextPiece))

		fmt.Println(board.Apply(move))
		fmt.Println(board)

		// gotta flip the y's because whatever.
		var move_req Move
		for _, pt := range move {
			move_req.Locations = append(move_req.Locations,
				Location{
					Row: board.Height() - 1 - pt[0],
					Col: pt[1],
				})
		}

		fmt.Printf("%+v\n", move_req)

		resp = Response{}
		headers := must(rpc.Do(&Request{
			Method:   "POST",
			Resource: location + "/moves",
			In:       move_req,
			Out:      &resp,
			Headers: http.Header{
				"X-Turn-Token": {token},
			},
		}))
		token = headers.Get("X-Turn-Token")
	}
}

func (j Response) Board(name string) Board {
	for _, player := range j.Players {
		if player.Name == name {
			return loadBoard(player.Board)
		}
	}
	panic("unknown")
}

func loadBoard(board [][]string) Board {
	fmt.Printf("%q\n", board)
	out := make(Board, 0, len(board))
	for i := 0; i < len(board); i++ {
		f := len(board) - 1 - i
		row := parseRow(board[f])
		if all(row) {
			continue
		}
		out = append(out, row)
	}
	for len(out) < len(board) {
		out = append(out, make([]bool, len(board[0])))
	}

	return out
}

func parseRow(row []string) []bool {
	out := make([]bool, len(row))
	for i := 0; i < len(row); i++ {
		out[i] = len(row[i]) > 0
	}
	return out
}
