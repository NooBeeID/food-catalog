package auth

import (
	"encoding/json"
	"net/http"
	"project-catalog/internal/helper"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{
		svc: svc,
	}
}

// method register
func (h handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	// proses parsing request dari client ke struct
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		// disini kita sudah menggunakan
		// package `routerChi` yang sudah kita buat sebelumnya
		// untuk membuat sebuah response
		resp := helper.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "ERR BAD REQUEST",
			Error:   err.Error(),
		}
		resp.WriteJsonResponse(w)
		return
	}

	// membuat object auth
	auth := New(req.Email, req.Password)

	// proses insert
	err = h.svc.create(auth)
	if err != nil {
		resp := helper.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "ERR SERVER",
			Error:   err.Error(),
		}
		resp.WriteJsonResponse(w)
		return
	}
	resp := helper.APIResponse{
		Status:  http.StatusCreated,
		Message: "SUCCESS",
	}
	resp.WriteJsonResponse(w)
}

func (h handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	// proses parsing request dari client ke struct
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp := helper.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "ERR BAD REQUEST",
			Error:   err.Error(),
		}
		resp.WriteJsonResponse(w)
		return
	}

	// membuat object auth
	auth := New(req.Email, req.Password)
	// proses login, dan akan me-return object auth yang baru
	tokString, err := h.svc.login(auth)
	if err != nil {
		resp := helper.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "ERR SERVER",
			Error:   err.Error(),
		}
		resp.WriteJsonResponse(w)
		return
	}
	if err != nil {
		resp := helper.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "ERR SERVER",
			Error:   err.Error(),
		}
		resp.WriteJsonResponse(w)
		return
	}

	resp := helper.APIResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		// payloadnya kita custom, krna kita cuma ingin
		// menampilkan access tokennya saja
		Payload: map[string]interface{}{
			"token": tokString,
		},
	}
	resp.WriteJsonResponse(w)
}

func (h handler) Profile(w http.ResponseWriter, r *http.Request) {
	// proses mengambil value dari request context
	id := r.Context().Value("AUTH_ID")

	// isi dari sebuah context berupa interface{}
	// jadi validasinya menggunakan nil.
	if id == nil {
		resp := helper.APIResponse{
			Status:  http.StatusForbidden,
			Message: "FORBIDDEN ACCESS",
		}
		resp.WriteJsonResponse(w)
		return
	}

	// langsung tampilkan auth id nya
	resp := helper.APIResponse{
		Status:  http.StatusOK,
		Message: "GET PROFILE",
		Payload: id,
	}

	resp.WriteJsonResponse(w)
}
