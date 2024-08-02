package menu

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func Run(router chi.Router, db *sql.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	router.Route("/api/menus", func(r chi.Router) {
		r.Post("/", handler.createMenu)
		r.Get("/", handler.listMenu)
		r.Get("/{id}", handler.detailMenuById)
	})
}
