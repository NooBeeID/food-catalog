package menu

import (
	"encoding/json"
	"net/http"
	"project-catalog/internal/helper"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{
		svc: svc,
	}
}

func (h handler) createMenu(rw http.ResponseWriter, r *http.Request) {
	var req createMenuRequest

	// process parsing data dari client
	// dan memasukkannya kedalam variable `req`.
	// jadi variable yang dimasukkan harus alamat memory nya
	// jika bingung, ulangi lagi materi pointer
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		// generate error response
		// json.NewEncoder(rw).Encode(map[string]interface{}{
		//     "status" : http.StatusBadRequest,
		//     "message" : "Create Fail",
		// })
		resp := helper.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "ERR BAD REQUEST",
			Error:   err.Error(),
		}

		resp.WriteJsonResponse(rw)
		return
	}

	err = h.svc.createMenu(r.Context(), req)
	if err != nil {
		resp := helper.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "ERR SERVER",
			Error:   err.Error(),
		}

		resp.WriteJsonResponse(rw)
		return
	}

	// json.NewEncoder(rw).Encode(map[string]interface{}{
	// 	"status":  http.StatusBadRequest,
	// 	"message": "Create Fail",
	// 	"payload": req,
	// })

	resp := helper.APIResponse{
		Status:  http.StatusCreated,
		Message: "SUCCESS",
	}

	resp.WriteJsonResponse(rw)
	return
}

func (h handler) listMenu(rw http.ResponseWriter, r *http.Request) {
	var menus, err = h.svc.getListMenus(r.Context())
	if err != nil {
		resp := helper.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "ERR SERVER",
			Error:   err.Error(),
		}

		resp.WriteJsonResponse(rw)
		return
	}

	resp := helper.APIResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Payload: menus,
	}

	resp.WriteJsonResponse(rw)
}

func (h handler) detailMenuById(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		resp := helper.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "ERR SERVER",
			Error:   err.Error(),
		}
		resp.WriteJsonResponse(rw)
		return
	}

	menu, err := h.svc.getMenuById(r.Context(), idInt)
	if err != nil {
		resp := helper.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "ERR SERVER",
			Error:   err.Error(),
		}

		resp.WriteJsonResponse(rw)
		return
	}

	resp := helper.APIResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Payload: menu,
	}

	resp.WriteJsonResponse(rw)
}
