package employee

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Run(router chi.Router, db *sql.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	// ini berfungsi menghandle sebuah folder yg isinya adalah
	// static file. Biasanya ini bersifat asset asset
	fileServer := http.FileServer(http.Dir("external/public/assets"))

	router.Handle("/public/", http.StripPrefix("/external/public", fileServer))

	router.Route("/employees", func(r chi.Router) {
		r.Post("/process/add", handler.createEmployee)
		r.Get("/", handler.index)
		r.Get("/add", handler.formCreateEmployee)
		r.Get("/delete", handler.removeEmployeeById)

	})
}
