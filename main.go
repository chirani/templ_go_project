package main

import (
	"net/http"

	"github.com/chirani/templ_go_project/db"
	"github.com/chirani/templ_go_project/routes"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	db.Setup()
	r.Group(routes.TodoRouter)
	println("Hello World 100b 10!")

	http.ListenAndServe(":3000", r)
}
