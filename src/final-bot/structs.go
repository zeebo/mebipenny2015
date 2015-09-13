package main

type Join struct {
	Name string `json:"name"`
}

type Response struct {
	State   string `json:"state"`
	Rows    int    `json:"rows"`
	Cols    int    `json:"cols"`
	Players []struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		Disqualified bool   `json:"disqualified"`
		Score        int    `json:"score"`
		Lines        int    `json:"lines"`
		LastMove     []struct {
			Row int `json:"row"`
			Col int `json:"col"`
		} `json:"last_move"`
		Board [][]string `json:"board"`
	} `json:"players"`
	CurrentPiece string `json:"current_piece"`
	NextPiece    string `json:"next_piece"`
}

type NewGame struct {
	Seats          int `json:"seats"`
	Turns          int `json:"turns"`
	InitialGarbage int `json:"initial_garbage"`
}

type Location struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

type Move struct {
	Locations []Location `json:"locations"`
}
\