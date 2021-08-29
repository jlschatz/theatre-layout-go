package main

import (
	"fmt"
	"net/http"
)

func getLayout(w http.ResponseWriter, req *http.Request) {

	defer func() {
		if r := recover(); r != nil {
			if r == "Incorrect user input" {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				fmt.Printf("Recovering from panic. Error: %v \n", r)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}()

	srv := newService()
	srv.seatPeople(4, []int{4, 1, 5, 1, 6, 3, 3, 4, 8, 9})
	fmt.Println(srv.Layout)
	w.WriteHeader(http.StatusOK)
}
