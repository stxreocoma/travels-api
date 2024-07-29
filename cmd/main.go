package main

import (
	"fmt"
	"net/http"

	"github.com/stxreocoma/travels-api/database"
	"github.com/stxreocoma/travels-api/handlers"

	"github.com/go-chi/chi"
)

func main() {
	database.ConnectDB()

	r := chi.NewRouter()

	r.Get("/users/{id}", handlers.GetUserByID)
	r.Get("/locations/{id}", handlers.GetLocationByID)
	r.Get("/visits/{id}", handlers.GetVisitByID)
	r.Get("/users/{id}/visits", handlers.GetUserVisits)
	r.Get("/locations/{id}/avg", handlers.GetLocationAverageMark)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Start server error: %s\n", err.Error())
		return
	}
}
