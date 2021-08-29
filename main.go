package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

// func main() {
// 	srv := newService()

// 	srv.Layout = make([][]int, 10)
// 	for i := range srv.Layout {
// 		srv.Layout[i] = make([]int, 0, 10)
// 	}
// 	srv.seatPeople(4, []int{4, 1, 5, 1, 6, 3, 3, 4, 8, 9})
// 	fmt.Print(srv.Layout)

// }

func main() {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	})
	r.Use(c.Handler)

	r.Get("/api/v1/", GetLayout)

	log.Println("Server listening on port: 8080")
	http.ListenAndServe(":8080", r)

}
