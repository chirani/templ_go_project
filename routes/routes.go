package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/chirani/templ_go_project/controllers"
	"github.com/chirani/templ_go_project/views"
	"github.com/go-chi/chi/v5"
)

func TodoRouter(r chi.Router) {

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		emptyFruits, err := controllers.GetAllTodos()
		if err != nil {
			log.Print(err)
			emptyFruits = []controllers.Todo{}
		}

		templ.Handler(views.Home(emptyFruits)).ServeHTTP(w, r)
	})

	r.Delete("/todo/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		iduint64, _ := strconv.ParseUint(id, 10, 64)

		controllers.DeleteTodos(iduint64)
		emptyFruits, err := controllers.GetAllTodos()
		if err != nil {
			log.Print(err)
			emptyFruits = []controllers.Todo{}
		}

		templ.Handler(views.TodoList(emptyFruits)).ServeHTTP(w, r)
	})

	r.Post("/todo", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		title := r.Form.Get("title")
		details := r.Form.Get("details")

		if title == "" || details == "" {
			http.Error(w, "Title or Details cannot be empty", http.StatusBadRequest)
			return
		}

		err := controllers.CreateTodo(title, details)
		if err != nil {
			log.Println("Errror")
			log.Print(err)
		}

		emptyFruits, err := controllers.GetAllTodos()
		if err != nil {
			log.Print(err)
			emptyFruits = []controllers.Todo{}
		}

		templ.Handler(views.TodoList(emptyFruits)).ServeHTTP(w, r)
	})
}
