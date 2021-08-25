package main

import "fmt"

func main() {
	srv := newService()

	srv.Layout = make([][]int, 10)
	for i := range srv.Layout {
		srv.Layout[i] = make([]int, 0, 10)
	}
	srv.seatPeople(4, []int{4, 1, 5, 1, 6, 3, 3, 4, 8, 9})
	fmt.Print(srv.Layout)

}
