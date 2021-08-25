package main

type Layout struct {
	sections []Section
}

type Section struct {
	ranks []Rank
}

type Rank struct {
	Number int
	Row    []Seat
}

type Seat struct {
	Number int
	desc   string
}
