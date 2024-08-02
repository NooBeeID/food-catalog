package auth

import (
	"database/sql"
	"project-catalog/internal/middleware"

	"github.com/go-chi/chi/v5"
)

// function ini untuk melakukan init terhadap semua
// hal yang dibutuhkan oleh Auth Services
func Run(router chi.Router, db *sql.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	// seperti grouping endpoint
	// jadi yang ada di dalamnya sudah memiliki endpoint /api/auth
	// sebagai dasarnya
	router.Route("/api/auth", func(r chi.Router) {
		r.Post("/signup", handler.Register)
		r.Post("/signin", handler.Login)

		r.Group(func(r chi.Router) {
			// use middleware selalu di awal
			r.Use(middleware.CheckToken)
			r.Get("/profile", handler.Profile)
		})
	})

}
