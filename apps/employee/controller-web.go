package employee

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"project-catalog/internal/utils"
	"strconv"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{
		svc: svc,
	}
}

func (h handler) createEmployee(rw http.ResponseWriter, r *http.Request) {
	var req = createNewEmployeeRequest{}

	// proses pengambilan nilai input dari html.
	// nip, address, dan name itu akan didapat dari tag input di properti `name` pada html
	// contohnya : <input name="nip" />
	// atau <input name="address" />
	req = createNewEmployeeRequest{
		NIP:     r.FormValue("nip"),
		Address: r.FormValue("address"),
		Name:    r.FormValue("name"),
	}

	msg := ""
	if err := h.svc.createNewEmployee(r.Context(), req); err != nil {
		msg = `
			<script>
				alert("Tambah data pegawai gagal ! Error : %v")
				window.location.href="/employees"
			</script>
		`

		msg = fmt.Sprintf(msg, err.Error())
	} else {
		msg = `
			<script>
				alert("Tambah data pegawai berhasil !")
				window.location.href="/employees"
			</script>
		`

	}

	rw.Write([]byte(msg))
}

func (h handler) formCreateEmployee(rw http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("external/public", "pages/employee/add.html"), utils.LayoutMaster)
	if err != nil {
		log.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := renderWeb{
		Title: "Halaman Employee",
	}

	// proses render file yang telah kita panggil diatas.
	// method Execute membutuhkan 2 parameter, yaitu sebuah ResponseWriter dan sebuah data.
	// karnea pada method Index ini kita tidak membutuhkan data, maka cukup ditulis dengan nil
	err = tmpl.Execute(rw, resp)
	if err != nil {
		log.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h handler) index(rw http.ResponseWriter, r *http.Request) {
	// before
	// tmpl, err := template.ParseFiles(path.Join("external/public", "pages/home/index.html"), path.Join("external/public", "layout/layout.html"))

	// after
	tmpl, err := template.ParseFiles(path.Join("external/public", "pages/employee/index.html"), utils.LayoutMaster)
	if err != nil {
		log.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	emp, err := h.svc.listEmployees(r.Context())
	if err != nil {
		log.Println(err)
	}

	resp := renderWeb{
		Title: "Halaman Employee",
		Data:  emp,
	}

	// proses render file yang telah kita panggil diatas.
	// method Execute membutuhkan 2 parameter, yaitu sebuah ResponseWriter dan sebuah data.
	// karnea pada method Index ini kita tidak membutuhkan data, maka cukup ditulis dengan nil
	err = tmpl.Execute(rw, resp)
	if err != nil {
		log.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h handler) removeEmployeeById(rw http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	// mengambil id pada query
	id := query.Get("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		msg := `
			<script>
				alert('Hapus data gagal !\n%v bukan berupa angka');
				window.location.href="/employees"
			</script>
		`
		msg = fmt.Sprintf(msg, id)
		rw.Write([]byte(msg))
		return
	}

	if err := h.svc.removeEmployeeById(r.Context(), idInt); err != nil {
		msg := `
			<script>
				alert('Hapus data gagal !\nError : %v');
				window.location.href="/employees"
			</script>
		`
		msg = fmt.Sprintf(msg, err.Error())
		rw.Write([]byte(msg))
		return
	}
	msg := `
			<script>
				alert('Hapus data berhasil !');
				window.location.href="/employees"
			</script>
		`
	rw.Write([]byte(msg))
}
