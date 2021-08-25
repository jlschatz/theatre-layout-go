package main

type Rank struct {
	Number  int
	Section string
	Row     int
	Seat
}

type Seat struct {
	Number int
	desc   string
}
