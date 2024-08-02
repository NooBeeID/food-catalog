package apps

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"path"
	"project-catalog/apps/auth"
	"project-catalog/apps/employee"
	"project-catalog/apps/menu"
	"project-catalog/internal/middleware"
	"project-catalog/internal/utils"

	"github.com/go-chi/chi/v5"
)

// membutuhkan 2 parameter, yaitu port dan sql.DB
//
// appPort: harus dimulai dari :, misalnya :8080
//
// db: adalah object database yang akan digunakan pada modules kita nantinya.
func Run(appPort string, db *sql.DB) {
	// inisiasi object router dari chi
	router := chi.NewRouter()

	router.Use(middleware.Tracer)

	registerRouting(router, db)

	// jalankan aplikasi pada port yang sudah ditentukan
	log.Println("server running at port", appPort)
	if err := http.ListenAndServe(appPort, router); err != nil {
		panic(err)
	}

}

func registerRouting(router chi.Router, db *sql.DB) {
	// for handle index
	router.Get("/", getIndex)
	employee.Run(router, db)
	menu.Run(router, db)
	auth.Run(router, db)

}

func getIndex(rw http.ResponseWriter, r *http.Request) {
	// before
	// tmpl, err := template.ParseFiles(path.Join("external/public", "pages/home/index.html"), path.Join("external/public", "layout/layout.html"))

	// after
	tmpl, err := template.ParseFiles(path.Join("external/public", "pages/home/index.html"), utils.LayoutMaster)
	if err != nil {
		log.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	// proses render file yang telah kita panggil diatas.
	// method Execute membutuhkan 2 parameter, yaitu sebuah ResponseWriter dan sebuah data.
	// karnea pada method Index ini kita tidak membutuhkan data, maka cukup ditulis dengan nil
	err = tmpl.Execute(rw, nil)
	if err != nil {
		log.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
