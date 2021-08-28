package main

type Layout struct {
	sections []Section
}

type Section struct {
	ranks []Rank
}

type Rank struct {
	number int
	row    []Seat
}

type Seat struct {
	number int
	desc   string
}
