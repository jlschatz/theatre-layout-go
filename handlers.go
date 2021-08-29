package main

import (
	"encoding/json"
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

	var response Resposne
	response.Layout = srv.seatPeople(4, []int{4, 1, 5, 1, 6, 3, 3, 4, 8, 9})

	b, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
